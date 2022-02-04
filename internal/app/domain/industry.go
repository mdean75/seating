package domain

type Industry struct {
	ID string
	Name string
}

func NewIndustry(id, name string) Industry {
	return Industry{
		ID: id,
		Name: name,
	}
}