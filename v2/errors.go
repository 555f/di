package di

import "errors"

// Ошибки контейнера.
var (
	// ErrNotFound возвращается, когда зависимость не зарегистрирована.
	ErrNotFound = errors.New("dependency not registered")

	// ErrInvalidRegistration возвращается, когда по ключу лежит запись
	// несовместимого типа.
	ErrInvalidRegistration = errors.New("invalid registration type")

	// ErrNoFactory возвращается, когда у фабричной регистрации нет фабрики.
	ErrNoFactory = errors.New("no factory function provided")

	// ErrUnsupportedKind возвращается для неизвестного жизненного цикла.
	ErrUnsupportedKind = errors.New("unsupported kind")

	// ErrCircularDependency возвращается при циклической зависимости.
	ErrCircularDependency = errors.New("circular dependency")
)
