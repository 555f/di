package di

// Injector - функция, которая создает значение T из зависимостей скоупа.
//
// Зависимости разрешаются в момент вызова инжектора, а не при его создании,
// поэтому порядок регистрации не важен.
type Injector[T any] func() (T, error)

// Injector создает инжектор, который разрешает зависимость D1 и передает ее
// в constructor.
func (s *Scope) Injector[T any, D1 any](f func(D1) T) Injector[T] {
	return func() (T, error) {
		var zero T

		dep1, err := s.Resolve[D1]()
		if err != nil {
			return zero, err
		}

		return f(dep1), nil
	}
}

// Injector2 создает инжектор, который разрешает зависимости D1 и D2.
func (s *Scope) Injector2[T any, D1, D2 any](f func(D1, D2) T) Injector[T] {
	return func() (T, error) {
		var zero T

		dep1, err := s.Resolve[D1]()
		if err != nil {
			return zero, err
		}

		dep2, err := s.Resolve[D2]()
		if err != nil {
			return zero, err
		}

		return f(dep1, dep2), nil
	}
}

// Injector3 создает инжектор, который разрешает зависимости D1..D3.
func (s *Scope) Injector3[T any, D1, D2, D3 any](f func(D1, D2, D3) T) Injector[T] {
	return func() (T, error) {
		var zero T

		dep1, err := s.Resolve[D1]()
		if err != nil {
			return zero, err
		}

		dep2, err := s.Resolve[D2]()
		if err != nil {
			return zero, err
		}

		dep3, err := s.Resolve[D3]()
		if err != nil {
			return zero, err
		}

		return f(dep1, dep2, dep3), nil
	}
}

// Injector4 создает инжектор, который разрешает зависимости D1..D4.
func (s *Scope) Injector4[T any, D1, D2, D3, D4 any](f func(D1, D2, D3, D4) T) Injector[T] {
	return func() (T, error) {
		var zero T

		dep1, err := s.Resolve[D1]()
		if err != nil {
			return zero, err
		}

		dep2, err := s.Resolve[D2]()
		if err != nil {
			return zero, err
		}

		dep3, err := s.Resolve[D3]()
		if err != nil {
			return zero, err
		}

		dep4, err := s.Resolve[D4]()
		if err != nil {
			return zero, err
		}

		return f(dep1, dep2, dep3, dep4), nil
	}
}

// Injector5 создает инжектор, который разрешает зависимости D1..D5.
func (s *Scope) Injector5[T any, D1, D2, D3, D4, D5 any](f func(D1, D2, D3, D4, D5) T) Injector[T] {
	return func() (T, error) {
		var zero T

		dep1, err := s.Resolve[D1]()
		if err != nil {
			return zero, err
		}

		dep2, err := s.Resolve[D2]()
		if err != nil {
			return zero, err
		}

		dep3, err := s.Resolve[D3]()
		if err != nil {
			return zero, err
		}

		dep4, err := s.Resolve[D4]()
		if err != nil {
			return zero, err
		}

		dep5, err := s.Resolve[D5]()
		if err != nil {
			return zero, err
		}

		return f(dep1, dep2, dep3, dep4, dep5), nil
	}
}

// Injector6 создает инжектор, который разрешает зависимости D1..D6.
func (s *Scope) Injector6[T any, D1, D2, D3, D4, D5, D6 any](f func(D1, D2, D3, D4, D5, D6) T) Injector[T] {
	return func() (T, error) {
		var zero T

		dep1, err := s.Resolve[D1]()
		if err != nil {
			return zero, err
		}

		dep2, err := s.Resolve[D2]()
		if err != nil {
			return zero, err
		}

		dep3, err := s.Resolve[D3]()
		if err != nil {
			return zero, err
		}

		dep4, err := s.Resolve[D4]()
		if err != nil {
			return zero, err
		}

		dep5, err := s.Resolve[D5]()
		if err != nil {
			return zero, err
		}

		dep6, err := s.Resolve[D6]()
		if err != nil {
			return zero, err
		}

		return f(dep1, dep2, dep3, dep4, dep5, dep6), nil
	}
}

