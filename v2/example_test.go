package di_test

import (
	"errors"
	"fmt"

	"github.com/555f/di/v2"
)

// Example демонстрирует регистрацию и разрешение зависимости.
func Example() {
	container := di.NewContainer()

	container.Register("Hello, World!")

	value, err := container.Resolve[string]()
	if err != nil {
		fmt.Println("error:", err)

		return
	}

	fmt.Println(value)
	// Output: Hello, World!
}

// ExampleContainer_RegisterFactory демонстрирует регистрацию зависимости через
// фабрику.
func ExampleContainer_RegisterFactory() {
	container := di.NewContainer()

	container.RegisterFactory(func() (int, error) {
		return 42, nil
	})

	value, _ := container.Resolve[int]()

	fmt.Println(value)
	// Output: 42
}

// ExampleContainer_RegisterTransient демонстрирует создание зависимости при
// каждом разрешении.
func ExampleContainer_RegisterTransient() {
	container := di.NewContainer()

	counter := 0

	container.RegisterTransient(func() (int, error) {
		counter++

		return counter, nil
	})

	first, _ := container.Resolve[int]()
	second, _ := container.Resolve[int]()

	fmt.Println(first, second)
	// Output: 1 2
}

// ExampleContainer_ResolveNamed демонстрирует разрешение именованной
// зависимости.
func ExampleContainer_ResolveNamed() {
	container := di.NewContainer()

	container.RegisterFactoryNamed("helloWorld", func() (string, error) {
		return "Hello, World!", nil
	})

	value, _ := container.ResolveNamed[string]("helloWorld")

	fmt.Println(value)
	// Output: Hello, World!
}

// ExampleContainer_Resolve_unregistered демонстрирует ошибку разрешения
// незарегистрированной зависимости.
func ExampleContainer_Resolve_unregistered() {
	container := di.NewContainer()

	if _, err := container.Resolve[string](); err != nil {
		fmt.Println("resolve failed")
	}
	// Output: resolve failed
}

// ExampleContainer_MustResolve демонстрирует разрешение зависимости без
// обработки ошибки.
func ExampleContainer_MustResolve() {
	container := di.NewContainer()

	container.Register(42)

	fmt.Println(container.MustResolve[int]())
	// Output: 42
}

// ExampleContainer_Injector демонстрирует автоматическое внедрение зависимости
// в конструктор.
func ExampleContainer_Injector() {
	container := di.NewContainer()

	container.Register("Hello")

	greet := container.Injector(func(name string) string {
		return name + ", World!"
	})

	message, _ := greet()

	fmt.Println(message)
	// Output: Hello, World!
}

// ExampleRegisterModules демонстрирует регистрацию зависимостей через модули.
func ExampleRegisterModules() {
	container := di.NewContainer()

	di.RegisterModules(container,
		di.ModuleFunc(func(c *di.Container) {
			c.Register(1)
		}),
		di.ModuleFunc(func(c *di.Container) {
			c.Register("two")
		}),
	)

	number, _ := container.Resolve[int]()
	word, _ := container.Resolve[string]()

	fmt.Println(number, word)
	// Output: 1 two
}

// ExampleContainer_BeginScope демонстрирует scoped-зависимости: экземпляр
// создается один раз в скоупе и различается между скоупами.
func ExampleContainer_BeginScope() {
	container := di.NewContainer()

	counter := 0

	container.RegisterScoped(func() (int, error) {
		counter++

		return counter, nil
	})

	first := container.BeginScope()
	second := container.BeginScope()

	firstA, _ := first.Resolve[int]()
	firstB, _ := first.Resolve[int]()
	secondA, _ := second.Resolve[int]()

	fmt.Println(firstA, firstB, secondA)
	// Output: 1 1 2
}

// ExampleContainer_Provide демонстрирует сборку синглтона из типовых
// зависимостей.
func ExampleContainer_Provide() {
	container := di.NewContainer()
	container.Register(21)
	container.Register(true)

	container.Provide2(func(number int, exclaim bool) (string, error) {
		if exclaim {
			return fmt.Sprintf("%d!", number), nil
		}

		return fmt.Sprint(number), nil
	})

	value, _ := container.Resolve[string]()

	fmt.Println(value)
	// Output: 21!
}

// ExampleContainer_Verify демонстрирует проверку графа зависимостей.
func ExampleContainer_Verify() {
	container := di.NewContainer()

	container.Provide2(func(number int, exclaim bool) (string, error) {
		return fmt.Sprint(number), nil
	})

	fmt.Println(errors.Is(container.Verify(), di.ErrNotFound))
	// Output: true
}

// ExampleContainer_ResolveAll демонстрирует разрешение всех зависимостей типа в
// порядке регистрации.
func ExampleContainer_ResolveAll() {
	container := di.NewContainer()
	container.RegisterMany(1, 2, 3)

	values, _ := container.ResolveAll[int]()

	fmt.Println(values)
	// Output: [1 2 3]
}

// ExampleContainer_RegisterAs демонстрирует привязку конкретного типа к интерфейсу.
func ExampleContainer_RegisterAs() {
	container := di.NewContainer()
	container.RegisterAs[greeter](simpleGreeter{msg: "hello"})

	fmt.Println(container.MustResolve[greeter]().Greet())
	// Output: hello
}

// ExampleContainer_Dispose демонстрирует закрытие синглтонов, реализующих
// io.Closer.
func ExampleContainer_Dispose() {
	container := di.NewContainer()
	container.RegisterFactory(func() (*closerSpy, error) {
		return newCloserSpy("db", nil), nil
	})

	spy := container.MustResolve[*closerSpy]()

	_ = container.Dispose()

	fmt.Println(spy.isClosed())
	// Output: true
}
