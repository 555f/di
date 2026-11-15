package di

import "context"

// Container хранит зарегистрированные зависимости.
type Container struct {
	root *Scope
}

// NewContainer создает пустой контейнер зависимостей.
func NewContainer() *Container {
	return &Container{root: newRootScope()}
}

// BeginScope создает дочерний скоуп контейнера.
//
// Scoped-зависимости создаются один раз внутри скоупа и кэшируются в нём.
func (c *Container) BeginScope() *Scope {
	return c.root.BeginScope()
}

// Register регистрирует готовый экземпляр зависимости.
func (c *Container) Register[T any](instance T) {
	c.root.Register(instance)
}

// RegisterNamed регистрирует готовый экземпляр под именем name.
func (c *Container) RegisterNamed[T any](name string, instance T) {
	c.root.RegisterNamed(name, instance)
}

// RegisterMany добавляет экземпляры типа T, не перезаписывая ранее
// зарегистрированные. Порядок сохраняется.
func (c *Container) RegisterMany[T any](instances ...T) {
	for _, instance := range instances {
		c.root.data.store(any(newTypeKey[T]("")), newInstanceRegistration("", instance), true)
	}
}

// RegisterFactory регистрирует фабрику зависимости. Фабрика вызывается не
// более одного раза.
func (c *Container) RegisterFactory[T any](factory func() (T, error)) {
	c.storeFactory(newTypeKey[T](""), kindSingleton, noContext(factory))
}

// RegisterFactoryNamed регистрирует именованную фабрику зависимости.
func (c *Container) RegisterFactoryNamed[T any](name string, factory func() (T, error)) {
	c.storeFactory(newTypeKey[T](name), kindSingleton, noContext(factory))
}

// RegisterTransient регистрирует фабрику, которая вызывается при каждом разрешении зависимости.
func (c *Container) RegisterTransient[T any](factory func() (T, error)) {
	c.storeFactory(newTypeKey[T](""), kindTransient, noContext(factory))
}

// RegisterTransientNamed регистрирует именованную фабрику, которая вызывается
// при каждом разрешении зависимости.
func (c *Container) RegisterTransientNamed[T any](name string, factory func() (T, error)) {
	c.storeFactory(newTypeKey[T](name), kindTransient, noContext(factory))
}

// RegisterScoped регистрирует scoped зависимость: она создается один раз в каждом скоупе.
func (c *Container) RegisterScoped[T any](factory func() (T, error)) {
	c.storeFactory(newTypeKey[T](""), kindScoped, noContext(factory))
}

// RegisterScopedNamed регистрирует именованную scoped зависимость.
func (c *Container) RegisterScopedNamed[T any](name string, factory func() (T, error)) {
	c.storeFactory(newTypeKey[T](name), kindScoped, noContext(factory))
}

// RegisterFactoryContext регистрирует фабрику, которая получает контекст
// разрешения. Для синглтона используется контекст первого вызова.
func (c *Container) RegisterFactoryContext[T any](factory func(ctx context.Context) (T, error)) {
	c.storeFactory(newTypeKey[T](""), kindSingleton, factory)
}

// RegisterFactoryContextNamed регистрирует именованную фабрику-синглтон с контекстом.
func (c *Container) RegisterFactoryContextNamed[T any](name string, factory func(ctx context.Context) (T, error)) {
	c.storeFactory(newTypeKey[T](name), kindSingleton, factory)
}

// RegisterTransientContext регистрирует фабрику, которая вызывается при каждом
// разрешении зависимости с контекстом.
func (c *Container) RegisterTransientContext[T any](factory func(ctx context.Context) (T, error)) {
	c.storeFactory(newTypeKey[T](""), kindTransient, factory)
}

// RegisterTransientContextNamed регистрирует именованную фабрику, которая
// вызывается при каждом разрешении зависимости с контекстом.
func (c *Container) RegisterTransientContextNamed[T any](name string, factory func(ctx context.Context) (T, error)) {
	c.storeFactory(newTypeKey[T](name), kindTransient, factory)
}

// RegisterScopedContext регистрирует scoped фабрику с контекстом.
func (c *Container) RegisterScopedContext[T any](factory func(ctx context.Context) (T, error)) {
	c.storeFactory(newTypeKey[T](""), kindScoped, factory)
}

