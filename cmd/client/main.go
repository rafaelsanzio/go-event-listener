package main

import (
	"log"
	"time"

	"github.com/rafaelsanzio/go-event-listener/internal/pkg/event"
	"github.com/rafaelsanzio/go-event-listener/internal/user"
)

func main() {
	// Register events with listeners at application boot
	dispatcher := event.NewDispatcher()
	if err := dispatcher.Register(user.Listener{}, user.Created, user.Updated, user.Deleted); err != nil {
		log.Fatalln(err)
	}

	// Dispatch registered events. Valid.
	go func() {
		err := dispatcher.Dispatch(user.Created, user.CreatedEvent{
			Time: time.Now().UTC(),
			ID:   "any_id",
		})
		if err != nil {
			log.Println(err)
		}
	}()

	go func() {
		err := dispatcher.Dispatch(user.Updated, user.UpdatedEvent{
			Time:     time.Now().UTC(),
			ID:       "other_id",
			Key:      "name",
			OldValue: "rafael",
			NewValue: "sanzio",
		})
		if err != nil {
			log.Println(err)
		}
	}()

	go func() {
		err := dispatcher.Dispatch(user.Deleted, user.DeletedEvent{
			Time: time.Now().UTC(),
			ID:   "another_id",
			Who:  "admin",
		})
		if err != nil {
			log.Println(err)
		}
	}()

	// Dispatch a wrong event type to registered event name. Error.
	go dispatcher.Dispatch(user.Created, nil)
	go dispatcher.Dispatch(user.Updated, "hi")
	go dispatcher.Dispatch(user.Created, 123)
	go dispatcher.Dispatch(user.Updated, struct{}{})
	go dispatcher.Dispatch(user.Created, make(chan int))

	select {}
}
