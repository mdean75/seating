/**
filename:	common.go
purpose:	provides common functions for the application api and original app
update:		3/3/2020
comments:	all of the functions that are common between the original app and the api.
*/

package api

import (
	"fmt"
	"math/rand"
	"time"
)

func addAttendeePairing(seat1 Attendee, seat2 Attendee, list []Attendee) {
	for i := range list {
		if list[i].ID == seat1.ID {
			list[i].PairedWith = append(list[i].PairedWith, seat2.ID)
			list[i].PairedWithName = append(list[i].PairedWithName, seat2.Name)
		}
		if list[i].ID == seat2.ID {
			list[i].PairedWith = append(list[i].PairedWith, seat1.ID)
			list[i].PairedWithName = append(list[i].PairedWithName, seat1.Name)
		}
	}
}

func arrayContains(needle int, haystack []int) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}

	return false
}

func (a *AppData) peek(i int) Attendee {
	return a.Attendees[i]
}

func (a *AppData) selectPartner(seat1 Attendee, c []Attendee) Attendee {
	var seat2 Attendee
	//i := randomInt(0, len(d.Attendees))
	//Seat2 = d.peek(i)
	for {
		i := randomInt(0, len(a.Attendees))
		seat2 = a.peek(i)

		if seat2.Industry != seat1.Industry {
			if seat1.PairedWith == nil {
				// remove from slice - swap to end and reslice
				a.Attendees[i] = a.Attendees[len(a.Attendees)-1]
				a.Attendees = a.Attendees[:len(a.Attendees)-1]

				return seat2
			}
			for _, v := range c {
				if v.ID == seat2.ID {
					if !arrayContains(seat1.ID, seat2.PairedWith) {
						// these two have not been paired before
						// remove from slice - swap to end and reslice
						a.Attendees[i] = a.Attendees[len(a.Attendees)-1]
						a.Attendees = a.Attendees[:len(a.Attendees)-1]

						return seat2
					}

				}
			}

		}
	}
}

func (a *AppData) arrayShift() Attendee {
	t := a.Attendees[0]
	a.Attendees = a.Attendees[1:]

	return t
}

func arrayPop(s []Attendee) []Attendee {
	index := len(s) - 1
	ret := make([]Attendee, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func (a *AppData) generateID() int {
	var i int
	var b bool

	for !b {
		i = randomInt(1, 1000)
		b = a.isUniqueID(i)
		fmt.Println()
	}
	return i
}

func (a *AppData) isUniqueID(id int) bool {
	for _, attendee := range a.Attendees {
		if attendee.ID == id {
			return false // this id already exists, return false
		}
	}
	return true
}

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
