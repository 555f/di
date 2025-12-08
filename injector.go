package di

type Injector[T any] func(*Container) (T, error)

// NewInjector создает инжектор с автоматическим внедрением одной зависимости
func NewInjector[T any, D1 any](constructor func(D1) T) Injector[T] {
	return func(container *Container) (T, error) {
		var zero T

		dep1, err := Resolve[D1](container)
		if err != nil {
			return zero, err
		}

		result := constructor(dep1)

		return result, nil
	}
}

func NewInjector2[T any, D1, D2 any](constructor func(D1, D2) T) Injector[T] {
	return func(container *Container) (T, error) {
		var zero T

		dep1, err := Resolve[D1](container)
		if err != nil {
			return zero, err
		}

		dep2, err := Resolve[D2](container)
		if err != nil {
			return zero, err
		}

		result := constructor(dep1, dep2)

		return result, nil
	}
}

func NewInjector3[T any, D1, D2, D3 any](constructor func(D1, D2, D3) T) Injector[T] {
	return func(container *Container) (T, error) {
		var zero T

		dep1, err := Resolve[D1](container)
		if err != nil {
			return zero, err
		}

		dep2, err := Resolve[D2](container)
		if err != nil {
			return zero, err
		}

		dep3, err := Resolve[D3](container)
		if err != nil {
			return zero, err
		}

		result := constructor(dep1, dep2, dep3)

		return result, nil
	}
}

func NewInjector4[T any, D1, D2, D3, D4 any](constructor func(D1, D2, D3, D4) T) Injector[T] {
	return func(container *Container) (T, error) {
		var zero T

		dep1, err := Resolve[D1](container)
		if err != nil {
			return zero, err
		}

		dep2, err := Resolve[D2](container)
		if err != nil {
			return zero, err
		}

		dep3, err := Resolve[D3](container)
		if err != nil {
			return zero, err
		}

		dep4, err := Resolve[D4](container)
		if err != nil {
			return zero, err
		}

		result := constructor(dep1, dep2, dep3, dep4)

		return result, nil
	}
}

func NewInjector5[T any, D1, D2, D3, D4, D5 any](constructor func(D1, D2, D3, D4, D5) T) Injector[T] {
	return func(container *Container) (T, error) {
		var zero T

		dep1, err := Resolve[D1](container)
		if err != nil {
			return zero, err
		}

		dep2, err := Resolve[D2](container)
		if err != nil {
			return zero, err
		}

		dep3, err := Resolve[D3](container)
		if err != nil {
			return zero, err
		}

		dep4, err := Resolve[D4](container)
		if err != nil {
			return zero, err
		}

		dep5, err := Resolve[D5](container)
		if err != nil {
			return zero, err
		}

		result := constructor(dep1, dep2, dep3, dep4, dep5)

		return result, nil
	}
}

func NewInjector6[T any, D1, D2, D3, D4, D5, D6 any](constructor func(D1, D2, D3, D4, D5, D6) T) Injector[T] {
	return func(container *Container) (T, error) {
		var zero T

		dep1, err := Resolve[D1](container)
		if err != nil {
			return zero, err
		}

		dep2, err := Resolve[D2](container)
		if err != nil {
			return zero, err
		}

		dep3, err := Resolve[D3](container)
		if err != nil {
			return zero, err
		}

		dep4, err := Resolve[D4](container)
		if err != nil {
			return zero, err
		}

		dep5, err := Resolve[D5](container)
		if err != nil {
			return zero, err
		}

		dep6, err := Resolve[D6](container)
		if err != nil {
			return zero, err
		}

		result := constructor(dep1, dep2, dep3, dep4, dep5, dep6)

		return result, nil
	}
}

