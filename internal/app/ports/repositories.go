package ports

import "seating/internal/app/domain"

type ID string

type GroupRepository interface {
	Save(domain.Group) (ID, error)
	Get(string) (domain.Group, error)
	GetAll() ([]domain.Group, error)
	Delete(string) error
}

type EventRepository interface {
	Save(domain.Event) (ID, error)
	Get(string) (domain.Event, error)
	Delete(string) error
	GetListCount(string) (int, error)
	GetEventsForGroup(string) ([]domain.Event, error)
	PairingRepository
}

type PairingRepository interface {
	SaveRound(string, []domain.Pair) error
}

type AttendeeRepository interface {
	Save(domain.Attendee, string) (ID, error)
	Get(string) (domain.Attendee, error)
	Delete(string) error
}

type IndustryRepository interface {
	Save(domain.Industry) (ID, error)
	Get(string) (domain.Industry, error)
	Delete(string) error
}
