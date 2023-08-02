package main

import (
	"fmt"

	"github.com/asaskevich/EventBus"
)

func handleEvent(payload interface{}) {
    // Handle the event here
    // Cast the payload to the appropriate type if needed


      fmt.Println(payload.(string))

}


func  main(){


eventBus := EventBus.New()

// Subscribe the handler to an event
eventBus.Subscribe("eventName", handleEvent)

// Publish an event with payload
eventBus.Publish("eventName", "eventPayload")

// Unsubscribe the handler from the event
eventBus.Unsubscribe("eventName", handleEvent)

}