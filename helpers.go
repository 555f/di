package di

import "fmt"

func typeName[T any](t T) string {
	name := fmt.Sprintf("%T", t)
	if name == "<nil>" {
		name = fmt.Sprintf("%T", new(T))
	}

	return name
}
