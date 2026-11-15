package di_test

import (
	"errors"
	"testing"

	"github.com/555f/di"
)

type (
	injDep1  int
	injDep2  int
	injDep3  int
	injDep4  int
	injDep5  int
	injDep6  int
	injDep7  int
	injDep8  int
	injDep9  int
	injDep10 int
)

type injResult struct {
	total int
}

// TestNewInjector проверяет внедрение одной зависимости.
func TestNewInjector(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	di.Register(container, injDep1(1))

	injector := di.NewInjector(func(d1 injDep1) int {
		return int(d1)
	})

	got, err := injector(container)
	if err != nil {
		t.Fatalf("injector(container) error = %v, want nil", err)
	}

	if want := 1; got != want {
		t.Errorf("injector(container) = %d, want %d", got, want)
	}
}

// TestNewInjector2 проверяет внедрение двух зависимостей.
func TestNewInjector2(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	di.Register(container, injDep1(1))
	di.Register(container, injDep2(2))

	injector := di.NewInjector2(func(d1 injDep1, d2 injDep2) int {
		return int(d1) + int(d2)
	})

	got, err := injector(container)
	if err != nil {
		t.Fatalf("injector(container) error = %v, want nil", err)
	}

	if want := 3; got != want {
		t.Errorf("injector(container) = %d, want %d", got, want)
	}
}

// TestNewInjector3 проверяет внедрение трех зависимостей.
func TestNewInjector3(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	di.Register(container, injDep1(1))
	di.Register(container, injDep2(2))
	di.Register(container, injDep3(3))

	injector := di.NewInjector3(func(d1 injDep1, d2 injDep2, d3 injDep3) int {
		return int(d1) + int(d2) + int(d3)
	})

	got, err := injector(container)
	if err != nil {
		t.Fatalf("injector(container) error = %v, want nil", err)
	}

	if want := 6; got != want {
		t.Errorf("injector(container) = %d, want %d", got, want)
	}
}

// TestNewInjector4 проверяет внедрение четырех зависимостей.
func TestNewInjector4(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	di.Register(container, injDep1(1))
	di.Register(container, injDep2(2))
	di.Register(container, injDep3(3))
	di.Register(container, injDep4(4))

	injector := di.NewInjector4(func(d1 injDep1, d2 injDep2, d3 injDep3, d4 injDep4) int {
		return int(d1) + int(d2) + int(d3) + int(d4)
	})

	got, err := injector(container)
	if err != nil {
		t.Fatalf("injector(container) error = %v, want nil", err)
	}

	if want := 10; got != want {
		t.Errorf("injector(container) = %d, want %d", got, want)
	}
}

// TestNewInjector5 проверяет внедрение пяти зависимостей.
func TestNewInjector5(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	di.Register(container, injDep1(1))
	di.Register(container, injDep2(2))
	di.Register(container, injDep3(3))
	di.Register(container, injDep4(4))
	di.Register(container, injDep5(5))

	injector := di.NewInjector5(func(d1 injDep1, d2 injDep2, d3 injDep3, d4 injDep4, d5 injDep5) int {
		return int(d1) + int(d2) + int(d3) + int(d4) + int(d5)
	})

	got, err := injector(container)
	if err != nil {
		t.Fatalf("injector(container) error = %v, want nil", err)
	}

	if want := 15; got != want {
		t.Errorf("injector(container) = %d, want %d", got, want)
	}
}

// TestNewInjector6 проверяет внедрение шести зависимостей.
func TestNewInjector6(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	di.Register(container, injDep1(1))
	di.Register(container, injDep2(2))
	di.Register(container, injDep3(3))
	di.Register(container, injDep4(4))
	di.Register(container, injDep5(5))
	di.Register(container, injDep6(6))

	injector := di.NewInjector6(func(d1 injDep1, d2 injDep2, d3 injDep3, d4 injDep4, d5 injDep5, d6 injDep6) int {
		return int(d1) + int(d2) + int(d3) + int(d4) + int(d5) + int(d6)
	})

	got, err := injector(container)
	if err != nil {
		t.Fatalf("injector(container) error = %v, want nil", err)
	}

	if want := 21; got != want {
		t.Errorf("injector(container) = %d, want %d", got, want)
	}
}

