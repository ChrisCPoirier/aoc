package day20

type flipFlop struct {
	name        string
	on          bool
	high        bool
	destination []module
	notify      chan notify
}

func (f *flipFlop) Name() string {
	return f.name
}

func (f *flipFlop) signal(src module, high bool) bool {
	if high {
		return false
	}

	f.high = high

	if !f.high {
		f.on = !f.on
	}

	return !f.high
}

func (f *flipFlop) pulse() []module {
	if f.high {
		return nil
	}

	next := []module{}
	for _, dest := range f.destination {
		if dest.signal(f, f.on) {
			next = append(next, dest)
		}

		f.notify <- notify{f.Name(), f.on}
	}
	return next
}

func (f *flipFlop) addDestination(m module) {
	// m.addSource(f)
	f.destination = append(f.destination, m)
}

func (f *flipFlop) addSource(m module) {}

func (b *flipFlop) Destinations() []module {
	return b.destination
}

func (f *flipFlop) willSend() bool {
	return f.on
}
