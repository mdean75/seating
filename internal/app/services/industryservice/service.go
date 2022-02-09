package industryservice

import (
	"seating/internal/app/domain"
	"seating/internal/app/ports"
)

type service struct {
	industryRepository ports.IndustryRepository
}

func New(industryRepo ports.IndustryRepository) *service {
	return &service{industryRepository: industryRepo}
}

func (s *service) CreateIndustry(name string) (domain.Industry, error) {
	industry := domain.NewIndustry("", name)

	id, err := s.industryRepository.Save(industry)
	if err != nil {
		return domain.Industry{}, err
	}

	industry.ID = string(id)
	return industry, nil
}

func (s *service) GetIndustry(industryID string) (domain.Industry, error) {
	group, err := s.industryRepository.Get(industryID)
	if err != nil {
		return domain.Industry{}, err
	}

	return group, nil
}

func (s *service) DeleteIndustry(industryID string) error {
	return s.industryRepository.Delete(industryID)
}
