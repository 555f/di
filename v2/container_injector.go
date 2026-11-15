package di

// Injector создает инжектор, привязанный к контейнеру.
func (c *Container) Injector[T any, D1 any](f func(D1) T) Injector[T] {
	return c.root.Injector(f)
}

// Injector2 создает инжектор, привязанный к контейнеру.
func (c *Container) Injector2[T any, D1, D2 any](f func(D1, D2) T) Injector[T] {
	return c.root.Injector2(f)
}

// Injector3 создает инжектор, привязанный к контейнеру.
func (c *Container) Injector3[T any, D1, D2, D3 any](f func(D1, D2, D3) T) Injector[T] {
	return c.root.Injector3(f)
}

// Injector4 создает инжектор, привязанный к контейнеру.
func (c *Container) Injector4[T any, D1, D2, D3, D4 any](f func(D1, D2, D3, D4) T) Injector[T] {
	return c.root.Injector4(f)
}

// Injector5 создает инжектор, привязанный к контейнеру.
func (c *Container) Injector5[T any, D1, D2, D3, D4, D5 any](f func(D1, D2, D3, D4, D5) T) Injector[T] {
	return c.root.Injector5(f)
}

// Injector6 создает инжектор, привязанный к контейнеру.
func (c *Container) Injector6[T any, D1, D2, D3, D4, D5, D6 any](f func(D1, D2, D3, D4, D5, D6) T) Injector[T] {
	return c.root.Injector6(f)
}

// Injector7 создает инжектор, привязанный к контейнеру.
func (c *Container) Injector7[T any, D1, D2, D3, D4, D5, D6, D7 any](f func(D1, D2, D3, D4, D5, D6, D7) T) Injector[T] {
	return c.root.Injector7(f)
}

// Injector8 создает инжектор, привязанный к контейнеру.
func (c *Container) Injector8[T any, D1, D2, D3, D4, D5, D6, D7, D8 any](f func(D1, D2, D3, D4, D5, D6, D7, D8) T) Injector[T] {
	return c.root.Injector8(f)
}

// Injector9 создает инжектор, привязанный к контейнеру.
func (c *Container) Injector9[T any, D1, D2, D3, D4, D5, D6, D7, D8, D9 any](f func(D1, D2, D3, D4, D5, D6, D7, D8, D9) T) Injector[T] {
	return c.root.Injector9(f)
}

// Injector10 создает инжектор, привязанный к контейнеру.
func (c *Container) Injector10[T any, D1, D2, D3, D4, D5, D6, D7, D8, D9, D10 any](f func(D1, D2, D3, D4, D5, D6, D7, D8, D9, D10) T) Injector[T] {
	return c.root.Injector10(f)
}
