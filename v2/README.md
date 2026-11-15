# di (v2)

Библиотека внедрения зависимостей для Go. Контейнер разрешает зависимости по
типу, поддерживает жизненные циклы singleton/transient/scoped, скоупы и
контекстные фабрики. `Provide` собирает граф зависимостей, `Verify` проверяет
его, `Dispose` закрывает ресурсы.

## Требования

Go 1.27 и новее.

## Установка

```bash
go get github.com/555f/di/v2
```

## Быстрый старт

```go
package main

import (
	"fmt"

	"github.com/555f/di/v2"
)

type Database struct {
	ConnectionString string
}

func main() {
	container := di.NewContainer()
	container.Register(Database{ConnectionString: "postgres://user:password@localhost/dbname"})

	db, err := container.Resolve[Database]()
	if err != nil {
		panic(err)
	}

	fmt.Println(db.ConnectionString)
}
```

## Регистрация зависимостей

Имена методов следуют одной схеме: `Register*` + опциональный суффикс `Named` +
опциональный суффикс `Context`. Ключом служит тип (и имя): `Resolve[string]` и
`Register[string]` работают с одной и той же зависимостью.

| Метод | Что регистрирует | Жизненный цикл |
|---|---|---|
| `Register(instance)` | готовый экземпляр | singleton |
| `RegisterNamed(name, instance)` | готовый экземпляр под именем | singleton |
| `RegisterMany(instances...)` | несколько экземпляров, не перезаписывая | singleton |
| `RegisterFactory(factory)` | `func() (T, error)` | singleton |
| `RegisterFactoryNamed(name, factory)` | фабрику под именем | singleton |
| `RegisterTransient(factory)` | `func() (T, error)` | transient |
| `RegisterTransientNamed(name, factory)` | фабрику под именем | transient |
| `RegisterScoped(factory)` | `func() (T, error)` | scoped |
| `RegisterScopedNamed(name, factory)` | фабрику под именем | scoped |
| `RegisterFactoryContext(factory)` | `func(ctx) (T, error)` | singleton |
| `RegisterFactoryContextNamed(name, factory)` | фабрику под именем | singleton |
| `RegisterTransientContext(factory)` | `func(ctx) (T, error)` | transient |
| `RegisterTransientContextNamed(name, factory)` | фабрику под именем | transient |
| `RegisterScopedContext(factory)` | `func(ctx) (T, error)` | scoped |
| `RegisterScopedContextNamed(name, factory)` | фабрику под именем | scoped |
| `RegisterAs[I](instance)` | экземпляр под интерфейсом `I` | singleton |
| `RegisterFactoryAs[I](factory)` | фабрику под интерфейсом `I` | singleton |
| `RegisterTransientAs[I](factory)` | фабрику под интерфейсом `I` | transient |

Повторная регистрация того же ключа (тип + имя) перезаписывает предыдущую
запись. `RegisterMany` добавляет экземпляры, не удаляя старые.

## Разрешение зависимостей

```go
value, err := container.Resolve[Database]()
if err != nil {
	// зависимость не зарегистрирована или фабрика вернула ошибку
}

named, err := container.ResolveNamed[string]("helloWorld")

values, err := container.ResolveAll[int]() // все регистрации типа int по порядку

// Паника, если разрешить не удалось
value := container.MustResolve[Database]()
named := container.MustResolveNamed[string]("helloWorld")
```

Зависимости, зарегистрированные под именем, доступны только через `*Named`
методы: безымянное разрешение и инжекторы их не видят. `ResolveAll` возвращает
все регистрации указанного типа, включая именованные, в порядке регистрации;
при отсутствии возвращается пустой срез и `nil`.

Управление регистрациями: `Has[T]`, `HasNamed[T](name)`, `Remove[T]`,
`RemoveNamed[T](name)`.

## Жизненные циклы

