package di

import "context"

// Scope область разрешения зависимостей.
//
// Скоуп наследует регистрации родителя (вплоть до контейнера) и может
// переопределять их локально через [Scope.Register]/[Scope.RegisterNamed].
// Порядок поиска: скоуп -> родитель -> ... -> корень.
//
// Scoped-зависимость создается один раз внутри скоупа и кэшируется в нём;
// синглтоны и транзиенты разрешаются через корневой контейнер. Открытые
// скоупы закрывает вызывающий через [Scope.Dispose]: контейнер их не
// отслеживает.
type Scope struct {
	parent *Scope
	data   *scopeData
}

func newRootScope() *Scope {
	return &Scope{data: newScopeData()}
}

// BeginScope создает вложенный скоуп.
func (s *Scope) BeginScope() *Scope {
	return &Scope{parent: s, data: newScopeData()}
}

// Register регистрирует готовый экземпляр в скоупе, переопределяя регистрацию
// родителя для этого и вложенных скоупов.
func (s *Scope) Register[T any](instance T) {
	s.data.store(any(newTypeKey[T]("")), newInstanceRegistration("", instance), false)
}

// RegisterNamed регистрирует готовый экземпляр под именем name локально в
// скоупе.
func (s *Scope) RegisterNamed[T any](name string, instance T) {
	s.data.store(any(newTypeKey[T](name)), newInstanceRegistration(name, instance), false)
}

// Resolve возвращает зарегистрированную зависимость типа T.
func (s *Scope) Resolve[T any]() (T, error) {
	return s.ResolveContext[T](context.Background())
}

// ResolveNamed возвращает зависимость типа T, зарегистрированную под именем
// name.
func (s *Scope) ResolveNamed[T any](name string) (T, error) {
	return s.ResolveNamedContext[T](context.Background(), name)
}

// ResolveContext возвращает зависимость типа T, пробрасывая ctx в фабрику.
// Для синглтона используется контекст первого вызова.
func (s *Scope) ResolveContext[T any](ctx context.Context) (T, error) {
	return resolveFrom(ctx, &resolveState{scope: s}, newTypeKey[T](""))
}

// ResolveNamedContext возвращает именованную зависимость типа T, пробрасывая
// ctx в фабрику.
func (s *Scope) ResolveNamedContext[T any](ctx context.Context, name string) (T, error) {
	return resolveFrom(ctx, &resolveState{scope: s}, newTypeKey[T](name))
}

// MustResolve возвращает зависимость типа T и паникует, если разрешить ее не
// удалось.
func (s *Scope) MustResolve[T any]() T {
	return mustResolve(s.Resolve[T]())
}

// MustResolveNamed возвращает именованную зависимость типа T и паникует, если
// разрешить ее не удалось.
func (s *Scope) MustResolveNamed[T any](name string) T {
	return mustResolve(s.ResolveNamed[T](name))
}

// MustResolveContext возвращает зависимость типа T и паникует, если разрешить
// ее не удалось.
func (s *Scope) MustResolveContext[T any](ctx context.Context) T {
	return mustResolve(s.ResolveContext[T](ctx))
}

// ResolveAll возвращает все зарегистрированные зависимости типа T в порядке
// регистрации. Если таких зависимостей нет, возвращается пустой срез и nil.
//
// Учитываются регистрации корневого контейнера: локальные переопределения
// скоупа в список не добавляются, но видны через [Scope.Resolve].
func (s *Scope) ResolveAll[T any]() ([]T, error) {
	order := s.rootScope().data.snapshotOrder()

	result := make([]T, 0, len(order))

	for _, raw := range order {
		reg, ok := raw.(*registration[T])
		if !ok {
			continue
		}

		value, err := resolveRegistration(context.Background(), s, reg)
		if err != nil {
			return nil, err
		}

		result = append(result, value)
	}

	return result, nil
}

// Dispose закрывает scoped-экземпляры этого скоупа в обратном порядке
// создания (LIFO). Скоупы, созданные от этого, закрывает вызывающий:
// родитель их не отслеживает.
func (s *Scope) Dispose() error {
	return disposeAll(s.data.snapshotScopedOrder())
}

func (s *Scope) lookup(key any) (any, bool) {
	for current := s; current != nil; current = current.parent {
		if entry, ok := current.data.load(key); ok {
			return entry, true
		}
	}

	return nil, false
}

func (s *Scope) scoped[T any](key any, owner *registration[T]) *scopedValue[T] {
	value := s.data.getOrCreateScoped(key, func() any {
		return &scopedValue[T]{owner: owner}
	})

	return value.(*scopedValue[T])
}

func (s *Scope) rootScope() *Scope {
	current := s
	for current.parent != nil {
		current = current.parent
	}

	return current
}

func mustResolve[T any](value T, err error) T {
	if err != nil {
		panic("failed to resolve dependency: " + err.Error())
	}

	return value
}
