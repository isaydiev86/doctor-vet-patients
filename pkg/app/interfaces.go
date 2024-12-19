package app

type Logger interface {
	Debug(msg string, args ...any)
	Error(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Fatal(msg string, args ...any)
}

func NewLifecycleComponent(name string, cmp Lifecycle) LifecycleComponent {
	return LifecycleComponent{name, cmp}
}

type LifecycleComponent struct {
	string
	Lifecycle
}
