package di_test

import (
	"testing"

	"github.com/555f/di"
)

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct {
	prefix string
}

func (c *ConsoleLogger) Log(message string) {

}

type UserService struct {
	logger Logger
}

func NewUserService(logger Logger) *UserService {
	return &UserService{logger: logger}
}

// Test_RegisterResolve проверяет регистрацию и разрешение простых зависимостей.
func Test_RegisterResolve(t *testing.T) {
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

// Test_RegisterFactoryResolve проверяет регистрацию и разрешение через фабричные функции.
func Test_RegisterFactoryResolve(t *testing.T) {
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

// Test_RegisterFactoryResolveNamed проверяет регистрацию и разрешение через именованные фабричные функции.
func Test_RegisterFactoryResolveNamed(t *testing.T) {
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

// Test_Singleton проверяет, что создается только один раз.
func Test_Singleton(t *testing.T) {
	c := di.NewContainer()

	counter := 0

	// Регистрируем синглтон
	di.RegisterFactory(c, func() (int, error) {
		counter++
		return counter, nil
	})

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

// Test_ResolveNonExistent проверяет разрешение незарегистрированной зависимости.
func Test_ResolveNonExistent(t *testing.T) {
	c := di.NewContainer()

	// Пытаемся разрешить незарегистрированную зависимость
	_, err := di.Resolve[string](c)
	if err == nil {
		t.Fatalf("Expected dependency not to be found")
	}
}

func TestAutoWiring_Injectors(t *testing.T) {
	container := di.NewContainer()

	di.Register[Logger](container, &ConsoleLogger{prefix: "[AUTO]"})

	injector := di.NewInjector(NewUserService)
	di.RegisterFactory(container, func() (*UserService, error) {
		return injector(container)
	})

	userService, err := di.Resolve[*UserService](container)
	if err != nil {
		t.Fatalf("Failed to resolve with auto-wiring: %v", err)
	}

	if userService == nil {
		t.Fatal("Auto-wired UserService should not be nil")
	}

	if userService.logger == nil {
		t.Fatal("UserService logger should not be nil")
	}
}
