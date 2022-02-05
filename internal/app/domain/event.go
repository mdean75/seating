package domain

import (
	"reflect"
	"time"
)

type Event struct {
	ID      string
	Date    time.Time
	GroupID string
	Group   Group
}

func NewEvent(id, groupID string, date time.Time, group Group) Event {
	return Event{
		ID:      id,
		Date:    date,
		GroupID: groupID,
		Group:   group,
	}
}

func (e Event) IsZero() bool {
	return reflect.DeepEqual(e, Event{})
}
