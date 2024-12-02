package day20

type broadcaster struct {
	name         string
	destinations []module
	notify       chan notify
}

func (b *broadcaster) Name() string {
	return b.name
}

func (b *broadcaster) pulse() []module {
	for _, dest := range b.destinations {
		dest.signal(b, false)
		b.notify <- notify{b.Name(), false}
	}
	return b.destinations
}

func (b *broadcaster) signal(module, bool) bool {
	return false
}

func (b *broadcaster) addDestination(m module) {
	b.destinations = append(b.destinations, m)
}

func (b *broadcaster) Destinations() []module {
	return b.destinations
}

func (b *broadcaster) addSource(m module) {}

func (b *broadcaster) willSend() bool {
	return false
}
