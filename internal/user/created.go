package user

import (
	"log"
	"time"

	"github.com/rafaelsanzio/go-event-listener/internal/pkg/event"
)

const Created event.Name = "user.created"

type CreatedEvent struct {
	ID   string
	Time time.Time
}

func (ce CreatedEvent) Handle() {
	log.Printf("creating: %+v\n", ce)
}
