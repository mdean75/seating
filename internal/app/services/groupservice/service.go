package groupservice

import "seating/internal/app/ports"

type service struct {
	groupRepository ports.GroupRepository
}

func New(groupRepo ports.GroupRepository)*service {
	return &service{groupRepository: groupRepo}
}

func (s *service) CreateGroup(displayName, shortName string) (string, error) {
	id, err := s.groupRepository.CreateGroup(displayName, shortName)
	if err != nil {
		return "", err
	}

	return string(id), nil
}