package di

import "context"

// Provide регистрирует синглтон типа T из конструктора без зависимостей.
func (c *Container) Provide[T any](constructor func() (T, error)) {
	reg := newProvideRegistration[T](nil)

	reg.factory = func(context.Context, *resolveState) (T, error) {
		return constructor()
	}

	c.root.data.store(reg.key, reg, false)
}

// Provide2 регистрирует синглтон типа T, зависящий от D1 и D2.
func (c *Container) Provide2[T any, D1, D2 any](constructor func(D1, D2) (T, error)) {
	reg := newProvideRegistration[T]([]any{any(newTypeKey[D1]("")), any(newTypeKey[D2](""))})

	reg.factory = func(ctx context.Context, st *resolveState) (T, error) {
		var zero T

		dep1, err := resolveFrom(ctx, st, newTypeKey[D1](""))
		if err != nil {
			return zero, err
		}

		dep2, err := resolveFrom(ctx, st, newTypeKey[D2](""))
		if err != nil {
			return zero, err
		}

		return constructor(dep1, dep2)
	}

	c.root.data.store(reg.key, reg, false)
}

// Provide3 регистрирует синглтон типа T, зависящий от D1..D3.
func (c *Container) Provide3[T any, D1, D2, D3 any](constructor func(D1, D2, D3) (T, error)) {
	reg := newProvideRegistration[T]([]any{
		any(newTypeKey[D1]("")),
		any(newTypeKey[D2]("")),
		any(newTypeKey[D3]("")),
	})

	reg.factory = func(ctx context.Context, st *resolveState) (T, error) {
		var zero T

		dep1, err := resolveFrom(ctx, st, newTypeKey[D1](""))
		if err != nil {
			return zero, err
		}

		dep2, err := resolveFrom(ctx, st, newTypeKey[D2](""))
		if err != nil {
			return zero, err
		}

		dep3, err := resolveFrom(ctx, st, newTypeKey[D3](""))
		if err != nil {
			return zero, err
		}

		return constructor(dep1, dep2, dep3)
	}

	c.root.data.store(reg.key, reg, false)
}

// Provide4 регистрирует синглтон типа T, зависящий от D1..D4.
func (c *Container) Provide4[T any, D1, D2, D3, D4 any](constructor func(D1, D2, D3, D4) (T, error)) {
	reg := newProvideRegistration[T]([]any{
		any(newTypeKey[D1]("")),
		any(newTypeKey[D2]("")),
		any(newTypeKey[D3]("")),
		any(newTypeKey[D4]("")),
	})

	reg.factory = func(ctx context.Context, st *resolveState) (T, error) {
		var zero T

		dep1, err := resolveFrom(ctx, st, newTypeKey[D1](""))
		if err != nil {
			return zero, err
		}

		dep2, err := resolveFrom(ctx, st, newTypeKey[D2](""))
		if err != nil {
			return zero, err
		}

		dep3, err := resolveFrom(ctx, st, newTypeKey[D3](""))
		if err != nil {
			return zero, err
		}

		dep4, err := resolveFrom(ctx, st, newTypeKey[D4](""))
		if err != nil {
			return zero, err
		}

		return constructor(dep1, dep2, dep3, dep4)
	}

	c.root.data.store(reg.key, reg, false)
}

// Provide5 регистрирует синглтон типа T, зависящий от D1..D5.
func (c *Container) Provide5[T any, D1, D2, D3, D4, D5 any](constructor func(D1, D2, D3, D4, D5) (T, error)) {
	reg := newProvideRegistration[T]([]any{
		any(newTypeKey[D1]("")),
		any(newTypeKey[D2]("")),
		any(newTypeKey[D3]("")),
		any(newTypeKey[D4]("")),
		any(newTypeKey[D5]("")),
	})

	reg.factory = func(ctx context.Context, st *resolveState) (T, error) {
		var zero T

		dep1, err := resolveFrom(ctx, st, newTypeKey[D1](""))
		if err != nil {
			return zero, err
		}

		dep2, err := resolveFrom(ctx, st, newTypeKey[D2](""))
		if err != nil {
			return zero, err
		}

		dep3, err := resolveFrom(ctx, st, newTypeKey[D3](""))
		if err != nil {
			return zero, err
		}

		dep4, err := resolveFrom(ctx, st, newTypeKey[D4](""))
		if err != nil {
			return zero, err
		}

		dep5, err := resolveFrom(ctx, st, newTypeKey[D5](""))
		if err != nil {
			return zero, err
		}

		return constructor(dep1, dep2, dep3, dep4, dep5)
	}

	c.root.data.store(reg.key, reg, false)
}

