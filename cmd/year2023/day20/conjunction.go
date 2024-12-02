package day20

type conjunction struct {
	name        string
	memory      map[module]bool
	destination []module
	notify      chan notify
}

func (c *conjunction) Name() string {
	return c.name
}

func (c *conjunction) signal(src module, high bool) bool {
	c.memory[src] = high

	return true
}

func (c *conjunction) addDestination(m module) {
	c.destination = append(c.destination, m)
}

func (c *conjunction) addSource(m module) {

	if c.memory == nil {
		c.memory = map[module]bool{}
	}

	c.memory[m] = false
}

func (c *conjunction) pulse() []module {
	high := false

	for _, h := range c.memory {
		if !h {
			high = true
			break
		}
	}

	next := []module{}
	for _, dest := range c.destination {
		if dest.signal(c, high) {
			next = append(next, dest)
		}
		c.notify <- notify{c.Name(), high}
	}

	return next
}

func (c *conjunction) willSend() bool {
	high := false

	for _, h := range c.memory {
		if !h {
			high = true
			break
		}
	}

	return high
}

func (b *conjunction) Destinations() []module {
	return b.destination
}
