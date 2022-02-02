package groupadapter

type Group struct {
	ID string `json:"id,omitempty" bson:"_id,omitempty"`
	DisplayName string `json:"displayName"`
	ShortName string `json:"shortName"`
}

func NewGroupRequest(displayName, shortName string) Group {
	return Group{
		DisplayName: displayName,
		ShortName: shortName,
	}
}