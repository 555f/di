package di

import "fmt"

type TypeKey[T any] struct {
	name string
}

func newTypeKey[T any](name ...string) TypeKey[T] {
	keyName := ""

	if len(name) > 0 {
		keyName = name[0]
	}

	return TypeKey[T]{name: keyName}
}

func (tk TypeKey[T]) String() string {
	return fmt.Sprintf("%T", tk)
}
