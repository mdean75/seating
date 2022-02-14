package attendeeservice

import (
	"seating/internal/app/domain"
	"seating/internal/app/ports"
)

type service struct {
	attendeeRepository ports.AttendeeRepository
}

func New(attendeeRepo ports.AttendeeRepository) *service {
	return &service{attendeeRepository: attendeeRepo}
}

func (s *service) CreateAttendee(name, companyName, industry, eventID string) (domain.Attendee, error) {
	attendee := domain.NewAttendee("", name, companyName, industry)
	id, err := s.attendeeRepository.Save(attendee, eventID)
	if err != nil {
		return domain.Attendee{}, err
	}

	attendee.ID = string(id)

	return attendee, nil
}

func (s *service) GetAttendee(attendeeID string) (domain.Attendee, error) {
	return s.attendeeRepository.Get(attendeeID)
}

func (s *service) DeleteAttendee(attendeeID string) error {
	return s.attendeeRepository.Delete(attendeeID)
}
