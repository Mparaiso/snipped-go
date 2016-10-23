package smartsnippets

type Event interface{}

// Signal is an implementation of the signal pattern
type Signal interface {
	Add(Listener)
	Remove(Listener)
	Dispatch(data Event) error
}

// Listener handle signals
type Listener interface {
	Handle(data Event) error
}

type funcListener struct {
	Listener func(data Event) error
}

func (fl funcListener) Handle(event Event) error {
	return fl.Listener(event)
}

func ListenerFunc(f func(data Event) error) *funcListener {
	return &funcListener{f}
}

type DefaultSignal struct {
	Listeners []Listener
}

func NewDefaultSignal() *DefaultSignal {
	return &DefaultSignal{Listeners: []Listener{}}
}

func (signal *DefaultSignal) Add(l Listener) {
	if signal.IndexOf(l) != -1 {
		return
	}
	signal.Listeners = append(signal.Listeners, l)
}

func (signal *DefaultSignal) IndexOf(l Listener) int {
	for i, listener := range signal.Listeners {
		if listener == l {
			return i
		}
	}
	return -1
}

func (signal *DefaultSignal) Remove(l Listener) {
	index := signal.IndexOf(l)
	if index == -1 {
		return
	}
	head := signal.Listeners[:index]
	if index == len(signal.Listeners)-1 {
		signal.Listeners = head
	} else {
		signal.Listeners = append(head, signal.Listeners[index+1:]...)
	}
}

func (signal *DefaultSignal) Dispatch(data Event) error {
	for _, listener := range signal.Listeners {
		if err := listener.Handle(data); err != nil {
			return err
		}
	}
	return nil
}
