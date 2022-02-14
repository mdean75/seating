package eventadapter

import (
	"seating/internal/app/domain"
	"seating/internal/handlers/attendeeadapter"
	"seating/internal/handlers/groupadapter"
	"time"
)

type ID string

type Event struct {
	ID           string                     `json:"id,omitempty"`
	Date         time.Time                  `json:"date"`
	GroupID      string                     `json:"groupId,omitempty"`
	Group        groupadapter.Group         `json:"group,omitempty"`
	Attendees    []attendeeadapter.Attendee `json:"attendees,omitempty"`
	PairingRound [][]Pair                   `json:"pairingRound,omitempty"`
}

// func NewEventRequest(groupID ID) Event {
// 	return Event{
// 		GroupID: string(groupID),
// 		Date: time.Now(),
// 	}
// }

type PairingRound struct {
	Pairs     []Pair
	Attendees []attendeeadapter.Attendee
}

type Pair struct {
	Seat1 attendeeadapter.Attendee
	Seat2 attendeeadapter.Attendee
}

func ConvertJSONEventFromDomain(event domain.Event) Event {
	var domAttendees []attendeeadapter.Attendee
	for _, attendee := range event.Attendees {
		var pairedWithAttendee []attendeeadapter.Attendee
		for _, p := range attendee.PairedWith {
			pairedWithAttendee = append(pairedWithAttendee, attendeeadapter.Attendee{
				ID:          p.ID,
				Name:        p.Name,
				CompanyName: p.CompanyName,
				Industry:    p.Industry,
				//PairedWith:  nil,
			})
		}

		domAttendees = append(domAttendees, attendeeadapter.Attendee{
			ID:          attendee.ID,
			Name:        attendee.Name,
			CompanyName: attendee.CompanyName,
			Industry:    attendee.Industry,
			PairedWith:  pairedWithAttendee,
		})
		//.NewAttendee(attendee.ID, attendee.Name, attendee.CompanyName, attendee.Industry))
	}

	var jsonRounds [][]Pair
	for _, round := range event.PairingRound {
		var seatinground []Pair
		for _, p := range round.Pairs {
			seat1 := attendeeadapter.Attendee{
				ID:          p.Seat1.ID,
				Name:        p.Seat1.Name,
				CompanyName: p.Seat1.CompanyName,
				Industry:    p.Seat1.Industry,
			}

			seat2 := attendeeadapter.Attendee{
				ID:          p.Seat2.ID,
				Name:        p.Seat2.Name,
				CompanyName: p.Seat2.CompanyName,
				Industry:    p.Seat2.Industry,
			}

			var pp Pair
			pp.Seat1 = seat1
			pp.Seat2 = seat2
			seatinground = append(seatinground, pp)
		}
		jsonRounds = append(jsonRounds, seatinground)
	}

	return Event{
		ID:           event.ID,
		Date:         event.Date,
		GroupID:      event.GroupID,
		Group:        groupadapter.ConvertJSONGroupFromDomain(event.Group),
		Attendees:    domAttendees,
		PairingRound: jsonRounds,
	}
}
