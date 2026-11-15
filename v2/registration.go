package di

import (
	"context"
	"fmt"
	"io"
	"sync"
)

// kind задает жизненный цикл зависимости. Тип внутренний: снаружи жизненный
// цикл выбирается методом регистрации.
type kind int

const (
	kindSingleton kind = iota
	kindTransient
	kindScoped
)

// registration одна зарегистрированная зависимость типа T.
//
// Фабрика получает [context.Context] и состояние разрешения, поэтому
// сгенерированные [Container.Provide] фабрики умеют передавать по цепочке
// посещенные ключи (детект циклов) и текущий скоуп.
type registration[T any] struct {
	key     any
	name    string
	kind    kind
	deps    []any
	factory func(ctx context.Context, st *resolveState) (T, error)

	mu       sync.Mutex
	instance *T
	closer   io.Closer
	closed   bool
}

// registrationInfo недженерик-представление регистрации, нужное [Container.Verify].
type registrationInfo interface {
	registrationKey() any
	dependencies() []any
}

// disposable закрывает принадлежащий владельцу экземпляр ровно один раз.
type disposable interface {
	dispose() error
}

var (
	_ registrationInfo = (*registration[int])(nil)
	_ disposable       = (*registration[int])(nil)
	_ disposable       = (*scopedValue[int])(nil)
)

func (r *registration[T]) resolveSingleton(ctx context.Context, st *resolveState) (T, error) {
	var zero T

	r.mu.Lock()
	defer r.mu.Unlock()

	if r.instance != nil {
		return *r.instance, nil
	}

	if r.factory == nil {
		return zero, fmt.Errorf("%w: %s", ErrNoFactory, r.name)
	}

	instance, err := r.factory(ctx, st)
	if err != nil {
		return zero, err
	}

	r.instance = &instance

	if closer, ok := any(instance).(io.Closer); ok {
		r.closer = closer
	}

	return instance, nil
}

func (r *registration[T]) resolveTransient(ctx context.Context, st *resolveState) (T, error) {
	var zero T

	if r.factory == nil {
		return zero, fmt.Errorf("%w: %s", ErrNoFactory, r.name)
	}

	return r.factory(ctx, st)
}

// dispose закрывает singleton-экземпляр, если он реализует [io.Closer].
// Транзиенты и scoped-зависимости контейнеру не принадлежат, поэтому для них
// метод ничего не делает.
func (r *registration[T]) dispose() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.closed {
		return nil
	}

	r.closed = true

	if r.kind != kindSingleton || r.closer == nil {
		return nil
	}

	return r.closer.Close()
}

func (r *registration[T]) registrationKey() any { return r.key }

func (r *registration[T]) dependencies() []any { return r.deps }

// scopedValue кэширует scoped-экземпляр одного скоупа.
type scopedValue[T any] struct {
	owner *registration[T]

	mu       sync.Mutex
	instance *T
	closer   io.Closer
	closed   bool
}

func (v *scopedValue[T]) get(ctx context.Context, st *resolveState) (T, error) {
	var zero T

	v.mu.Lock()
	defer v.mu.Unlock()

	if v.instance != nil {
		return *v.instance, nil
	}

	if v.owner.factory == nil {
		return zero, fmt.Errorf("%w: %s", ErrNoFactory, v.owner.name)
	}

	instance, err := v.owner.factory(ctx, st)
	if err != nil {
		return zero, err
	}

	v.instance = &instance

	if closer, ok := any(instance).(io.Closer); ok {
		v.closer = closer
	}

	return instance, nil
}

func (v *scopedValue[T]) dispose() error {
	v.mu.Lock()
	defer v.mu.Unlock()

	if v.closed {
		return nil
	}

	v.closed = true

	if v.closer == nil {
		return nil
	}

	return v.closer.Close()
}

// newInstanceRegistration создает регистрацию готового экземпляра (singleton).
func newInstanceRegistration[T any](name string, instance T) *registration[T] {
	r := &registration[T]{
		key:      any(newTypeKey[T](name)),
		name:     name,
		kind:     kindSingleton,
		instance: &instance,
	}

	if closer, ok := any(instance).(io.Closer); ok {
		r.closer = closer
	}

	return r
}

// newFactoryRegistration оборачивает пользовательскую фабрику в internal
// фабрику с состоянием. Если фабрика nil, регистрация создается без фабрики,
// чтобы разрешение вернуло [ErrNoFactory].
func newFactoryRegistration[T any](name string, k kind, factory func(ctx context.Context) (T, error)) *registration[T] {
	r := &registration[T]{
		key:  any(newTypeKey[T](name)),
		name: name,
		kind: k,
	}

	if factory != nil {
		r.factory = func(ctx context.Context, _ *resolveState) (T, error) {
			return factory(ctx)
		}
	}

	return r
}

// noContext адаптирует фабрику без контекста, сохраняя nil как nil.
func noContext[T any](factory func() (T, error)) func(ctx context.Context) (T, error) {
	if factory == nil {
		return nil
	}

	return func(context.Context) (T, error) {
		return factory()
	}
}

// castAs приводит конкретный тип T к интерфейсу I.
//
// Go не позволяет статически проверить «T реализует I» для двух
// неограниченных параметров типа, поэтому проверка выполняется во время
// выполнения: при несовпадении рантайм паникует с сообщением о конверсии.
func castAs[I any, T any](instance T) I {
	return any(instance).(I)
}

func zeroOf[T any]() T {
	var zero T

	return zero
}
