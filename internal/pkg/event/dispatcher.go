package event

import (
	"fmt"
)

type Dispatcher struct {
	Jobs   chan Job
	Events map[Name]Listener
}

func NewDispatcher() *Dispatcher {
	d := &Dispatcher{
		Jobs:   make(chan Job),
		Events: make(map[Name]Listener),
	}

	go d.consume()

	return d
}

func (d *Dispatcher) Register(listener Listener, names ...Name) error {
	for _, name := range names {
		if _, ok := d.Events[name]; ok {
			return fmt.Errorf("the '%s' event is already registered", name)
		}

		d.Events[name] = listener
	}

	return nil
}

func (d *Dispatcher) Dispatch(name Name, event interface{}) error {
	if _, ok := d.Events[name]; !ok {
		return fmt.Errorf("the '%s' event is not registered", name)
	}

	d.Jobs <- Job{EventName: name, EventType: event}

	return nil
}

func (d *Dispatcher) consume() {
	for job := range d.Jobs {
		d.Events[job.EventName].Listen(job.EventType)
	}
}
