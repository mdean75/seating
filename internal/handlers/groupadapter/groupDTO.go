package groupadapter

import "seating/internal/app/domain"

type Group struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	ShortName   string `json:"shortName"`
}

// func NewGroupRequest(displayName, shortName string) Group {
// 	return Group{
// 		DisplayName: displayName,
// 		ShortName: shortName,
// 	}
// }

func ConvertJSONGroupFromDomain(group domain.Group) Group {
	return Group{
		ID:          group.ID,
		DisplayName: group.DisplayName,
		ShortName:   group.ShortName,
	}
}

func ConvertDomainGroupFromJSON(group Group) domain.Group {
	return domain.NewGroup(group.ID, group.DisplayName, group.ShortName)
}
