package eventadapter

import (
	"seating/internal/app/domain"
	"seating/internal/handlers/groupadapter"
	"time"
)

type ID string

type Event struct {
	ID      string             `json:"id,omitempty"`
	Date    time.Time          `json:"date"`
	GroupID string             `json:"groupId"`
	Group   groupadapter.Group `json:"group"`
}

// func NewEventRequest(groupID ID) Event {
// 	return Event{
// 		GroupID: string(groupID),
// 		Date: time.Now(),
// 	}
// }

func ConvertJSONEventFromDomain(event domain.Event) Event {
	return Event{
		ID:      event.ID,
		Date:    event.Date,
		GroupID: event.GroupID,
		Group:   groupadapter.ConvertJSONGroupFromDomain(event.Group),
	}
}
