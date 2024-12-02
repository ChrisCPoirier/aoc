package day20

type base struct {
	name         string
	destinations []module
	notify       chan notify
}

func (b *base) Name() string {
	return b.name
}

func (b *base) pulse() []module {
	return nil
}

func (b *base) signal(m module, h bool) bool {
	return !h
}

func (b *base) addDestination(m module) {
	b.destinations = append(b.destinations, m)
}

func (b *base) Destinations() []module {
	return b.destinations
}

func (b *base) addSource(m module) {}

func (b *base) willSend() bool {
	return false
}
