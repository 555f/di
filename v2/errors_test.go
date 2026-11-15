package di_test

import (
	"errors"
	"testing"
	"time"

	"github.com/555f/di/v2"
)

// TestErrors_NotFound проверяет errors.Is для незарегистрированной зависимости.
func TestErrors_NotFound(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()

	_, err := container.Resolve[string]()
	if !errors.Is(err, di.ErrNotFound) {
		t.Fatalf("Resolve() error = %v, want ErrNotFound", err)
	}
}

// TestErrors_NotFoundNamed проверяет errors.Is для незарегистрированного имени.
func TestErrors_NotFoundNamed(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()

	_, err := container.ResolveNamed[string]("missing")
	if !errors.Is(err, di.ErrNotFound) {
		t.Fatalf("ResolveNamed() error = %v, want ErrNotFound", err)
	}
}

// TestErrors_NoFactory проверяет, что nil-фабрика дает ErrNoFactory.
func TestErrors_NoFactory(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.RegisterFactory[string](nil)

	_, err := container.Resolve[string]()
	if !errors.Is(err, di.ErrNoFactory) {
		t.Fatalf("Resolve() error = %v, want ErrNoFactory", err)
	}

	container.RegisterTransientNamed[int]("t", nil)

	if _, err := container.ResolveNamed[int]("t"); !errors.Is(err, di.ErrNoFactory) {
		t.Fatalf("ResolveNamed() error = %v, want ErrNoFactory", err)
	}
}

// TestErrors_CircularDependency проверяет, что цикл Provide(A->B->A) возвращает
// ErrCircularDependency и не зависает на мьютексе.
func TestErrors_CircularDependency(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.Register(0)
	container.Provide2(func(b *svcB, n int) (*svcA, error) {
		return &svcA{b: b}, nil
	})
	container.Provide2(func(a *svcA, n int) (*svcB, error) {
		return &svcB{a: a}, nil
	})

	done := make(chan error, 1)

	go func() {
		_, err := container.Resolve[*svcA]()
		done <- err
	}()

	select {
	case err := <-done:
		if !errors.Is(err, di.ErrCircularDependency) {
			t.Fatalf("Resolve() error = %v, want ErrCircularDependency", err)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("Resolve() hung on a circular dependency")
	}
}

// TestErrors_CircularDependencySelf проверяет самоссылку.
func TestErrors_CircularDependencySelf(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.Provide2(func(a *svcA, n int) (*svcA, error) {
		return a, nil
	})

	done := make(chan error, 1)

	go func() {
		_, err := container.Resolve[*svcA]()
		done <- err
	}()

	select {
	case err := <-done:
		if !errors.Is(err, di.ErrCircularDependency) {
			t.Fatalf("Resolve() error = %v, want ErrCircularDependency", err)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("Resolve() hung on a self-referencing dependency")
	}
}
