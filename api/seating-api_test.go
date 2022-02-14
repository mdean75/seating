package api

import "testing"

func TestAppData_generateID(t *testing.T) {
	type fields struct {
		Industries []string
		Attendees  []Attendee
		Pairs      []Pair
		ListCount  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: "test 1", fields: fields{
			Industries: nil,
			Attendees: []Attendee{{
				Name:           "george",
				ID:             6,
				Industry:       "",
				Business:       "",
				PairedWith:     nil,
				PairedWithName: nil,
			}, {Name: "larry", ID: 6}},
			Pairs:     nil,
			ListCount: 0,
		}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AppData{
				Industries: tt.fields.Industries,
				Attendees:  tt.fields.Attendees,
				Pairs:      tt.fields.Pairs,
				ListCount:  tt.fields.ListCount,
			}
			if got := a.generateID(); got != tt.want {
				t.Errorf("generateID() = %v, want %v", got, tt.want)
			}
		})
	}
}