- **singleton** (`Register`, `RegisterFactory`) — создается при первом
  разрешении и кэшируется: фабрика вызывается не более одного раза.
- **transient** (`RegisterTransient`) — фабрика вызывается при каждом
  разрешении.
- **scoped** (`RegisterScoped`) — создается один раз в каждом скоупе и
  кэшируется в нём.

```go
container := di.NewContainer()

counter := 0
container.RegisterTransient(func() (int, error) {
	counter++

	return counter, nil
})

first, _ := container.Resolve[int]()  // 1
second, _ := container.Resolve[int]() // 2
```

## Скоупы

`BeginScope` создает вложенную область разрешения. Scoped-зависимости
кэшируются в скоупе, синглтоны и транзиенты разрешаются через корневой
контейнер. Поиск регистрации идет от скоупа к родителю и далее к корню, поэтому
`Scope.Register` переопределяет регистрацию родителя для этого и вложенных
скоупов.

```go
container := di.NewContainer()
container.RegisterScoped(func() (*Session, error) {
	return newSession(), nil
})

scope := container.BeginScope()
defer scope.Dispose() // закрывает scoped-экземпляры скоупа (LIFO)

session := scope.MustResolve[*Session]()  // создан в скоупе
session2 := scope.MustResolve[*Session]() // тот же экземпляр
```

`Scope` повторяет методы `Container`: `Resolve`, `ResolveNamed`,
`MustResolve`, `ResolveContext`, `ResolveAll`, `Injector`..`Injector10`, а также
`Register`/`RegisterNamed` для локальных переопределений.

## Контекст

Контекстные фабрики получают `ctx` из `ResolveContext`, включая отмену, дедлайн
и request-значения.

```go
container.RegisterScopedContext(func(ctx context.Context) (*Request, error) {
	return newRequest(ctx), nil
})

request, err := scope.ResolveContext[*Request](ctx)
```

Для синглтона используется контекст первого вызова. `Resolve` и `ResolveNamed`
вызывают контекстные фабрики с `context.Background()`.

## Автоматическая сборка: Provide

`Provide` (и `Provide2`..`Provide10`) собирает синглтон из перечисленных
типовых зависимостей: контейнер разрешает их сам.

```go
container.Register(21)
container.Register(true)

container.Provide2[string, int, bool](func(number int, exclaim bool) (string, error) {
	if exclaim {
		return fmt.Sprintf("%d!", number), nil
	}

	return fmt.Sprint(number), nil
})
```

### Проверка графа

`Verify` проверяет граф, построенный по `Provide`: находит незарегистрированные
зависимости и циклы, объединяя ошибки через `errors.Join`.

```go
if err := container.Verify(); err != nil {
	// errors.Is(err, di.ErrNotFound) или di.ErrCircularDependency
}
```

Циклы обнаруживаются и при разрешении: повторный вход в тот же ключ внутри одной
цепочки возвращает `di.ErrCircularDependency` без зависания на мьютексе. Это
верно для графа из `Provide`. Если фабрика из `RegisterFactory` сама вызывает
`Resolve`, граф по ней не строится.

## Привязка к интерфейсу

```go
type Logger interface{ Log(string) }

container.RegisterAs[Logger](&ConsoleLogger{})
container.RegisterFactoryAs[Logger](func() (*ConsoleLogger, error) {
	return &ConsoleLogger{}, nil
})

logger := container.MustResolve[Logger]()
```

Go не позволяет статически проверить «`T` реализует `I`» для двух
неограниченных параметров типа, поэтому несовпадение обнаруживается во время
выполнения паникой конверсии интерфейса.

## Инжекторы

Инжектор разрешает зависимости из контейнера и передает их в конструктор.

```go
newUserService := container.Injector(NewUserService)

service, err := newUserService()
if err != nil {
	// одна из зависимостей не разрешилась, конструктор не вызывался
}
```

