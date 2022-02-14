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

func (s *service) GetListCount(eventID string) (int, error) {
	return s.eventRepository.GetListCount(eventID)

}

func (s *service) CreateEvent(groupID string, group domain.Group, date time.Time) (domain.Event, error) {
	// TODO: We will want this to be not the current time but rather the date of the event
	event := domain.NewEvent("", groupID, date, group)
	id, err := s.eventRepository.Save(event)
	if err != nil {
		return domain.Event{}, err
	}

	event.ID = string(id)

	return event, nil
}

func (s *service) GetEvent(eventID string) (domain.Event, error) {
	event, err := s.eventRepository.Get(eventID)
	if err != nil {
		return domain.Event{}, err
	}

	return event, nil
}

func (s *service) DeleteEvent(eventID string) error {
	return s.eventRepository.Delete(eventID)
}

func (s *service) CreatePairingRound(eventID string, pairs []domain.Pair) error {
	return s.eventRepository.SaveRound(eventID, pairs)
}

func (s *service) GetEventsForGroup(groupID string) ([]domain.Event, error) {
	return s.eventRepository.GetEventsForGroup(groupID)
}
