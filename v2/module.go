package di

// Module описывает набор зависимостей, который регистрируется в контейнере.
type Module interface {
	// Register регистрирует зависимости модуля в контейнере.
	Register(container *Container)
}

// ModuleFunc адаптирует обычную функцию к интерфейсу [Module].
type ModuleFunc func(*Container)

// Register вызывает функцию модуля.
func (f ModuleFunc) Register(container *Container) {
	f(container)
}

// ModuleGroup объединяет модули, которые регистрируются по порядку.
type ModuleGroup []Module

// Register регистрирует все модули группы по порядку.
func (mg ModuleGroup) Register(container *Container) {
	for _, module := range mg {
		module.Register(container)
	}
}

// NewModuleGroup объединяет модули в [ModuleGroup].
func NewModuleGroup(modules ...Module) ModuleGroup {
	return modules
}

// RegisterModules регистрирует модули в контейнере по порядку.
func RegisterModules(container *Container, modules ...Module) {
	for _, module := range modules {
		module.Register(container)
	}
}
