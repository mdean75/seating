package ports

import (
	"seating/internal/app/domain"
	"time"
)

type GroupService interface {
	CreateGroup(displayName, shortName string) (domain.Group, error)
	GetGroup(groupID string) (domain.Group, error)
	GetAllGroups() ([]domain.Group, error)
	DeleteGroup(groupID string) error
}

type EventService interface {
	CreateEvent(groupID string, group domain.Group, date time.Time) (domain.Event, error)
	GetEvent(eventID string) (domain.Event, error)
	DeleteEvent(eventID string) error
	CreatePairingRound(eventID string, attendees []domain.Pair) error
	GetListCount(eventID string) (int, error)
	GetEventsForGroup(groupID string) ([]domain.Event, error)
}

type AttendeeService interface {
	CreateAttendee(name, companyName, industry, eventID string) (domain.Attendee, error)
	GetAttendee(attendeeID string) (domain.Attendee, error)
	DeleteAttendee(attendeeID string) error
}

type IndustryService interface {
	CreateIndustry(name string) (domain.Industry, error)
	GetIndustry(industryID string) (domain.Industry, error)
	DeleteIndustry(industryID string) error
}
