package domain

type Group struct {
	ID string `bson:"_id,omitempty"`
	DisplayName string `bson:"displayName"`
	ShortName string `bson:"shortName"`
}