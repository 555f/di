package di

import (
	"errors"
	"fmt"
	"sync"
)

type Kind int

const (
	Singleton Kind = iota // Создается один раз и кэшируется
	Transient             // Создается каждый раз заново
)

type Registration[T any] struct {
	factory  func() (T, error)
	instance *T
	kind     Kind
	mutex    sync.RWMutex
}

type Container struct {
	registrations sync.Map
	scopes        sync.Map
}

func NewContainer() *Container {
	return &Container{}
}

func RegisterFactory[T any](
	c *Container,
	factory func() (T, error),
) {
	key := newTypeKey[T]()

	registerFactory(c, key, factory)
}

func RegisterFactoryNamed[T any](c *Container, name string, factory func() (T, error)) {
	key := newTypeKey[T](name)

	registerFactory[T](c, key, factory)
}

func RegisterTransient[T any](c *Container, factory func() (T, error)) {
	key := newTypeKey[T]()

	registerTransient(c, key, factory)
}

func RegisterTransientNamed[T any](c *Container, name string, factory func() (T, error)) {
	key := newTypeKey[T](name)

	registerTransient(c, key, factory)
}

func Register[T any](c *Container, instance T) {
	key := newTypeKey[T]()

	registerInstance(c, key, instance)
}

func RegisterInstanceNamed[T any](c *Container, name string, instance T) {
	key := newTypeKey[T](name)

	registerInstance(c, key, instance)
}

func registerInstance[T any](
	c *Container,
	key TypeKey[T],
	instance T,
) {
	registration := &Registration[T]{
		instance: &instance,
		kind:     Singleton,
	}

	c.registrations.Store(key, registration)
}

func registerFactory[T any](
	c *Container,
	key TypeKey[T],
	factory func() (T, error),
) {
	registration := &Registration[T]{
		factory: factory,
		kind:    Singleton,
	}

	c.registrations.Store(key, registration)
}

func registerTransient[T any](
	c *Container,
	key TypeKey[T],
	factory func() (T, error),
) {
	registration := &Registration[T]{
		factory: factory,
		kind:    Transient,
	}

	c.registrations.Store(key, registration)
}

func Resolve[T any](c *Container) (T, error) {
	key := newTypeKey[T]()

	return resolve[T](c, key)
}

func ResolveNamed[T any](c *Container, name string) (T, error) {
	key := newTypeKey[T](name)

	return resolve[T](c, key)
}

func MustResolve[T any](c *Container) T {
	key := newTypeKey[T]()

	result, err := resolve[T](c, key)
	if err != nil {
		panic("failed to resolve dependency: " + err.Error())
	}

	return result
}

func resolve[T any](
	c *Container,
	key TypeKey[T],
) (T, error) {
	var zero T

	value, ok := c.registrations.Load(key)
	if !ok {
		return zero, fmt.Errorf("dependency not registered: %s", key.String())
	}

	registration, ok := value.(*Registration[T])
	if !ok {
		return zero, errors.New("invalid registration type")
	}

	switch registration.kind {
	case Singleton:
		return resolveSingleton(registration)
	case Transient:
		return resolveTransient(registration)
	default:
		return zero, errors.New("unsupported kind")
	}
}

func resolveSingleton[T any](registration *Registration[T]) (T, error) {
	var zero T

	registration.mutex.RLock()
	if registration.instance != nil {
		instance := *registration.instance
		registration.mutex.RUnlock()

		return instance, nil
	}
	registration.mutex.RUnlock()

	registration.mutex.Lock()
	defer registration.mutex.Unlock()

	if registration.instance != nil {
		return *registration.instance, nil
	}

	if registration.factory == nil {
		return zero, errors.New("no factory function provided")
	}

	instance, err := registration.factory()
	if err != nil {
		return zero, err
	}

	registration.instance = &instance

	return instance, nil
}

func resolveTransient[T any](registration *Registration[T]) (T, error) {
	var zero T

	if registration.factory == nil {
		return zero, errors.New("no factory function provided")
	}

	return registration.factory()
}
