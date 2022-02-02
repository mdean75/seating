package ports

type ID string

type GroupRepository interface {
	CreateGroup(displayName, shortName string) (ID, error)
}

type EventRepository interface {
	CreateEvent(groupID ID, ) (ID, error)
}