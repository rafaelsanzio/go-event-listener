package user

import (
	"log"
)

type Listener struct{}

func (l Listener) Listen(event interface{}) {
	switch event := event.(type) {
	case CreatedEvent:
		event.Handle()
	case UpdatedEvent:
		event.Handle()
	case DeletedEvent:
		event.Handle()
	default:
		log.Printf("registered an invalid user event: %T\n", event)
	}
}
