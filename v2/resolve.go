package di

import (
	"context"
	"fmt"
)

// keyChain цепочка ключей, разрешаемых в рамках одного вызова.
//
// Проверка цепочки идет до взятия мьютекса регистрации, поэтому повторный
// вход в тот же ключ возвращает [ErrCircularDependency] вместо зависания на
// мьютексе. Цепочка неизменяемая: push создает новый узел.
type keyChain struct {
	parent *keyChain
	key    any
}

func (ch *keyChain) contains(key any) bool {
	for node := ch; node != nil; node = node.parent {
		if node.key == key {
			return true
		}
	}

	return false
}

func (ch *keyChain) push(key any) *keyChain {
	return &keyChain{parent: ch, key: key}
}

// resolveState контекст разрешения: скоуп, в котором идёт разрешение, и
// цепочка посещенных ключей.
type resolveState struct {
	scope *Scope
	chain *keyChain
}

func (st *resolveState) pushed(key any) *resolveState {
	return &resolveState{scope: st.scope, chain: st.chain.push(key)}
}

// resolveFrom разрешает зависимость key в скоупе состояния.
func resolveFrom[T any](ctx context.Context, st *resolveState, key typeKey[T]) (T, error) {
	var zero T

	if st.chain.contains(any(key)) {
		return zero, fmt.Errorf("%w: %s", ErrCircularDependency, key.name)
	}

	raw, ok := st.scope.lookup(any(key))
	if !ok {
		return zero, fmt.Errorf("%w: %s", ErrNotFound, key.name)
	}

	reg, ok := raw.(*registration[T])
	if !ok {
		return zero, fmt.Errorf("%w: %s", ErrInvalidRegistration, key.name)
	}

	next := st.pushed(any(key))

	switch reg.kind {
	case kindSingleton:
		return reg.resolveSingleton(ctx, next)
	case kindTransient:
		return reg.resolveTransient(ctx, next)
	case kindScoped:
		return st.scope.scoped(any(key), reg).get(ctx, next)
	default:
		return zero, fmt.Errorf("%w: %s", ErrUnsupportedKind, key.name)
	}
}

// keyName возвращает человекочитаемое имя ключа для сообщений об ошибках.
func keyName(key any) string {
	if named, ok := key.(fmt.Stringer); ok {
		return named.String()
	}

	return ""
}

// resolveRegistration разрешает конкретную регистрацию, минуя поиск по ключу.
func resolveRegistration[T any](ctx context.Context, s *Scope, reg *registration[T]) (T, error) {
	st := &resolveState{scope: s, chain: (&keyChain{}).push(reg.key)}

	switch reg.kind {
	case kindSingleton:
		return reg.resolveSingleton(ctx, st)
	case kindTransient:
		return reg.resolveTransient(ctx, st)
	case kindScoped:
		return s.scoped(reg.key, reg).get(ctx, st)
	default:
		return zeroOf[T](), fmt.Errorf("%w: %s", ErrUnsupportedKind, reg.name)
	}
}