// TestNewInjector7 проверяет внедрение семи зависимостей.
func TestNewInjector7(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	di.Register(container, injDep1(1))
	di.Register(container, injDep2(2))
	di.Register(container, injDep3(3))
	di.Register(container, injDep4(4))
	di.Register(container, injDep5(5))
	di.Register(container, injDep6(6))
	di.Register(container, injDep7(7))

	injector := di.NewInjector7(func(d1 injDep1, d2 injDep2, d3 injDep3, d4 injDep4, d5 injDep5, d6 injDep6, d7 injDep7) int {
		return int(d1) + int(d2) + int(d3) + int(d4) + int(d5) + int(d6) + int(d7)
	})

	got, err := injector(container)
	if err != nil {
		t.Fatalf("injector(container) error = %v, want nil", err)
	}

	if want := 28; got != want {
		t.Errorf("injector(container) = %d, want %d", got, want)
	}
}

// TestNewInjector8 проверяет внедрение восьми зависимостей.
func TestNewInjector8(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	di.Register(container, injDep1(1))
	di.Register(container, injDep2(2))
	di.Register(container, injDep3(3))
	di.Register(container, injDep4(4))
	di.Register(container, injDep5(5))
	di.Register(container, injDep6(6))
	di.Register(container, injDep7(7))
	di.Register(container, injDep8(8))

	injector := di.NewInjector8(func(d1 injDep1, d2 injDep2, d3 injDep3, d4 injDep4, d5 injDep5, d6 injDep6, d7 injDep7, d8 injDep8) int {
		return int(d1) + int(d2) + int(d3) + int(d4) + int(d5) + int(d6) + int(d7) + int(d8)
	})

	got, err := injector(container)
	if err != nil {
		t.Fatalf("injector(container) error = %v, want nil", err)
	}

	if want := 36; got != want {
		t.Errorf("injector(container) = %d, want %d", got, want)
	}
}

// TestNewInjector9 проверяет внедрение девяти зависимостей.
func TestNewInjector9(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	di.Register(container, injDep1(1))
	di.Register(container, injDep2(2))
	di.Register(container, injDep3(3))
	di.Register(container, injDep4(4))
	di.Register(container, injDep5(5))
	di.Register(container, injDep6(6))
	di.Register(container, injDep7(7))
	di.Register(container, injDep8(8))
	di.Register(container, injDep9(9))

	injector := di.NewInjector9(func(d1 injDep1, d2 injDep2, d3 injDep3, d4 injDep4, d5 injDep5, d6 injDep6, d7 injDep7, d8 injDep8, d9 injDep9) int {
		return int(d1) + int(d2) + int(d3) + int(d4) + int(d5) + int(d6) + int(d7) + int(d8) + int(d9)
	})

	got, err := injector(container)
	if err != nil {
		t.Fatalf("injector(container) error = %v, want nil", err)
	}

	if want := 45; got != want {
		t.Errorf("injector(container) = %d, want %d", got, want)
	}
}

// TestNewInjector10 проверяет внедрение десяти зависимостей.
func TestNewInjector10(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	di.Register(container, injDep1(1))
	di.Register(container, injDep2(2))
	di.Register(container, injDep3(3))
	di.Register(container, injDep4(4))
	di.Register(container, injDep5(5))
	di.Register(container, injDep6(6))
	di.Register(container, injDep7(7))
	di.Register(container, injDep8(8))
	di.Register(container, injDep9(9))
	di.Register(container, injDep10(10))

	injector := di.NewInjector10(func(d1 injDep1, d2 injDep2, d3 injDep3, d4 injDep4, d5 injDep5, d6 injDep6, d7 injDep7, d8 injDep8, d9 injDep9, d10 injDep10) int {
		return int(d1) + int(d2) + int(d3) + int(d4) + int(d5) + int(d6) + int(d7) + int(d8) + int(d9) + int(d10)
	})

	got, err := injector(container)
	if err != nil {
		t.Fatalf("injector(container) error = %v, want nil", err)
	}

	if want := 55; got != want {
		t.Errorf("injector(container) = %d, want %d", got, want)
	}
}

// TestNewInjector_MissingDependency проверяет возврат ошибки, если зависимость не зарегистрирована.
func TestNewInjector_MissingDependency(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()

	called := false

	injector := di.NewInjector(func(d1 injDep1) *injResult {
		called = true

		return &injResult{total: int(d1)}
	})

	got, err := injector(container)
	if err == nil {
		t.Fatal("injector(container) error = nil, want error")
	}

	if called {
		t.Error("constructor called, want it to be skipped when dependency is not registered")
	}

	if got != nil {
		t.Errorf("injector(container) = %v, want nil", got)
	}
}