Инжекторы есть для 1..10 зависимостей: `Injector`, `Injector2`, ..., `Injector10`
— и на `Container`, и на `Scope`. Зависимости разрешаются в момент вызова
инжектора, поэтому регистрировать их можно в любой момент до вызова.

## Очистка ресурсов: Dispose

`Container.Dispose` закрывает синглтоны, реализующие `io.Closer`, в обратном
порядке регистрации (LIFO). `Scope.Dispose` закрывает scoped-экземпляры скоупа.
Транзиенты принадлежат вызывающему коду и не закрываются.

```go
container := di.NewContainer()
container.RegisterFactory(func() (*sql.DB, error) {
	return sql.Open("postgres", dsn)
})

db := container.MustResolve[*sql.DB]()

defer container.Dispose() // db.Close() при остановке приложения
```

## Ошибки

Все ошибки разрешения оборачивают sentinel-ошибки через `%w` и проверяются через
`errors.Is`.

| Sentinel | Когда возникает |
|---|---|
| `ErrNotFound` | зависимость не зарегистрирована |
| `ErrInvalidRegistration` | по ключу лежит запись несовместимого типа |
| `ErrNoFactory` | у фабричной регистрации отсутствует фабрика |
| `ErrUnsupportedKind` | у регистрации неизвестный жизненный цикл |
| `ErrCircularDependency` | циклическая зависимость |

```go
if errors.Is(err, di.ErrNotFound) {
	// обработать отсутствующую зависимость
}
```

## Модули

Модуль группирует регистрации, чтобы собирать контейнер из независимых частей
приложения.

```go
type ServerModule struct{}

func (ServerModule) Register(container *di.Container) {
	container.RegisterFactory(func() (*http.Server, error) {
		return &http.Server{Addr: ":8080"}, nil
	})
}

func main() {
	container := di.NewContainer()

	di.RegisterModules(container,
		ServerModule{},
		di.ModuleFunc(func(c *di.Container) {
			c.Register("value")
		}),
	)
}
```

`di.NewModuleGroup` объединяет модули в `di.ModuleGroup`, который тоже
реализует интерфейс `di.Module`.

## Потокобезопасность

Все методы `Container` и `Scope` можно вызывать из нескольких горутин
одновременно. Разрешение синглтона защищено блокировкой: фабрика выполняется
один раз, остальные вызовы получают кэшированный экземпляр.

## Переход с v1

В v2 функции контейнера стали методами, появились скоупы и контекстные фабрики.

| v1 | v2 |
|---|---|
| `di.Register(c, value)` | `c.Register(value)` |
| `di.RegisterFactory(c, factory)` | `c.RegisterFactory(factory)` |
| `di.RegisterFactoryNamed(c, name, factory)` | `c.RegisterFactoryNamed(name, factory)` |
| `di.RegisterTransient(c, factory)` | `c.RegisterTransient(factory)` |
| `di.RegisterTransientNamed(c, name, factory)` | `c.RegisterTransientNamed(name, factory)` |
| `di.RegisterInstanceNamed(c, name, value)` | `c.RegisterNamed(name, value)` |
| `di.Resolve[T](c)` | `c.Resolve[T]()` |
| `di.ResolveNamed[T](c, name)` | `c.ResolveNamed[T](name)` |
| `di.MustResolve[T](c)` | `c.MustResolve[T]()` |
| `injector := di.NewInjector(ctor)`; `injector(c)` | `injector := c.Injector(ctor)`; `injector()` |
| `di.NewInjectorN(ctor)` для N = 2..10 | `c.InjectorN(ctor)` для N = 2..10 |
| `di.RegisterModules(c, modules...)` | без изменений |

Модули (`Module`, `ModuleFunc`, `ModuleGroup`, `NewModuleGroup`,
`RegisterModules`) перенесены в v2 без изменений API.

## Тестирование

```bash
cd v2
go test ./...
go test -race ./...
go vet ./...
```
