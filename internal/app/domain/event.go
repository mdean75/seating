package domain

import (
	"reflect"
	"time"
)

type Event struct {
	ID           string
	Date         time.Time
	GroupID      string
	Group        Group
	Attendees    []Attendee // new event constructor does not include attendees as they are added later usually one at a time
	PairingRound []PairingRound
}

func NewEvent(id, groupID string, date time.Time, group Group, attendees ...Attendee) Event {
	return Event{
		ID:           id,
		Date:         date,
		GroupID:      groupID,
		Group:        group,
		Attendees:    attendees,
		PairingRound: make([]PairingRound, 0),
	}
}

func (e Event) IsZero() bool {
	return reflect.DeepEqual(e, Event{})
}
