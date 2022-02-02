package eventadapter

import "time"

type ID string

type Event struct {
	ID string `json:"id,omitempty"`
	Date time.Time `json:"date"`
	GroupID string `json:"groupId"`
}

func NewEventRequest(groupID ID) Event {
	return Event{
		GroupID: string(groupID),
		Date: time.Now(),
	}
}