// Provide6 регистрирует синглтон типа T, зависящий от D1..D6.
func (c *Container) Provide6[T any, D1, D2, D3, D4, D5, D6 any](constructor func(D1, D2, D3, D4, D5, D6) (T, error)) {
	reg := newProvideRegistration[T]([]any{
		any(newTypeKey[D1]("")),
		any(newTypeKey[D2]("")),
		any(newTypeKey[D3]("")),
		any(newTypeKey[D4]("")),
		any(newTypeKey[D5]("")),
		any(newTypeKey[D6]("")),
	})

	reg.factory = func(ctx context.Context, st *resolveState) (T, error) {
		var zero T

		dep1, err := resolveFrom(ctx, st, newTypeKey[D1](""))
		if err != nil {
			return zero, err
		}

		dep2, err := resolveFrom(ctx, st, newTypeKey[D2](""))
		if err != nil {
			return zero, err
		}

		dep3, err := resolveFrom(ctx, st, newTypeKey[D3](""))
		if err != nil {
			return zero, err
		}

		dep4, err := resolveFrom(ctx, st, newTypeKey[D4](""))
		if err != nil {
			return zero, err
		}

		dep5, err := resolveFrom(ctx, st, newTypeKey[D5](""))
		if err != nil {
			return zero, err
		}

		dep6, err := resolveFrom(ctx, st, newTypeKey[D6](""))
		if err != nil {
			return zero, err
		}

		return constructor(dep1, dep2, dep3, dep4, dep5, dep6)
	}

	c.root.data.store(reg.key, reg, false)
}

// Provide7 регистрирует синглтон типа T, зависящий от D1..D7.
func (c *Container) Provide7[T any, D1, D2, D3, D4, D5, D6, D7 any](constructor func(D1, D2, D3, D4, D5, D6, D7) (T, error)) {
	reg := newProvideRegistration[T]([]any{
		any(newTypeKey[D1]("")),
		any(newTypeKey[D2]("")),
		any(newTypeKey[D3]("")),
		any(newTypeKey[D4]("")),
		any(newTypeKey[D5]("")),
		any(newTypeKey[D6]("")),
		any(newTypeKey[D7]("")),
	})

	if constructor != nil {
		reg.factory = func(ctx context.Context, st *resolveState) (T, error) {
			var zero T

			dep1, err := resolveFrom(ctx, st, newTypeKey[D1](""))
			if err != nil {
				return zero, err
			}

			dep2, err := resolveFrom(ctx, st, newTypeKey[D2](""))
			if err != nil {
				return zero, err
			}

			dep3, err := resolveFrom(ctx, st, newTypeKey[D3](""))
			if err != nil {
				return zero, err
			}

			dep4, err := resolveFrom(ctx, st, newTypeKey[D4](""))
			if err != nil {
				return zero, err
			}

			dep5, err := resolveFrom(ctx, st, newTypeKey[D5](""))
			if err != nil {
				return zero, err
			}

			dep6, err := resolveFrom(ctx, st, newTypeKey[D6](""))
			if err != nil {
				return zero, err
			}

			dep7, err := resolveFrom(ctx, st, newTypeKey[D7](""))
			if err != nil {
				return zero, err
			}

			return constructor(dep1, dep2, dep3, dep4, dep5, dep6, dep7)
		}
	}

	c.root.data.store(reg.key, reg, false)
}

// Provide8 регистрирует синглтон типа T, зависящий от D1..D8.
func (c *Container) Provide8[T any, D1, D2, D3, D4, D5, D6, D7, D8 any](constructor func(D1, D2, D3, D4, D5, D6, D7, D8) (T, error)) {
	reg := newProvideRegistration[T]([]any{
		any(newTypeKey[D1]("")),
		any(newTypeKey[D2]("")),
		any(newTypeKey[D3]("")),
		any(newTypeKey[D4]("")),
		any(newTypeKey[D5]("")),
		any(newTypeKey[D6]("")),
		any(newTypeKey[D7]("")),
		any(newTypeKey[D8]("")),
	})

	if constructor != nil {
		reg.factory = func(ctx context.Context, st *resolveState) (T, error) {
			var zero T

			dep1, err := resolveFrom(ctx, st, newTypeKey[D1](""))
			if err != nil {
				return zero, err
			}

			dep2, err := resolveFrom(ctx, st, newTypeKey[D2](""))
			if err != nil {
				return zero, err
			}

			dep3, err := resolveFrom(ctx, st, newTypeKey[D3](""))
			if err != nil {
				return zero, err
			}

			dep4, err := resolveFrom(ctx, st, newTypeKey[D4](""))
			if err != nil {
				return zero, err
			}

			dep5, err := resolveFrom(ctx, st, newTypeKey[D5](""))
			if err != nil {
				return zero, err
			}

			dep6, err := resolveFrom(ctx, st, newTypeKey[D6](""))
			if err != nil {
				return zero, err
			}

			dep7, err := resolveFrom(ctx, st, newTypeKey[D7](""))
			if err != nil {
				return zero, err
			}

			dep8, err := resolveFrom(ctx, st, newTypeKey[D8](""))
			if err != nil {
				return zero, err
			}

			return constructor(dep1, dep2, dep3, dep4, dep5, dep6, dep7, dep8)
		}
	}

	c.root.data.store(reg.key, reg, false)
}

