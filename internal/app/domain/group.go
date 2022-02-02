package domain

type Group struct {
	ID string 
	DisplayName string 
	ShortName string 
}

func NewGroup(id, displayName, shortName string) Group {
	return Group{
		ID: id,
		DisplayName: displayName,
		ShortName: shortName,
	}
}