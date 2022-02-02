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

	// id, err := s.groupRepository.CreateGroup(displayName, shortName)
	id, err := s.groupRepository.Save(group)
	if err != nil {
		return domain.Group{}, err
	}

	group.ID = string(id)
	return group, nil
}