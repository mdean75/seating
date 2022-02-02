package eventservice

import (
	"seating/internal/app/domain"
	"seating/internal/app/ports"
	"time"
)

type service struct {
	eventRepository ports.EventRepository
}

func New(eventRepo ports.EventRepository) *service {
	return &service{eventRepository: eventRepo}
}

func (s*service) CreateEvent(groupID string) (domain.Event, error) {
	event := domain.NewEvent("", groupID, time.Now())
	id, err := s.eventRepository.Save(event)
	if err != nil {
		return domain.Event{}, err
	}

	event.ID = string(id)

	return event, nil
}