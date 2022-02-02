package ports

type GroupService interface {
	CreateGroup(displayName, shortName string) (string, error)
}

type EventService interface {
	CreateEvent(groupID string) (string, error)
}