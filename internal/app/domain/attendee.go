package domain

type Attendee struct {
	ID             string
	Name           string
	CompanyName    string
	Industry       string
	PairedWith     []Attendee
	PairedWithID   []string
	PairedWithName []string
}

func NewAttendee(id, name, company, industry string) Attendee {
	return Attendee{
		ID:          id,
		Name:        name,
		CompanyName: company,
		Industry:    industry,
	}
}
