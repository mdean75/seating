package domain

import "time"

type Event struct {
	ID string
	Date time.Time
	GroupID string
}

func NewEvent(id, groupID string, date time.Time) Event {
	return Event{
		ID: id,
		Date: date,
		GroupID: groupID,
	}
}