func NewInjector7[T any, D1, D2, D3, D4, D5, D6, D7 any](constructor func(D1, D2, D3, D4, D5, D6, D7) T) Injector[T] {
	return func(container *Container) (T, error) {
		var zero T

		dep1, err := Resolve[D1](container)
		if err != nil {
			return zero, err
		}

		dep2, err := Resolve[D2](container)
		if err != nil {
			return zero, err
		}

		dep3, err := Resolve[D3](container)
		if err != nil {
			return zero, err
		}

		dep4, err := Resolve[D4](container)
		if err != nil {
			return zero, err
		}

		dep5, err := Resolve[D5](container)
		if err != nil {
			return zero, err
		}

		dep6, err := Resolve[D6](container)
		if err != nil {
			return zero, err
		}

		dep7, err := Resolve[D7](container)
		if err != nil {
			return zero, err
		}

		result := constructor(dep1, dep2, dep3, dep4, dep5, dep6, dep7)

		return result, nil
	}
}

func NewInjector8[T any, D1, D2, D3, D4, D5, D6, D7, D8 any](constructor func(D1, D2, D3, D4, D5, D6, D7, D8) T) Injector[T] {
	return func(container *Container) (T, error) {
		var zero T

		dep1, err := Resolve[D1](container)
		if err != nil {
			return zero, err
		}

		dep2, err := Resolve[D2](container)
		if err != nil {
			return zero, err
		}

		dep3, err := Resolve[D3](container)
		if err != nil {
			return zero, err
		}

		dep4, err := Resolve[D4](container)
		if err != nil {
			return zero, err
		}

		dep5, err := Resolve[D5](container)
		if err != nil {
			return zero, err
		}

		dep6, err := Resolve[D6](container)
		if err != nil {
			return zero, err
		}

		dep7, err := Resolve[D7](container)
		if err != nil {
			return zero, err
		}

		dep8, err := Resolve[D8](container)
		if err != nil {
			return zero, err
		}

		result := constructor(dep1, dep2, dep3, dep4, dep5, dep6, dep7, dep8)

		return result, nil
	}
}

func NewInjector9[T any, D1, D2, D3, D4, D5, D6, D7, D8, D9 any](constructor func(D1, D2, D3, D4, D5, D6, D7, D8, D9) T) Injector[T] {
	return func(container *Container) (T, error) {
		var zero T

		dep1, err := Resolve[D1](container)
		if err != nil {
			return zero, err
		}

		dep2, err := Resolve[D2](container)
		if err != nil {
			return zero, err
		}

		dep3, err := Resolve[D3](container)
		if err != nil {
			return zero, err
		}

		dep4, err := Resolve[D4](container)
		if err != nil {
			return zero, err
		}

		dep5, err := Resolve[D5](container)
		if err != nil {
			return zero, err
		}

		dep6, err := Resolve[D6](container)
		if err != nil {
			return zero, err
		}

		dep7, err := Resolve[D7](container)
		if err != nil {
			return zero, err
		}

		dep8, err := Resolve[D8](container)
		if err != nil {
			return zero, err
		}

		dep9, err := Resolve[D9](container)
		if err != nil {
			return zero, err
		}

		result := constructor(dep1, dep2, dep3, dep4, dep5, dep6, dep7, dep8, dep9)

		return result, nil
	}
}

func NewInjector10[T any, D1, D2, D3, D4, D5, D6, D7, D8, D9, D10 any](constructor func(D1, D2, D3, D4, D5, D6, D7, D8, D9, D10) T) Injector[T] {
	return func(container *Container) (T, error) {
		var zero T

		dep1, err := Resolve[D1](container)
		if err != nil {
			return zero, err
		}

		dep2, err := Resolve[D2](container)
		if err != nil {
			return zero, err
		}

		dep3, err := Resolve[D3](container)
		if err != nil {
			return zero, err
		}

		dep4, err := Resolve[D4](container)
		if err != nil {
			return zero, err
		}

		dep5, err := Resolve[D5](container)
		if err != nil {
			return zero, err
		}

		dep6, err := Resolve[D6](container)
		if err != nil {
			return zero, err
		}

		dep7, err := Resolve[D7](container)
		if err != nil {
			return zero, err
		}

		dep8, err := Resolve[D8](container)
		if err != nil {
			return zero, err
		}

		dep9, err := Resolve[D9](container)
		if err != nil {
			return zero, err
		}

		dep10, err := Resolve[D10](container)
		if err != nil {
			return zero, err
		}

		result := constructor(dep1, dep2, dep3, dep4, dep5, dep6, dep7, dep8, dep9, dep10)

		return result, nil
	}
}
