package di_test

import (
	"errors"
	"testing"

	"github.com/555f/di/v2"
)

type composed string

type total int

type (
	d1  int
	d2  int
	d3  int
	d4  int
	d5  int
	d6  int
	d7  int
	d8  int
	d9  int
	d10 int
)

// TestProvide_EqualsFactoryAndInjector проверяет, что Provide разрешает
// зависимости так же, как фабрика с инжектором.
func TestProvide_EqualsFactoryAndInjector(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.Register(1)
	container.Register("hello")
	container.Register(true)

	container.Provide3(func(n int, s string, b bool) (composed, error) {
		if !b {
			return "", errors.New("flag is false")
		}

		return composed(s + " " + string(rune('0'+n))), nil
	})

	got, err := container.Resolve[composed]()
	if err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}

	if got != "hello 1" {
		t.Errorf("Resolve() = %q, want %q", got, "hello 1")
	}
}

// TestProvide_MissingDependency проверяет ошибку при пропущенной зависимости.
func TestProvide_MissingDependency(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.Provide2(func(n int, b bool) (string, error) {
		return "", nil
	})

	_, err := container.Resolve[string]()
	if !errors.Is(err, di.ErrNotFound) {
		t.Fatalf("Resolve() error = %v, want ErrNotFound", err)
	}
}

// TestProvide10 проверяет Provide на максимальное кол-во параметров.
func TestProvide10(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.Register(d1(1))
	container.Register(d2(2))
	container.Register(d3(3))
	container.Register(d4(4))
	container.Register(d5(5))
	container.Register(d6(6))
	container.Register(d7(7))
	container.Register(d8(8))
	container.Register(d9(9))
	container.Register(d10(10))

	container.Provide10(
		func(a d1, b d2, c d3, d d4, e d5, f d6, g d7, h d8, i d9, j d10) (total, error) {
			return total(int(a) + int(b) + int(c) + int(d) + int(e) + int(f) + int(g) + int(h) + int(i) + int(j)), nil
		},
	)

	got, err := container.Resolve[total]()
	if err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}

	if got != 55 {
		t.Errorf("Resolve() = %d, want 55", got)
	}
}

// TestRegisterAs проверяет привязку конкретного типа к интерфейсу.
func TestRegisterAs(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.RegisterAs[greeter](simpleGreeter{msg: "hello"})

	got, err := container.Resolve[greeter]()
	if err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}

	if got.Greet() != "hello" {
		t.Errorf("Greet() = %q, want hello", got.Greet())
	}
}

// TestRegisterFactoryAs проверяет привязку фабрики к интерфейсу.
func TestRegisterFactoryAs(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.RegisterFactoryAs[greeter](func() (simpleGreeter, error) {
		return simpleGreeter{msg: "factory"}, nil
	})

	got := container.MustResolve[greeter]()
	if got.Greet() != "factory" {
		t.Errorf("Greet() = %q, want factory", got.Greet())
	}
}

// TestRegisterTransientAs проверяет транзиентную привязку к интерфейсу и то,
// что интерфейс и конкретный тип - разные ключи.
func TestRegisterTransientAs(t *testing.T) {
	t.Parallel()

	counter := 0

	container := di.NewContainer()
	container.RegisterTransientAs[greeter](func() (simpleGreeter, error) {
		counter++

		return simpleGreeter{msg: "x"}, nil
	})

	container.MustResolve[greeter]()
	container.MustResolve[greeter]()

	if counter != 2 {
		t.Errorf("transient factory called %d times, want 2", counter)
	}

	if container.Has[simpleGreeter]() {
		t.Error("concrete type must not be registered by RegisterTransientAs")
	}
}

// TestResolveAll_OrderAndEmpty проверяет сохранение порядка и пустой результат.
func TestResolveAll_OrderAndEmpty(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.Register(0)
	container.RegisterMany(1, 2)
	container.RegisterMany(3)

	got, err := container.ResolveAll[int]()
	if err != nil {
		t.Fatalf("ResolveAll() error = %v", err)
	}

	want := []int{0, 1, 2, 3}
	if len(got) != len(want) {
		t.Fatalf("ResolveAll() = %v, want %v", got, want)
	}

	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("ResolveAll() = %v, want %v", got, want)
		}
	}

	empty, err := container.ResolveAll[float64]()
	if err != nil {
		t.Fatalf("ResolveAll() error = %v, want nil", err)
	}

	if len(empty) != 0 {
		t.Errorf("ResolveAll() = %v, want empty slice", empty)
	}
}

// TestHasRemove проверяет Has, HasNamed, Remove и RemoveNamed.
func TestHasRemove(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.Register(1)
	container.RegisterNamed("named", "value")

	if !container.Has[int]() {
		t.Error("Has[int]() = false, want true")
	}

	if !container.HasNamed[string]("named") {
		t.Error("HasNamed[string](named) = false, want true")
	}

	if container.Has[bool]() {
		t.Error("Has[bool]() = true, want false")
	}

	container.Remove[int]()
	if container.Has[int]() {
		t.Error("Has[int]() = true after Remove, want false")
	}

	container.RemoveNamed[string]("named")
	if container.HasNamed[string]("named") {
		t.Error("HasNamed() = true after RemoveNamed, want false")
	}
}

// TestMustResolveNamed проверяет именованное разрешение и панику при промахе.
func TestMustResolveNamed(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.RegisterNamed("greeting", "hi")

	if got := container.MustResolveNamed[string]("greeting"); got != "hi" {
		t.Errorf("MustResolveNamed() = %q, want hi", got)
	}

	defer func() {
		if recover() == nil {
			t.Error("MustResolveNamed() did not panic for missing name")
		}
	}()

	container.MustResolveNamed[string]("missing")
}