// TestNewInjector_MissingMiddleDependency проверяет, что ошибка второй зависимости прерывает внедрение.
func TestNewInjector_MissingMiddleDependency(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	di.Register(container, injDep1(1))
	di.Register(container, injDep3(3))

	called := false

	injector := di.NewInjector3(func(d1 injDep1, d2 injDep2, d3 injDep3) *injResult {
		called = true

		return &injResult{total: int(d1) + int(d2) + int(d3)}
	})

	got, err := injector(container)
	if err == nil {
		t.Fatal("injector(container) error = nil, want error")
	}

	if called {
		t.Error("constructor called, want it to be skipped when injDep2 is not registered")
	}

	if got != nil {
		t.Errorf("injector(container) = %v, want nil", got)
	}
}

// TestNewInjector_FactoryError проверяет проброс ошибки фабрики зависимости.
func TestNewInjector_FactoryError(t *testing.T) {
	t.Parallel()

	wantErr := errors.New("factory failed")

	container := di.NewContainer()
	di.RegisterFactory(container, func() (injDep1, error) {
		return 0, wantErr
	})

	called := false

	injector := di.NewInjector(func(d1 injDep1) *injResult {
		called = true

		return &injResult{total: int(d1)}
	})

	got, err := injector(container)
	if !errors.Is(err, wantErr) {
		t.Fatalf("injector(container) error = %v, want %v", err, wantErr)
	}

	if called {
		t.Error("constructor called, want it to be skipped when dependency factory fails")
	}

	if got != nil {
		t.Errorf("injector(container) = %v, want nil", got)
	}
}

// TestNewInjector_TransientDependency проверяет, что транзиентная зависимость создается на каждый вызов.
func TestNewInjector_TransientDependency(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()

	counter := 0

	di.RegisterTransient(container, func() (injDep1, error) {
		counter++

		return injDep1(counter), nil
	})

	injector := di.NewInjector(func(d1 injDep1) int {
		return int(d1)
	})

	for want := 1; want <= 2; want++ {
		got, err := injector(container)
		if err != nil {
			t.Fatalf("injector(container) call %d error = %v, want nil", want, err)
		}

		if got != want {
			t.Errorf("injector(container) call %d = %d, want %d", want, got, want)
		}
	}

	if want := 2; counter != want {
		t.Errorf("transient factory called %d times, want %d", counter, want)
	}
}

// TestNewInjector_LazyResolution проверяет, что зависимости разрешаются в момент вызова, а не создания инжектора.
func TestNewInjector_LazyResolution(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()

	injector := di.NewInjector(func(d1 injDep1) int {
		return int(d1)
	})

	di.Register(container, injDep1(7))

	got, err := injector(container)
	if err != nil {
		t.Fatalf("injector(container) error = %v, want nil", err)
	}

	if want := 7; got != want {
		t.Errorf("injector(container) = %d, want %d", got, want)
	}
}

// TestNewInjector_IgnoresNamedRegistration проверяет, что инжектор использует только безымянную регистрацию.
func TestNewInjector_IgnoresNamedRegistration(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	di.RegisterInstanceNamed(container, "named", injDep1(1))

	injector := di.NewInjector(func(d1 injDep1) int {
		return int(d1)
	})

	_, err := injector(container)
	if err == nil {
		t.Fatal("injector(container) error = nil, want error for named-only registration")
	}
}

// TestNewInjector_DifferentContainers проверяет, что один инжектор разрешает зависимости из переданного контейнера.
func TestNewInjector_DifferentContainers(t *testing.T) {
	t.Parallel()

	first := di.NewContainer()
	di.Register(first, injDep1(1))

	second := di.NewContainer()
	di.Register(second, injDep1(2))

	injector := di.NewInjector(func(d1 injDep1) int {
		return int(d1)
	})

	for _, tt := range []struct {
		name      string
		container *di.Container
		want      int
	}{
		{"first container", first, 1},
		{"second container", second, 2},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := injector(tt.container)
			if err != nil {
				t.Fatalf("injector(container) error = %v, want nil", err)
			}

			if got != tt.want {
				t.Errorf("injector(container) = %d, want %d", got, tt.want)
			}
		})
	}
}
