package di_test

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/555f/di/v2"
)

// TestSingleton_ConcurrentResolve проверяет, что фабрика синглтона вызывается
// ровно один раз при конкурентном разрешении.
func TestSingleton_ConcurrentResolve(t *testing.T) {
	t.Parallel()

	var calls atomic.Int32

	container := di.NewContainer()
	container.RegisterFactory(func() (int, error) {
		calls.Add(1)
		time.Sleep(time.Millisecond)

		return 42, nil
	})

	const workers = 32

	results := make([]int, workers)

	var wg sync.WaitGroup

	for i := range results {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()

			results[index] = container.MustResolve[int]()
		}(i)
	}

	wg.Wait()

	if got := calls.Load(); got != 1 {
		t.Errorf("factory called %d times, want 1", got)
	}

	for index, value := range results {
		if value != 42 {
			t.Errorf("results[%d] = %d, want 42", index, value)
		}
	}
}

// TestScope_ConcurrentScoped проверяет, что scoped-зависимость создается один
// раз внутри скоупа при конкурентном разрешении.
func TestScope_ConcurrentScoped(t *testing.T) {
	t.Parallel()

	var calls atomic.Int32

	container := di.NewContainer()
	container.RegisterScoped(func() (*closerSpy, error) {
		calls.Add(1)
		time.Sleep(time.Millisecond)

		return newCloserSpy("scoped", nil), nil
	})

	scope := container.BeginScope()

	const workers = 16

	results := make([]*closerSpy, workers)

	var wg sync.WaitGroup

	for i := range results {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()

			results[index] = scope.MustResolve[*closerSpy]()
		}(i)
	}

	wg.Wait()

	if got := calls.Load(); got != 1 {
		t.Errorf("scoped factory called %d times, want 1", got)
	}

	for index, value := range results {
		if value != results[0] {
			t.Errorf("results[%d] is a different instance than results[0]", index)
		}
	}
}