// RegisterScopedContextNamed регистрирует именованную scoped фабрику с контекстом.
func (c *Container) RegisterScopedContextNamed[T any](name string, factory func(ctx context.Context) (T, error)) {
	c.storeFactory(newTypeKey[T](name), kindScoped, factory)
}

// RegisterAs регистрирует конкретный экземпляр T под интерфейсом I.
func (c *Container) RegisterAs[I, T any](instance T) {
	c.Register(castAs[I](instance))
}

// RegisterFactoryAs регистрирует фабрику конкретного типа T под интерфейсом I.
func (c *Container) RegisterFactoryAs[I, T any](factory func() (T, error)) {
	c.registerAsFactory(newTypeKey[I](""), kindSingleton, factory)
}

// RegisterTransientAs регистрирует фабрику конкретного типа T под интерфейсом
// I; фабрика вызывается при каждом разрешении зависимости.
func (c *Container) RegisterTransientAs[I, T any](factory func() (T, error)) {
	c.registerAsFactory(newTypeKey[I](""), kindTransient, factory)
}

// Resolve возвращает зарегистрированную зависимость типа T.
func (c *Container) Resolve[T any]() (T, error) {
	return c.root.Resolve[T]()
}

// ResolveNamed возвращает именованную зависимость типа T.
func (c *Container) ResolveNamed[T any](name string) (T, error) {
	return c.root.ResolveNamed[T](name)
}

// ResolveContext возвращает зависимость типа T, пробрасывая ctx в фабрику
// (отмена, дедлайн, request-значения). Для синглтона используется контекст
// первого вызова.
func (c *Container) ResolveContext[T any](ctx context.Context) (T, error) {
	return c.root.ResolveContext[T](ctx)
}

// ResolveNamedContext возвращает именованную зависимость типа T, пробрасывая
// ctx в фабрику.
func (c *Container) ResolveNamedContext[T any](ctx context.Context, name string) (T, error) {
	return c.root.ResolveNamedContext[T](ctx, name)
}

// MustResolve возвращает зависимость типа T и паникует, если разрешить ее не
// удалось.
func (c *Container) MustResolve[T any]() T {
	return c.root.MustResolve[T]()
}

// MustResolveNamed возвращает именованную зависимость типа T и паникует, если
// разрешить ее не удалось.
func (c *Container) MustResolveNamed[T any](name string) T {
	return c.root.MustResolveNamed[T](name)
}

// MustResolveContext возвращает зависимость типа T и паникует, если разрешить
// ее не удалось.
func (c *Container) MustResolveContext[T any](ctx context.Context) T {
	return c.root.MustResolveContext[T](ctx)
}

// ResolveAll возвращает все зарегистрированные зависимости типа T в порядке
// регистрации. Если таких зависимостей нет, возвращается пустой срез и nil.
func (c *Container) ResolveAll[T any]() ([]T, error) {
	return c.root.ResolveAll[T]()
}

// Has сообщает, зарегистрирована ли зависимость типа T без имени.
func (c *Container) Has[T any]() bool {
	_, ok := c.root.lookup(any(newTypeKey[T]("")))

	return ok
}

// HasNamed сообщает, зарегистрирована ли зависимость типа T под именем name.
func (c *Container) HasNamed[T any](name string) bool {
	_, ok := c.root.lookup(any(newTypeKey[T](name)))

	return ok
}

// Remove удаляет регистрацию зависимости типа T без имени.
func (c *Container) Remove[T any]() {
	c.root.data.remove(any(newTypeKey[T]("")))
}

// RemoveNamed удаляет регистрацию зависимости типа T под именем name.
func (c *Container) RemoveNamed[T any](name string) {
	c.root.data.remove(any(newTypeKey[T](name)))
}

func (c *Container) storeFactory[T any](key typeKey[T], k kind, factory func(ctx context.Context) (T, error)) {
	c.root.data.store(any(key), newFactoryRegistration(key.name, k, factory), false)
}

func (c *Container) registerAsFactory[I, T any](key typeKey[I], k kind, factory func() (T, error)) {
	if factory == nil {
		c.storeFactory(key, k, nil)

		return
	}

	c.storeFactory(key, k, func(context.Context) (I, error) {
		value, err := factory()
		if err != nil {
			return zeroOf[I](), err
		}

		return castAs[I](value), nil
	})
}
