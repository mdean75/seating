package groupservice

import (
	"seating/internal/app/domain"
	"seating/internal/app/ports"
)

type service struct {
	groupRepository ports.GroupRepository
}

func New(groupRepo ports.GroupRepository)*service {
	return &service{groupRepository: groupRepo}
}

func (s *service) CreateGroup(displayName, shortName string) (domain.Group, error) {
	group := domain.NewGroup("", displayName, shortName)

	id, err := s.groupRepository.Save(group)
	if err != nil {
		return domain.Group{}, err
	}

	group.ID = string(id)
	return group, nil
}

func (s *service) GetGroup(groupID string) (domain.Group, error) {
	group, err := s.groupRepository.Get(groupID)
	if err != nil {
		return domain.Group{}, nil
	}

	return group, nil
}