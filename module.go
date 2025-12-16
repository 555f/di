package di

type Module interface {
	Register(container *Container)
}

type ModuleFunc func(*Container)

func (f ModuleFunc) Register(container *Container) {
	f(container)
}

type ModuleGroup []Module

func (mg ModuleGroup) Register(container *Container) {
	for _, module := range mg {
		module.Register(container)
	}
}

func NewModuleGroup(modules ...Module) ModuleGroup {
	return modules
}

func RegisterModules(container *Container, modules ...Module) {
	for _, module := range modules {
		module.Register(container)
	}
}
