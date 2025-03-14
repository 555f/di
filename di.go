package di

import (
	"fmt"
	"sync"
)

type Containerer interface{}

// Container — контейнер для управления зависимостями.
type Container struct {
	dependencies sync.Map
	singletons   sync.Map
}

// RegisterFactoryOptions опции для регистрации фабричной функции.
type RegisterFactoryOptions struct {
	singleton bool // создавать только один экземпляр зависимости
}

type RegisterFactoryOption func(*RegisterFactoryOptions)

// WithSingleton опция для регистрации фабричной функции как синглтон.
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
	var zero T

	register(c, typeName(zero), value, false)
}

// Register регистрирует именованную зависимость как готовый экземпляр.
func RegisterNamed[T any](c *Container, name string, value T) {
	register(c, name, value, false)
}

// RegisterFactory регистрирует именованную фабричную функцию для создания зависимости.
func RegisterFactory[T any](c *Container, factory func() (T, error), opts ...RegisterFactoryOption) {
	var zero T

	RegisterFactoryaNamed(c, typeName(zero), factory, opts...)
}

// RegisterFactory регистрирует фабричную функцию для создания зависимости.
func RegisterFactoryaNamed[T any](c *Container, name string, factory func() (T, error), opts ...RegisterFactoryOption) {
	o := &RegisterFactoryOptions{}
	for _, applyOpt := range opts {
		applyOpt(o)
	}

	register(c, name, factory, o.singleton)
}

func ResolveNamed[T any](c *Container, name string) (T, error) {
	return resolve[T](c, name)
}

func Resolve[T any](c *Container) (T, error) {
	var zero T

	name := typeName(zero)

	return resolve[T](c, name)
}

func register[T any](c *Container, name string, provider T, singleton bool) {
	if singleton {
		c.singletons.Store(name, provider)
	} else {
		c.dependencies.Store(name, provider)
	}
}

// Resolve извлекает зависимость из контейнера.
func resolve[T any](c *Container, name string) (T, error) {
	var zero T

	if factory, ok := c.singletons.Load(name); ok {
		if singletonFactory, isFunc := factory.(func() (T, error)); isFunc {
			value, err := singletonFactory()
			if err != nil {
				return zero, err
			}
			c.dependencies.Store(name, value)
			c.singletons.Delete(name)
			return value, nil
		}
	}

	value, ok := c.dependencies.Load(name)
	if !ok {
		return zero, fmt.Errorf("dpendency %s not registered", name)
	}

	if factory, isFunc := value.(func() (T, error)); isFunc {
		return factory()
	}

	return value.(T), nil
}