// Provide9 регистрирует синглтон типа T, зависящий от D1..D9.
func (c *Container) Provide9[T any, D1, D2, D3, D4, D5, D6, D7, D8, D9 any](constructor func(D1, D2, D3, D4, D5, D6, D7, D8, D9) (T, error)) {
	reg := newProvideRegistration[T]([]any{
		any(newTypeKey[D1]("")),
		any(newTypeKey[D2]("")),
		any(newTypeKey[D3]("")),
		any(newTypeKey[D4]("")),
		any(newTypeKey[D5]("")),
		any(newTypeKey[D6]("")),
		any(newTypeKey[D7]("")),
		any(newTypeKey[D8]("")),
		any(newTypeKey[D9]("")),
	})

	if constructor != nil {
		reg.factory = func(ctx context.Context, st *resolveState) (T, error) {
			var zero T

			dep1, err := resolveFrom(ctx, st, newTypeKey[D1](""))
			if err != nil {
				return zero, err
			}

			dep2, err := resolveFrom(ctx, st, newTypeKey[D2](""))
			if err != nil {
				return zero, err
			}

			dep3, err := resolveFrom(ctx, st, newTypeKey[D3](""))
			if err != nil {
				return zero, err
			}

			dep4, err := resolveFrom(ctx, st, newTypeKey[D4](""))
			if err != nil {
				return zero, err
			}

			dep5, err := resolveFrom(ctx, st, newTypeKey[D5](""))
			if err != nil {
				return zero, err
			}

			dep6, err := resolveFrom(ctx, st, newTypeKey[D6](""))
			if err != nil {
				return zero, err
			}

			dep7, err := resolveFrom(ctx, st, newTypeKey[D7](""))
			if err != nil {
				return zero, err
			}

			dep8, err := resolveFrom(ctx, st, newTypeKey[D8](""))
			if err != nil {
				return zero, err
			}

			dep9, err := resolveFrom(ctx, st, newTypeKey[D9](""))
			if err != nil {
				return zero, err
			}

			return constructor(dep1, dep2, dep3, dep4, dep5, dep6, dep7, dep8, dep9)
		}
	}

	c.root.data.store(reg.key, reg, false)
}

// Provide10 регистрирует синглтон типа T, зависящий от D1..D10.
func (c *Container) Provide10[T any, D1, D2, D3, D4, D5, D6, D7, D8, D9, D10 any](constructor func(D1, D2, D3, D4, D5, D6, D7, D8, D9, D10) (T, error)) {
	reg := newProvideRegistration[T]([]any{
		any(newTypeKey[D1]("")),
		any(newTypeKey[D2]("")),
		any(newTypeKey[D3]("")),
		any(newTypeKey[D4]("")),
		any(newTypeKey[D5]("")),
		any(newTypeKey[D6]("")),
		any(newTypeKey[D7]("")),
		any(newTypeKey[D8]("")),
		any(newTypeKey[D9]("")),
		any(newTypeKey[D10]("")),
	})

	if constructor != nil {
		reg.factory = func(ctx context.Context, st *resolveState) (T, error) {
			var zero T

			dep1, err := resolveFrom(ctx, st, newTypeKey[D1](""))
			if err != nil {
				return zero, err
			}

			dep2, err := resolveFrom(ctx, st, newTypeKey[D2](""))
			if err != nil {
				return zero, err
			}

			dep3, err := resolveFrom(ctx, st, newTypeKey[D3](""))
			if err != nil {
				return zero, err
			}

			dep4, err := resolveFrom(ctx, st, newTypeKey[D4](""))
			if err != nil {
				return zero, err
			}

			dep5, err := resolveFrom(ctx, st, newTypeKey[D5](""))
			if err != nil {
				return zero, err
			}

			dep6, err := resolveFrom(ctx, st, newTypeKey[D6](""))
			if err != nil {
				return zero, err
			}

			dep7, err := resolveFrom(ctx, st, newTypeKey[D7](""))
			if err != nil {
				return zero, err
			}

			dep8, err := resolveFrom(ctx, st, newTypeKey[D8](""))
			if err != nil {
				return zero, err
			}

			dep9, err := resolveFrom(ctx, st, newTypeKey[D9](""))
			if err != nil {
				return zero, err
			}

			dep10, err := resolveFrom(ctx, st, newTypeKey[D10](""))
			if err != nil {
				return zero, err
			}

			return constructor(dep1, dep2, dep3, dep4, dep5, dep6, dep7, dep8, dep9, dep10)
		}
	}

	c.root.data.store(reg.key, reg, false)
}