// Injector7 создает инжектор, который разрешает зависимости D1..D7.
func (s *Scope) Injector7[T any, D1, D2, D3, D4, D5, D6, D7 any](f func(D1, D2, D3, D4, D5, D6, D7) T) Injector[T] {
	return func() (T, error) {
		var zero T

		dep1, err := s.Resolve[D1]()
		if err != nil {
			return zero, err
		}

		dep2, err := s.Resolve[D2]()
		if err != nil {
			return zero, err
		}

		dep3, err := s.Resolve[D3]()
		if err != nil {
			return zero, err
		}

		dep4, err := s.Resolve[D4]()
		if err != nil {
			return zero, err
		}

		dep5, err := s.Resolve[D5]()
		if err != nil {
			return zero, err
		}

		dep6, err := s.Resolve[D6]()
		if err != nil {
			return zero, err
		}

		dep7, err := s.Resolve[D7]()
		if err != nil {
			return zero, err
		}

		return f(dep1, dep2, dep3, dep4, dep5, dep6, dep7), nil
	}
}

// Injector8 создает инжектор, который разрешает зависимости D1..D8.
func (s *Scope) Injector8[T any, D1, D2, D3, D4, D5, D6, D7, D8 any](f func(D1, D2, D3, D4, D5, D6, D7, D8) T) Injector[T] {
	return func() (T, error) {
		var zero T

		dep1, err := s.Resolve[D1]()
		if err != nil {
			return zero, err
		}

		dep2, err := s.Resolve[D2]()
		if err != nil {
			return zero, err
		}

		dep3, err := s.Resolve[D3]()
		if err != nil {
			return zero, err
		}

		dep4, err := s.Resolve[D4]()
		if err != nil {
			return zero, err
		}

		dep5, err := s.Resolve[D5]()
		if err != nil {
			return zero, err
		}

		dep6, err := s.Resolve[D6]()
		if err != nil {
			return zero, err
		}

		dep7, err := s.Resolve[D7]()
		if err != nil {
			return zero, err
		}

		dep8, err := s.Resolve[D8]()
		if err != nil {
			return zero, err
		}

		return f(dep1, dep2, dep3, dep4, dep5, dep6, dep7, dep8), nil
	}
}

// Injector9 создает инжектор, который разрешает зависимости D1..D9.
func (s *Scope) Injector9[T any, D1, D2, D3, D4, D5, D6, D7, D8, D9 any](f func(D1, D2, D3, D4, D5, D6, D7, D8, D9) T) Injector[T] {
	return func() (T, error) {
		var zero T

		dep1, err := s.Resolve[D1]()
		if err != nil {
			return zero, err
		}

		dep2, err := s.Resolve[D2]()
		if err != nil {
			return zero, err
		}

		dep3, err := s.Resolve[D3]()
		if err != nil {
			return zero, err
		}

		dep4, err := s.Resolve[D4]()
		if err != nil {
			return zero, err
		}

		dep5, err := s.Resolve[D5]()
		if err != nil {
			return zero, err
		}

		dep6, err := s.Resolve[D6]()
		if err != nil {
			return zero, err
		}

		dep7, err := s.Resolve[D7]()
		if err != nil {
			return zero, err
		}

		dep8, err := s.Resolve[D8]()
		if err != nil {
			return zero, err
		}

		dep9, err := s.Resolve[D9]()
		if err != nil {
			return zero, err
		}

		return f(dep1, dep2, dep3, dep4, dep5, dep6, dep7, dep8, dep9), nil
	}
}

// Injector10 создает инжектор, который разрешает зависимости D1..D10.
func (s *Scope) Injector10[T any, D1, D2, D3, D4, D5, D6, D7, D8, D9, D10 any](f func(D1, D2, D3, D4, D5, D6, D7, D8, D9, D10) T) Injector[T] {
	return func() (T, error) {
		var zero T

		dep1, err := s.Resolve[D1]()
		if err != nil {
			return zero, err
		}

		dep2, err := s.Resolve[D2]()
		if err != nil {
			return zero, err
		}

		dep3, err := s.Resolve[D3]()
		if err != nil {
			return zero, err
		}

		dep4, err := s.Resolve[D4]()
		if err != nil {
			return zero, err
		}

		dep5, err := s.Resolve[D5]()
		if err != nil {
			return zero, err
		}

		dep6, err := s.Resolve[D6]()
		if err != nil {
			return zero, err
		}

		dep7, err := s.Resolve[D7]()
		if err != nil {
			return zero, err
		}

		dep8, err := s.Resolve[D8]()
		if err != nil {
			return zero, err
		}

		dep9, err := s.Resolve[D9]()
		if err != nil {
			return zero, err
		}

		dep10, err := s.Resolve[D10]()
		if err != nil {
			return zero, err
		}

		return f(dep1, dep2, dep3, dep4, dep5, dep6, dep7, dep8, dep9, dep10), nil
	}
}
