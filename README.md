# Библиотека внедрения зависимостей

## Вступление
Эта библиотека предоставляет простой и гибкий способ управления зависимостями в ваших приложениях Go с помощью обобщений (generics) и фабричных функций.

## Как использовать

### Регистрация зависимости
```go
package main

import (
  "fmt"
  "github.com/555f/di"
)

type Database struct {
  ConnectionString string
}

func main() {
  c := di.NewContainer()
  di.Register(c, Database{ConnectionString: "postgres://user:password@localhost/dbname"})

  fmt.Println("Database registered!")
}
```

### Разрешение зависимости

```go
package main

import (
  "fmt"
  "github.com/555f/di"
)

func main() {
  c := di.NewContainer()
  di.Register(c, "Hello, World!")

  value, err := di.Resolve[string](с)
  if err == nil {
    fmt.Println(value) // Выведет: Hello, World!
  }
}
```

## Описание методов

### di.NewContainer

Создает новый контейнер для внедрения зависимостей.

### di.Register

Регистрирует зависимость как готовый к использованию экземпляр.

```go
func Register[T any](с *Container, value T)
```

#### Параметры

<code>c</code>: контейнер для регистрации зависимости.

<code>value</code>: Экземпляр для регистрации.

### di.RegisterFactory

Регистрирует фабричную функцию для создания зависимости. Фабричная функция будет вызвана только один раз, а результат будет кэширован.

```go
func RegisterFactory[T any](c *Container, factory func() (T, error))
```

#### Параметры

<code>c</code>: контейнер для регистрации зависимости.

<code>factory</code>: фабричная функция, которая создаёт зависимость.

### di.Resolve

Разрешает зависимость из контейнера. Если зависимость зарегистрирована как фабричная функция, она будет вызвана для создания экземпляра.

```go
func Resolve[T any]() (T, error)
```

#### Возвращаемое значение

<code>T</code>: экземпляр зависимости.
<code>error</code>: ошибки при разрешении зависимости.

## Поддержка

Если у вас есть какие-либо вопросы или проблемы, пожалуйста, откройте задачу на GitHub.
