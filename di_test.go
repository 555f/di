package di_test

import (
	"testing"

	"github.com/555f/di"
)

// TestRegisterResolve проверяет регистрацию и разрешение простых зависимостей.
func TestRegisterResolve(t *testing.T) {
	c := di.NewContainer()

	di.Register(c, 42)

	value, err := di.Resolve[int](c)
	if err != nil {
		t.Fatalf("Dependency not found: %s", err)
	}

	if value != 42 {
		t.Errorf("Expected 42, got %v", value)
	}
}

// TestRegisterFactoryResolve проверяет регистрацию и разрешение через фабричные функции.
func TestRegisterFactoryResolve(t *testing.T) {
	c := di.NewContainer()

	di.RegisterFactory(c, func() (string, error) {
		return "Hello, World!", nil
	})

	// Разрешаем зависимость
	value, err := di.Resolve[string](c)
	if err != nil {
		t.Fatalf("Dependency not found: %s", err)
	}

	if value != "Hello, World!" {
		t.Errorf(`Expected "Hello, World!", got %v`, value)
	}
}

// TestRegisterFactoryResolveNamed проверяет регистрацию и разрешение через именованные фабричные функции.
func TestRegisterFactoryResolveNamed(t *testing.T) {
	c := di.NewContainer()

	di.RegisterFactoryNamed(c, "helloWorld", func() (string, error) {
		return "Hello, World!", nil
	})

	// Разрешаем зависимость
	value, err := di.ResolveNamed[string](c, "helloWorld")
	if err != nil {
		t.Fatalf("Dependency not found: %s", err)
	}

	if value != "Hello, World!" {
		t.Errorf(`Expected "Hello, World!", got %v`, value)
	}
}

// TestSingleton проверяет, что синглтон создается только один раз.
func TestSingleton(t *testing.T) {
	c := di.NewContainer()

	counter := 0

	// Регистрируем синглтон
	di.RegisterFactory(c, func() (int, error) {
		counter++
		return counter, nil
	}, di.WithSingleton())

	// Первый вызов Resolve должен вызвать фабричную функцию
	value1, err := di.Resolve[int](c)
	if err != nil {
		t.Fatalf("Dependency not found: %s", err)
	}

	if value1 != 1 {
		t.Errorf("Expected 1, got %v", value1)
	}

	// Второй вызов Resolve должен вернуть кэшированное значение
	value2, err := di.Resolve[int](c)
	if err != nil {
		t.Fatalf("Dependency not found: %s", err)
	}

	if value2 != 1 {
		t.Errorf("Expected 1, got %v", value2)
	}

	if counter != 1 {
		t.Errorf("Factory function called %v times, expected 1", counter)
	}
}

// TestResolveNonExistent проверяет разрешение незарегистрированной зависимости.
func TestResolveNonExistent(t *testing.T) {
	c := di.NewContainer()

	// Пытаемся разрешить незарегистрированную зависимость
	_, err := di.Resolve[string](c)
	if err == nil {
		t.Fatalf("Expected dependency not to be found")
	}
}
