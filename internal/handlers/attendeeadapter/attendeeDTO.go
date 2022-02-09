package attendeeadapter

import "seating/internal/app/domain"

type Attendee struct {
	ID          string     `json:"id,omitempty"`
	Name        string     `json:"name"`
	CompanyName string     `json:"companyName"`
	Industry    string     `json:"industry"`
	PairedWith  []Attendee `json:"pairedWith"`
}

func convertJSONAttendeeFromDomain(attendee domain.Attendee) Attendee {
	var jsonPairedWith []Attendee

	for _, a := range attendee.PairedWith {
		jsonPairedWith = append(jsonPairedWith, Attendee{
			ID:          a.ID,
			Name:        a.Name,
			CompanyName: a.CompanyName,
			Industry:    a.Industry,
			//PairedWith:  a.,
		})
	}

	return Attendee{
		ID:          attendee.ID,
		Name:        attendee.Name,
		CompanyName: attendee.CompanyName,
		Industry:    attendee.Industry,
		PairedWith:  jsonPairedWith,
	}
}
