package di_test

import (
	"testing"

	"github.com/555f/di/v2"
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

	c.Register(42)

	value, err := c.Resolve[int]()
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

	c.RegisterFactory(func() (string, error) {
		return "Hello, World!", nil
	})

	value, err := c.Resolve[string]()
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

	c.RegisterFactoryNamed("helloWorld", func() (string, error) {
		return "Hello, World!", nil
	})

	value, err := c.ResolveNamed[string]("helloWorld")
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

	c.RegisterFactory(func() (int, error) {
		counter++
		return counter, nil
	})

	value1, err := c.Resolve[int]()
	if err != nil {
		t.Fatalf("Dependency not found: %s", err)
	}

	if value1 != 1 {
		t.Errorf("Expected 1, got %v", value1)
	}

	value2, err := c.Resolve[int]()
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

	_, err := c.Resolve[string]()
	if err == nil {
		t.Fatalf("Expected dependency not to be found")
	}
}

func TestAutoWiring_Injectors(t *testing.T) {
	c := di.NewContainer()

	c.Register[Logger](&ConsoleLogger{prefix: "[AUTO]"})

	injector := c.Injector(NewUserService)
	c.RegisterFactory(func() (*UserService, error) {
		return injector()
	})

	userService, err := c.Resolve[*UserService]()
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
