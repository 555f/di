package di

import (
	"fmt"
	"sync"
)

// Container - контейнер для управления зависимостями.
type Container struct {
	dependencies sync.Map
	singletons   sync.Map
}

// RegisterFactoryOptions - опции для регистрации фабричной функции.
type RegisterFactoryOptions struct {
	singleton bool
}

type RegisterFactoryOption func(*RegisterFactoryOptions)

// WithSingleton - опция для регистрации фабричной функции как синглтон.
func WithSingleton() RegisterFactoryOption {
	return func(options *RegisterFactoryOptions) {
		options.singleton = true
	}
}

// NewContainer создает новый контейнер.
func NewContainer() *Container {
	return &Container{}
}

// Register регистрирует зависимость как готовый экземпляр.
func Register[T any](c *Container, value T) {
	registerInstance(c, typeName(*new(T)), value)
}

// RegisterNamed регистрирует именованную зависимость как готовый экземпляр.
func RegisterNamed[T any](c *Container, name string, value T) {
	registerInstance(c, buildName(name, *new(T)), value)
}

// RegisterFactory регистрирует фабричную функцию для создания зависимости.
func RegisterFactory[T any](c *Container, factory func() (T, error), opts ...RegisterFactoryOption) {
	registerFactory(c, "", factory, opts...)
}

// RegisterFactoryNamed регистрирует именованную фабричную функцию для создания зависимости.
func RegisterFactoryNamed[T any](c *Container, name string, factory func() (T, error), opts ...RegisterFactoryOption) {
	registerFactory(c, name, factory, opts...)
}

// Resolve извлекает зависимость из контейнера.
func Resolve[T any](c *Container) (T, error) {
	return resolve[T](c, typeName(*new(T)))
}

// ResolveNamed извлекает именованную зависимость из контейнера.
func ResolveNamed[T any](c *Container, name string) (T, error) {
	return resolve[T](c, buildName(name, *new(T)))
}

func registerInstance[T any](c *Container, name string, value T) {
	c.dependencies.Store(name, value)
}

func registerFactory[T any](c *Container, name string, factory func() (T, error), opts ...RegisterFactoryOption) {
	options := &RegisterFactoryOptions{}
	for _, opt := range opts {
		opt(options)
	}

	key := buildName(name, *new(T))
	if options.singleton {
		c.singletons.Store(key, factory)
	} else {
		c.dependencies.Store(key, factory)
	}
}

func resolve[T any](c *Container, name string) (T, error) {
	var zero T

	if singleton, ok := c.singletons.Load(name); ok {
		if factory, isFunc := singleton.(func() (T, error)); isFunc {
			instance, err := factory()
			if err != nil {
				return zero, err
			}
			c.dependencies.Store(name, instance)
			c.singletons.Delete(name)
			return instance, nil
		}
	}

	value, ok := c.dependencies.Load(name)
	if !ok {
		return zero, fmt.Errorf("dependency %s not registered", name)
	}

	if factory, isFunc := value.(func() (T, error)); isFunc {
		return factory()
	}

	return value.(T), nil
}

func buildName[T any](name string, _ T) string {
	return name + typeName(*new(T))
}
