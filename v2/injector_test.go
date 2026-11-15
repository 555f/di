package di_test

import (
	"errors"
	"testing"

	"github.com/555f/di/v2"
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

// TestInjector проверяет внедрение одной зависимости.
func TestInjector(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.Register(injDep1(1))

	injector := container.Injector(func(d1 injDep1) int {
		return int(d1)
	})

	got, err := injector()
	if err != nil {
		t.Fatalf("injector() error = %v, want nil", err)
	}

	if want := 1; got != want {
		t.Errorf("injector() = %d, want %d", got, want)
	}
}

// TestInjector2 проверяет внедрение двух зависимостей.
func TestInjector2(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.Register(injDep1(1))
	container.Register(injDep2(2))

	injector := container.Injector2(func(d1 injDep1, d2 injDep2) int {
		return int(d1) + int(d2)
	})

	got, err := injector()
	if err != nil {
		t.Fatalf("injector() error = %v, want nil", err)
	}

	if want := 3; got != want {
		t.Errorf("injector() = %d, want %d", got, want)
	}
}

// TestInjector3 проверяет внедрение трех зависимостей.
func TestInjector3(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.Register(injDep1(1))
	container.Register(injDep2(2))
	container.Register(injDep3(3))

	injector := container.Injector3(func(d1 injDep1, d2 injDep2, d3 injDep3) int {
		return int(d1) + int(d2) + int(d3)
	})

	got, err := injector()
	if err != nil {
		t.Fatalf("injector() error = %v, want nil", err)
	}

	if want := 6; got != want {
		t.Errorf("injector() = %d, want %d", got, want)
	}
}

// TestInjector4 проверяет внедрение четырех зависимостей.
func TestInjector4(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.Register(injDep1(1))
	container.Register(injDep2(2))
	container.Register(injDep3(3))
	container.Register(injDep4(4))

	injector := container.Injector4(func(d1 injDep1, d2 injDep2, d3 injDep3, d4 injDep4) int {
		return int(d1) + int(d2) + int(d3) + int(d4)
	})

	got, err := injector()
	if err != nil {
		t.Fatalf("injector() error = %v, want nil", err)
	}

	if want := 10; got != want {
		t.Errorf("injector() = %d, want %d", got, want)
	}
}

// TestInjector5 проверяет внедрение пяти зависимостей.
func TestInjector5(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.Register(injDep1(1))
	container.Register(injDep2(2))
	container.Register(injDep3(3))
	container.Register(injDep4(4))
	container.Register(injDep5(5))

	injector := container.Injector5(func(d1 injDep1, d2 injDep2, d3 injDep3, d4 injDep4, d5 injDep5) int {
		return int(d1) + int(d2) + int(d3) + int(d4) + int(d5)
	})

	got, err := injector()
	if err != nil {
		t.Fatalf("injector() error = %v, want nil", err)
	}

	if want := 15; got != want {
		t.Errorf("injector() = %d, want %d", got, want)
	}
}

// TestInjector6 проверяет внедрение шести зависимостей.
func TestInjector6(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.Register(injDep1(1))
	container.Register(injDep2(2))
	container.Register(injDep3(3))
	container.Register(injDep4(4))
	container.Register(injDep5(5))
	container.Register(injDep6(6))

	injector := container.Injector6(func(d1 injDep1, d2 injDep2, d3 injDep3, d4 injDep4, d5 injDep5, d6 injDep6) int {
		return int(d1) + int(d2) + int(d3) + int(d4) + int(d5) + int(d6)
	})

	got, err := injector()
	if err != nil {
		t.Fatalf("injector() error = %v, want nil", err)
	}

	if want := 21; got != want {
		t.Errorf("injector() = %d, want %d", got, want)
	}
}

// TestInjector7 проверяет внедрение семи зависимостей.
func TestInjector7(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.Register(injDep1(1))
	container.Register(injDep2(2))
	container.Register(injDep3(3))
	container.Register(injDep4(4))
	container.Register(injDep5(5))
	container.Register(injDep6(6))
	container.Register(injDep7(7))

	injector := container.Injector7(func(d1 injDep1, d2 injDep2, d3 injDep3, d4 injDep4, d5 injDep5, d6 injDep6, d7 injDep7) int {
		return int(d1) + int(d2) + int(d3) + int(d4) + int(d5) + int(d6) + int(d7)
	})

	got, err := injector()
	if err != nil {
		t.Fatalf("injector() error = %v, want nil", err)
	}

	if want := 28; got != want {
		t.Errorf("injector() = %d, want %d", got, want)
	}
}

// TestInjector8 проверяет внедрение восьми зависимостей.
func TestInjector8(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.Register(injDep1(1))
	container.Register(injDep2(2))
	container.Register(injDep3(3))
	container.Register(injDep4(4))
	container.Register(injDep5(5))
	container.Register(injDep6(6))
	container.Register(injDep7(7))
	container.Register(injDep8(8))

	injector := container.Injector8(func(d1 injDep1, d2 injDep2, d3 injDep3, d4 injDep4, d5 injDep5, d6 injDep6, d7 injDep7, d8 injDep8) int {
		return int(d1) + int(d2) + int(d3) + int(d4) + int(d5) + int(d6) + int(d7) + int(d8)
	})

	got, err := injector()
	if err != nil {
		t.Fatalf("injector() error = %v, want nil", err)
	}

	if want := 36; got != want {
		t.Errorf("injector() = %d, want %d", got, want)
	}
}

// TestInjector9 проверяет внедрение девяти зависимостей.
func TestInjector9(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.Register(injDep1(1))
	container.Register(injDep2(2))
	container.Register(injDep3(3))
	container.Register(injDep4(4))
	container.Register(injDep5(5))
	container.Register(injDep6(6))
	container.Register(injDep7(7))
	container.Register(injDep8(8))
	container.Register(injDep9(9))

	injector := container.Injector9(func(d1 injDep1, d2 injDep2, d3 injDep3, d4 injDep4, d5 injDep5, d6 injDep6, d7 injDep7, d8 injDep8, d9 injDep9) int {
		return int(d1) + int(d2) + int(d3) + int(d4) + int(d5) + int(d6) + int(d7) + int(d8) + int(d9)
	})

	got, err := injector()
	if err != nil {
		t.Fatalf("injector() error = %v, want nil", err)
	}

	if want := 45; got != want {
		t.Errorf("injector() = %d, want %d", got, want)
	}
}

// TestInjector10 проверяет внедрение десяти зависимостей.
func TestInjector10(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.Register(injDep1(1))
	container.Register(injDep2(2))
	container.Register(injDep3(3))
	container.Register(injDep4(4))
	container.Register(injDep5(5))
	container.Register(injDep6(6))
	container.Register(injDep7(7))
	container.Register(injDep8(8))
	container.Register(injDep9(9))
	container.Register(injDep10(10))

	injector := container.Injector10(func(d1 injDep1, d2 injDep2, d3 injDep3, d4 injDep4, d5 injDep5, d6 injDep6, d7 injDep7, d8 injDep8, d9 injDep9, d10 injDep10) int {
		return int(d1) + int(d2) + int(d3) + int(d4) + int(d5) + int(d6) + int(d7) + int(d8) + int(d9) + int(d10)
	})

	got, err := injector()
	if err != nil {
		t.Fatalf("injector() error = %v, want nil", err)
	}

	if want := 55; got != want {
		t.Errorf("injector() = %d, want %d", got, want)
	}
}

// TestInjector_MissingDependency проверяет возврат ошибки, если зависимость не зарегистрирована.
func TestInjector_MissingDependency(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()

	called := false

	injector := container.Injector(func(d1 injDep1) *injResult {
		called = true

		return &injResult{total: int(d1)}
	})

	got, err := injector()
	if err == nil {
		t.Fatal("injector() error = nil, want error")
	}

	if called {
		t.Error("constructor called, want it to be skipped when dependency is not registered")
	}

	if got != nil {
		t.Errorf("injector() = %v, want nil", got)
	}
}

// TestInjector_MissingMiddleDependency проверяет, что ошибка второй зависимости прерывает внедрение.
func TestInjector_MissingMiddleDependency(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.Register(injDep1(1))
	container.Register(injDep3(3))

	called := false

	injector := container.Injector3(func(d1 injDep1, d2 injDep2, d3 injDep3) *injResult {
		called = true

		return &injResult{total: int(d1) + int(d2) + int(d3)}
	})

	got, err := injector()
	if err == nil {
		t.Fatal("injector() error = nil, want error")
	}

	if called {
		t.Error("constructor called, want it to be skipped when injDep2 is not registered")
	}

	if got != nil {
		t.Errorf("injector() = %v, want nil", got)
	}
}

// TestInjector_FactoryError проверяет проброс ошибки фабрики зависимости.
func TestInjector_FactoryError(t *testing.T) {
	t.Parallel()

	wantErr := errors.New("factory failed")

	container := di.NewContainer()
	container.RegisterFactory(func() (injDep1, error) {
		return 0, wantErr
	})

	called := false

	injector := container.Injector(func(d1 injDep1) *injResult {
		called = true

		return &injResult{total: int(d1)}
	})

	got, err := injector()
	if !errors.Is(err, wantErr) {
		t.Fatalf("injector() error = %v, want %v", err, wantErr)
	}

	if called {
		t.Error("constructor called, want it to be skipped when dependency factory fails")
	}

	if got != nil {
		t.Errorf("injector() = %v, want nil", got)
	}
}

// TestInjector_TransientDependency проверяет, что транзиентная зависимость создается на каждый вызов.
func TestInjector_TransientDependency(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()

	counter := 0

	container.RegisterTransient(func() (injDep1, error) {
		counter++

		return injDep1(counter), nil
	})

	injector := container.Injector(func(d1 injDep1) int {
		return int(d1)
	})

	for want := 1; want <= 2; want++ {
		got, err := injector()
		if err != nil {
			t.Fatalf("injector() call %d error = %v, want nil", want, err)
		}

		if got != want {
			t.Errorf("injector() call %d = %d, want %d", want, got, want)
		}
	}

	if want := 2; counter != want {
		t.Errorf("transient factory called %d times, want %d", counter, want)
	}
}

// TestInjector_LazyResolution проверяет, что зависимости разрешаются в момент вызова, а не создания инжектора.
func TestInjector_LazyResolution(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()

	injector := container.Injector(func(d1 injDep1) int {
		return int(d1)
	})

	container.Register(injDep1(7))

	got, err := injector()
	if err != nil {
		t.Fatalf("injector() error = %v, want nil", err)
	}

	if want := 7; got != want {
		t.Errorf("injector() = %d, want %d", got, want)
	}
}

// TestInjector_BoundToContainer проверяет, что инжектор разрешает зависимости из контейнера, для которого создан.
func TestInjector_BoundToContainer(t *testing.T) {
	t.Parallel()

	other := di.NewContainer()
	other.Register(injDep1(99))

	container := di.NewContainer()
	container.Register(injDep1(1))

	injector := container.Injector(func(d1 injDep1) int {
		return int(d1)
	})

	got, err := injector()
	if err != nil {
		t.Fatalf("injector() error = %v, want nil", err)
	}

	if want := 1; got != want {
		t.Errorf("injector() = %d, want %d", got, want)
	}
}

// TestInjector_IgnoresNamedRegistration проверяет, что инжектор использует только безымянную регистрацию.
func TestInjector_IgnoresNamedRegistration(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.RegisterNamed("named", injDep1(1))

	injector := container.Injector(func(d1 injDep1) int {
		return int(d1)
	})

	_, err := injector()
	if err == nil {
		t.Fatal("injector() error = nil, want error for named-only registration")
	}
}
