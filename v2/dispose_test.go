package di_test

import (
	"errors"
	"testing"

	"github.com/555f/di/v2"
)

// TestDispose_LIFOAndOnce проверяет закрытие в обратном порядке и ровно один раз.
func TestDispose_LIFOAndOnce(t *testing.T) {
	t.Parallel()

	var events []string

	first := newCloserSpy("first", &events)
	second := newCloserSpy("second", &events)
	third := newCloserSpy("third", &events)

	container := di.NewContainer()
	container.RegisterNamed("first", first)
	container.RegisterNamed("second", second)
	container.RegisterNamed("third", third)

	for _, name := range []string{"first", "second", "third"} {
		if _, err := container.ResolveNamed[*closerSpy](name); err != nil {
			t.Fatalf("ResolveNamed(%q) error = %v", name, err)
		}
	}

	if err := container.Dispose(); err != nil {
		t.Fatalf("Dispose() error = %v, want nil", err)
	}

	if err := container.Dispose(); err != nil {
		t.Fatalf("second Dispose() error = %v, want nil", err)
	}

	want := []string{"third", "second", "first"}
	if len(events) != len(want) {
		t.Fatalf("close order = %v, want %v", events, want)
	}

	for i, name := range want {
		if events[i] != name {
			t.Fatalf("close order = %v, want %v", events, want)
		}
	}

	for _, spy := range []*closerSpy{first, second, third} {
		if got := spy.closeCount(); got != 1 {
			t.Errorf("%s closed %d times, want 1", spy.name, got)
		}
	}
}

// TestDispose_JoinsErrors проверяет объединение ошибок через errors.Join.
func TestDispose_JoinsErrors(t *testing.T) {
	t.Parallel()

	errFirst := errors.New("first failed")
	errSecond := errors.New("second failed")

	first := newCloserSpy("first", nil)
	first.err = errFirst

	second := newCloserSpy("second", nil)
	second.err = errSecond

	container := di.NewContainer()
	container.RegisterNamed("first", first)
	container.RegisterNamed("second", second)

	if _, err := container.ResolveNamed[*closerSpy]("first"); err != nil {
		t.Fatalf("ResolveNamed() error = %v", err)
	}

	if _, err := container.ResolveNamed[*closerSpy]("second"); err != nil {
		t.Fatalf("ResolveNamed() error = %v", err)
	}

	err := container.Dispose()
	if !errors.Is(err, errFirst) || !errors.Is(err, errSecond) {
		t.Fatalf("Dispose() error = %v, want both errors joined", err)
	}
}

// TestDispose_IgnoresTransientsAndUnresolved проверяет, что транзиенты не
// закрываются, а неразрешенные синглтоны ничего не закрывают.
func TestDispose_IgnoresTransientsAndUnresolved(t *testing.T) {
	t.Parallel()

	var events []string

	transient := newCloserSpy("transient", &events)
	unresolved := newCloserSpy("unresolved", &events)

	container := di.NewContainer()
	container.RegisterTransientNamed("transient", func() (*closerSpy, error) {
		return transient, nil
	})
	container.RegisterFactoryNamed("singleton", func() (*closerSpy, error) {
		return unresolved, nil
	})

	if _, err := container.ResolveNamed[*closerSpy]("transient"); err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}

	if err := container.Dispose(); err != nil {
		t.Fatalf("Dispose() error = %v, want nil", err)
	}

	if transient.isClosed() {
		t.Error("transient was closed, want it to be left to the caller")
	}

	if unresolved.isClosed() {
		t.Error("unresolved singleton was closed, want nothing to close")
	}
}

// TestDispose_SingletonFactory проверяет закрытие синглтона, созданного фабрикой.
func TestDispose_SingletonFactory(t *testing.T) {
	t.Parallel()

	spy := newCloserSpy("factory", nil)

	container := di.NewContainer()
	container.RegisterFactory(func() (*closerSpy, error) {
		return spy, nil
	})

	if _, err := container.Resolve[*closerSpy](); err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}

	if err := container.Dispose(); err != nil {
		t.Fatalf("Dispose() error = %v", err)
	}

	if !spy.isClosed() {
		t.Error("singleton created by factory was not closed")
	}
}
