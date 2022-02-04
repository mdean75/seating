package attendeeadapter

import "seating/internal/app/domain"

type Attendee struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name"`
	CompanyName string `json:"companyName"`
	Industry string `json:"industry"`
}

func convertJSONAttendeeFromDomain(attendee domain.Attendee) Attendee {
	return Attendee{
		ID: attendee.ID,
		Name: attendee.Name,
		CompanyName: attendee.CompanyName,
		Industry: attendee.Industry,
	}
}