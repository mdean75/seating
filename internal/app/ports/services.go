package ports

import "seating/internal/app/domain"

type GroupService interface {
	CreateGroup(displayName, shortName string) (domain.Group, error)
	GetGroup(groupID string) (domain.Group, error)
	DeleteGroup(groupID string) error
}

type EventService interface {
	CreateEvent(groupID string) (domain.Event, error)
	GetEvent(eventID string) (domain.Event, error)
	DeleteEvent(eventID string) error
}

type AttendeeService interface {
	CreateAttendee(name, companyName, industry string) (domain.Attendee, error)
}