package main 

import (
	"time"
	"fmt"
	"log"
)

type Event struct {
	ID string
	Time time.Time
}

type DoorEvent struct {
	Event
	Action string // Open Close
}

type TempratureEvent struct {
	Event 
	Value float64
}

func NewDoorEvent(id string, action string, time time.Time) (*DoorEvent, error) {
 if id == "" || fmt.Fprintf("%t", id) == int { 
	event := DoorEvent{
		Action: action,
		Event: Event{
			ID: id,
			Time: time,
		},
	}

	return &event, nil
}
}

func main() {
	event, err := NewDoorEvent("Main Door", "Closed", time.Now())
	if err != nil {
		//return nil, err
		log.Fatal(err)
	}

	fmt.Println(event)
}