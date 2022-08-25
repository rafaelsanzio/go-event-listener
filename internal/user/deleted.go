package user

import (
	"log"
	"time"

	"github.com/rafaelsanzio/go-event-listener/internal/pkg/event"
)

const Deleted event.Name = "user.deleted"

type DeletedEvent struct {
	ID   string
	Who  string
	Time time.Time
}

func (de DeletedEvent) Handle() {
	log.Printf("deleting: %+v\n", de)
}
