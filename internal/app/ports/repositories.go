package ports

import "seating/internal/app/domain"

type ID string

type GroupRepository interface {
	Save(domain.Group) (ID, error)
	Get(string) (domain.Group, error)
}

type EventRepository interface {
	Save(domain.Event) (ID, error)
	Get(string) (domain.Event, error)
	Delete(string) error
}