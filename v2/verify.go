package di

import (
	"errors"
	"fmt"
)

// Verify обходит граф зависимостей и сообщает о незарегистрированных зависимостях и циклах.
func (c *Container) Verify() error {
	const (
		visiting = 1
		visited  = 2
	)

	data := c.root.data
	states := make(map[any]int)

	var errs []error

	var visit func(key any)
	visit = func(key any) {
		switch states[key] {
		case visited:
			return
		case visiting:
			errs = append(errs, fmt.Errorf("%w: %s", ErrCircularDependency, keyName(key)))

			return
		}

		states[key] = visiting

		raw, ok := data.load(key)
		if !ok {
			errs = append(errs, fmt.Errorf("%w: %s", ErrNotFound, keyName(key)))
			states[key] = visited

			return
		}

		if info, ok := raw.(registrationInfo); ok {
			for _, dep := range info.dependencies() {
				visit(dep)
			}
		}

		states[key] = visited
	}

	for _, raw := range data.snapshotOrder() {
		info, ok := raw.(registrationInfo)
		if !ok || info.dependencies() == nil {
			continue
		}

		visit(info.registrationKey())
	}

	return errors.Join(errs...)
}

// Dispose закрывает синглтоны, реализующие io.Closer, в обратном порядке регистрации (LIFO).
func (c *Container) Dispose() error {
	return disposeAll(c.root.data.snapshotOrder())
}

// newProvideRegistration создает singleton-регистрацию для [Container.Provide].
func newProvideRegistration[T any](deps []any) *registration[T] {
	key := newTypeKey[T]("")

	if deps == nil {
		deps = []any{}
	}

	return &registration[T]{
		key:  any(key),
		kind: kindSingleton,
		deps: deps,
	}
}
