package user

import (
	"log"
	"time"

	"github.com/rafaelsanzio/go-event-listener/internal/pkg/event"
)

const Updated event.Name = "user.updated"

type UpdatedEvent struct {
	ID       string
	Key      string
	OldValue string
	NewValue string
	Time     time.Time
}

func (ue UpdatedEvent) Handle() {
	log.Printf("updating: %+v\n", ue)
}
