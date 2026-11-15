# Библиотека внедрения зависимостей

Библиотека управляет зависимостями в приложениях Go: зависимости регистрируются
как готовые значения или фабричные функции и разрешаются по типу с помощью
обобщений.

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

Регистрирует готовый экземпляр.

```go
func Register[T any](с *Container, value T)
```

#### Параметры

<code>c</code>: контейнер для регистрации зависимости.

<code>value</code>: экземпляр для регистрации.

### di.RegisterFactory

Регистрирует фабрику зависимости. Фабрика вызывается один раз, результат кэшируется.

```go
func RegisterFactory[T any](c *Container, factory func() (T, error))
```

#### Параметры

<code>c</code>: контейнер для регистрации зависимости.

<code>factory</code>: фабричная функция, которая создаёт зависимость.

### di.Resolve

Возвращает зависимость из контейнера. Если она зарегистрирована фабрикой, фабрика создаёт экземпляр.

```go
func Resolve[T any]() (T, error)
```

#### Возвращаемое значение

<code>T</code>: экземпляр зависимости.
<code>error</code>: ошибки при разрешении зависимости.

## Поддержка

Вопросы и проблемы: откройте задачу на GitHub.
