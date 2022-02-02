package ports

import "seating/internal/app/domain"

type ID string

type GroupRepository interface {
	// CreateGroup(displayName, shortName string) (ID, error)
	Save(domain.Group) (ID, error) // change to save()
}

type EventRepository interface {
	Save(domain.Event) (ID, error)
}