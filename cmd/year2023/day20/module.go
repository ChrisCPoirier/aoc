package day20

type module interface {
	Name() string
	signal(module, bool) bool
	pulse() []module
	addDestination(module)
	Destinations() []module
	addSource(module)
	willSend() bool
}
