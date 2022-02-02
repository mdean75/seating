package eventservice

import "seating/internal/app/ports"

type service struct {
	eventRepository ports.EventRepository
}

func New(eventRepo ports.EventRepository) *service {
	return &service{eventRepository: eventRepo}
}

func (s*service) CreateEvent(groupID string) (string, error) {
	id, err := s.eventRepository.CreateEvent(ports.ID(groupID))
	if err != nil {
		return "", err
	}

	return string(id), nil
}