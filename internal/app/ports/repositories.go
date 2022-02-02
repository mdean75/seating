package ports

import "seating/internal/app/domain"

type ID string

type GroupRepository interface {
	Save(domain.Group) (ID, error)
}

type EventRepository interface {
	Save(domain.Event) (ID, error)
}