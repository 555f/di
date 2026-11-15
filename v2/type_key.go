package di

// typeKey идентифицирует зависимость по типу T и опциональному имени.
type typeKey[T any] struct {
	name string
}

func newTypeKey[T any](name string) typeKey[T] {
	return typeKey[T]{name: name}
}

// String возвращает имя регистрации.
func (k typeKey[T]) String() string {
	return k.name
}
