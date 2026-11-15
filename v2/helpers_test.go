package di_test

import "sync"

// closerSpy записывает факт и порядок закрытия, чтобы проверять Dispose.
type closerSpy struct {
	name string
	err  error

	mu     sync.Mutex
	events *[]string
	closes int
	closed bool
}

func newCloserSpy(name string, events *[]string) *closerSpy {
	return &closerSpy{name: name, events: events}
}

func (s *closerSpy) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.closes++
	s.closed = true

	if s.events != nil {
		*s.events = append(*s.events, s.name)
	}

	return s.err
}

func (s *closerSpy) closeCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.closes
}

func (s *closerSpy) isClosed() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.closed
}

// greeter и реализация нужны для проверки RegisterAs.
type greeter interface {
	Greet() string
}

type simpleGreeter struct {
	msg string
}

func (g simpleGreeter) Greet() string {
	return g.msg
}

// svcA и svcB образуют цикл для проверки детекта циклических зависимостей.
type svcA struct {
	b *svcB
}

type svcB struct {
	a *svcA
}

// ctxKey — ключ для проверки проброса значений контекста в фабрику.
type ctxKey struct{}
