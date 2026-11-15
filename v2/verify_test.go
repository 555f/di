package di_test

import (
	"errors"
	"testing"

	"github.com/555f/di/v2"
)

// TestVerify_OK проверяет успешный граф зависимостей.
func TestVerify_OK(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.Register(true)
	container.Provide(func() (int, error) { return 1, nil })
	container.Provide2(func(n int, b bool) (string, error) {
		return "ok", nil
	})

	if err := container.Verify(); err != nil {
		t.Fatalf("Verify() error = %v, want nil", err)
	}
}

// TestVerify_Missing проверяет пропущенную зависимость.
func TestVerify_Missing(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.Provide2(func(n int, b bool) (string, error) {
		return "ok", nil
	})

	err := container.Verify()
	if !errors.Is(err, di.ErrNotFound) {
		t.Fatalf("Verify() error = %v, want ErrNotFound", err)
	}
}

// TestVerify_Cycle проверяет цикл без разрешения зависимостей.
func TestVerify_Cycle(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.Provide2(func(b *svcB, n int) (*svcA, error) {
		return &svcA{b: b}, nil
	})
	container.Provide2(func(a *svcA, n int) (*svcB, error) {
		return &svcB{a: a}, nil
	})

	err := container.Verify()
	if !errors.Is(err, di.ErrCircularDependency) {
		t.Fatalf("Verify() error = %v, want ErrCircularDependency", err)
	}
}

// TestVerify_IgnoresManualFactories проверяет, что регистрации с ручным
// Resolve не ломают проверку.
func TestVerify_IgnoresManualFactories(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.RegisterFactory(func() (int, error) {
		value, err := container.Resolve[string]()
		if err != nil {
			return 0, err
		}

		return len(value), nil
	})

	if err := container.Verify(); err != nil {
		t.Fatalf("Verify() error = %v, want nil", err)
	}
}
