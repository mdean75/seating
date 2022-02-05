package domain

import "reflect"

// TODO: I need to decide how to populate the shortname and what it will be used for. Should this be generated somehow here?

type Group struct {
	ID          string
	DisplayName string
	ShortName   string
}

func NewGroup(id, displayName, shortName string) Group {
	return Group{
		ID:          id,
		DisplayName: displayName,
		ShortName:   shortName,
	}
}

func (g Group) IsZero() bool {
	return reflect.DeepEqual(g, Group{})
}
