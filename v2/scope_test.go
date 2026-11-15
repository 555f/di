package di_test

import (
	"context"
	"errors"
	"testing"

	"github.com/555f/di/v2"
)

// TestScope_ScopedCachedPerScope проверяет кэширование scoped внутри скоупа и
// разделение экземпляров между скоупами.
func TestScope_ScopedCachedPerScope(t *testing.T) {
	t.Parallel()

	counter := 0

	container := di.NewContainer()
	container.RegisterScoped(func() (*closerSpy, error) {
		counter++

		return newCloserSpy("scoped", nil), nil
	})

	first := container.BeginScope()
	second := container.BeginScope()

	firstA := first.MustResolve[*closerSpy]()
	firstB := first.MustResolve[*closerSpy]()
	secondA := second.MustResolve[*closerSpy]()

	if firstA != firstB {
		t.Error("scoped dependency is not cached within one scope")
	}

	if firstA == secondA {
		t.Error("scoped dependency is shared between scopes")
	}

	if counter != 2 {
		t.Errorf("scoped factory called %d times, want 2", counter)
	}
}

// TestScope_NestedAndOverride проверяет вложенные скоупы и локальное
// переопределение регистрации.
func TestScope_NestedAndOverride(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.Register("root")

	child := container.BeginScope()
	child.Register("child")

	grandchild := child.BeginScope()

	if got := container.MustResolve[string](); got != "root" {
		t.Errorf("container resolve = %q, want root", got)
	}

	if got := child.MustResolve[string](); got != "child" {
		t.Errorf("child resolve = %q, want child", got)
	}

	if got := grandchild.MustResolve[string](); got != "child" {
		t.Errorf("grandchild resolve = %q, want inherited child override", got)
	}
}

// TestScope_DisposeClosesScoped проверяет, что Scope.Dispose закрывает scoped в LIFO.
func TestScope_DisposeClosesScoped(t *testing.T) {
	t.Parallel()

	var events []string

	first := newCloserSpy("first", &events)
	second := newCloserSpy("second", &events)

	container := di.NewContainer()
	container.RegisterScopedNamed("first", func() (*closerSpy, error) {
		return first, nil
	})
	container.RegisterScopedNamed("second", func() (*closerSpy, error) {
		return second, nil
	})

	scope := container.BeginScope()
	scope.MustResolveNamed[*closerSpy]("first")
	scope.MustResolveNamed[*closerSpy]("second")

	if err := scope.Dispose(); err != nil {
		t.Fatalf("Dispose() error = %v", err)
	}

	if !first.isClosed() || !second.isClosed() {
		t.Fatal("scoped dependencies were not closed")
	}

	if len(events) != 2 || events[0] != "second" || events[1] != "first" {
		t.Errorf("close order = %v, want [second first]", events)
	}
}

// TestScope_ContainerDisposeIgnoresScoped проверяет, что контейнер не закрывает
// scoped-зависимости скоупа.
func TestScope_ContainerDisposeIgnoresScoped(t *testing.T) {
	t.Parallel()

	spy := newCloserSpy("scoped", nil)

	container := di.NewContainer()
	container.RegisterScoped(func() (*closerSpy, error) {
		return spy, nil
	})

	scope := container.BeginScope()
	scope.MustResolve[*closerSpy]()

	if err := container.Dispose(); err != nil {
		t.Fatalf("Dispose() error = %v", err)
	}

	if spy.isClosed() {
		t.Error("container closed scoped dependency owned by a scope")
	}

	if err := scope.Dispose(); err != nil {
		t.Fatalf("scope Dispose() error = %v", err)
	}

	if !spy.isClosed() {
		t.Error("scope did not close its scoped dependency")
	}
}

// TestScope_ContextCancellation проверяет проброс отмены контекста в фабрику.
func TestScope_ContextCancellation(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.RegisterScopedContext(func(ctx context.Context) (*closerSpy, error) {
		if err := ctx.Err(); err != nil {
			return nil, err
		}

		return newCloserSpy("scoped", nil), nil
	})

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	scope := container.BeginScope()

	_, err := scope.ResolveContext[*closerSpy](ctx)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("ResolveContext() error = %v, want context.Canceled", err)
	}
}

// TestScope_ContextValue проверяет проброс значений контекста в транзиентную фабрику.
func TestScope_ContextValue(t *testing.T) {
	t.Parallel()

	container := di.NewContainer()
	container.RegisterTransientContext(func(ctx context.Context) (string, error) {
		value, _ := ctx.Value(ctxKey{}).(string)

		return value, nil
	})

	ctx := context.WithValue(context.Background(), ctxKey{}, "request")

	got, err := container.ResolveContext[string](ctx)
	if err != nil {
		t.Fatalf("ResolveContext() error = %v", err)
	}

	if got != "request" {
		t.Errorf("ResolveContext() = %q, want request", got)
	}
}

// TestSingleton_UsesFirstContext проверяет, что для синглтона берется контекст
// первого вызова.
func TestSingleton_UsesFirstContext(t *testing.T) {
	t.Parallel()

	calls := 0

	container := di.NewContainer()
	container.RegisterFactoryContext(func(ctx context.Context) (string, error) {
		calls++

		value, _ := ctx.Value(ctxKey{}).(string)

		return value, nil
	})

	firstCtx := context.WithValue(context.Background(), ctxKey{}, "first")
	secondCtx := context.WithValue(context.Background(), ctxKey{}, "second")

	got, err := container.ResolveContext[string](firstCtx)
	if err != nil {
		t.Fatalf("ResolveContext() error = %v", err)
	}

	if got != "first" {
		t.Errorf("first resolve = %q, want first", got)
	}

	got, err = container.ResolveContext[string](secondCtx)
	if err != nil {
		t.Fatalf("ResolveContext() error = %v", err)
	}

	if got != "first" {
		t.Errorf("cached singleton = %q, want first", got)
	}

	if calls != 1 {
		t.Errorf("factory called %d times, want 1", calls)
	}
}
