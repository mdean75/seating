package industryadapter

import "seating/internal/app/domain"

type Industry struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

func convertJSONIndustryFromDomain(industry domain.Industry) Industry {
	return Industry{
		ID:   industry.ID,
		Name: industry.Name,
	}
}
