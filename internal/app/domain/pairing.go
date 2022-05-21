package domain

import (
	"fmt"
	"math/rand"
	"time"
)

type PairingRound struct {
	Pairs     []Pair
	Attendees map[string]Attendee
}

// NewPairingRound returns a PairingRound instance with a non-nil Pairs slice ready for use.
func NewPairingRound(attendees map[string]Attendee) (PairingRound, error) {
	//return PairingRound{Pairs: make([]Pair, 0)}
	a := PairingRound{Attendees: attendees}
	if len(a.Attendees) < 5 {
		return PairingRound{}, fmt.Errorf("unable to build seating chart, not enough attendees")
	}

	if len(a.Attendees)%2 != 0 {
		//a.Attendees = append(a.Attendees, Attendee{
		//	ID:   "placeholder",
		//	Name: "Placeholder",
		//})
		a.Attendees["placeholder"] = Attendee{ID: "placeholder", Name: "Placeholder"}
	}

	//tempAttendees := make([]Attendee, len(attendees))
	//num := copy(tempAttendees, attendees)  // copy attendees to c so we aren't modify the original attendees
	//fmt.Printf("number copied: %d \n", num)

	for ok := true; ok; ok = len(a.Attendees) > 2 {
		var p Pair

		// TODO: NOT THE GREATEST AT ALL!! REFACTOR TO ADD KEYS TO PAIRING STRUCT AND CONSTRUCTOR
		var keys []string
		for _, att := range a.Attendees {
			keys = append(keys, att.ID)
		}

		//p.Seat1 = a.arrayShift() // slice is updated
		p.Seat1 = a.Attendees[keys[0]] // select first attendee then remove from map
		delete(a.Attendees, p.Seat1.ID)
		p.Seat2 = a.selectPartner(p.Seat1, a.Attendees)

		// null out the paired with array so it's not recursive
		//p.Seat1.PairedWith = []Attendee{}
		//p.Seat2.PairedWith = []Attendee{}

		//p.Seat1.

		p.Seat1.PairedWith = append(p.Seat1.PairedWith, p.Seat2)
		p.Seat2.PairedWith = append(p.Seat2.PairedWith, p.Seat1)

		p.Seat1.PairedWithName = append(p.Seat1.PairedWithName, p.Seat2.Name)
		p.Seat1.PairedWithID = append(p.Seat1.PairedWithID, p.Seat2.ID)

		p.Seat2.PairedWithName = append(p.Seat2.PairedWithName, p.Seat1.Name)
		p.Seat2.PairedWithID = append(p.Seat2.PairedWithID, p.Seat1.ID)

		a.Pairs = append(a.Pairs, p)
		//addAttendeePairing(p.Seat1, p.Seat2, attendees)

	}

	var keys []string
	for _, att := range a.Attendees {
		keys = append(keys, att.ID)
	}

	// CHECK WE SHOULD ONLY HAVE 2 NOW
	if len(keys) != 2 {
		fmt.Println("expected key length 2, have: ", len(keys), keys)
	}

	//lastPair1, a.Attendees := arrayShift(a.Attendees)
	pp := Pair{
		Seat1: a.Attendees[keys[0]],
		Seat2: a.Attendees[keys[1]],
	}

	pp.Seat1.PairedWith = append(pp.Seat1.PairedWith, pp.Seat2)
	pp.Seat2.PairedWith = append(pp.Seat2.PairedWith, pp.Seat1)

	pp.Seat1.PairedWithName = append(pp.Seat1.PairedWithName, pp.Seat2.Name)
	pp.Seat2.PairedWithName = append(pp.Seat2.PairedWithName, pp.Seat1.Name)

	pp.Seat1.PairedWithID = append(pp.Seat1.PairedWithID, pp.Seat2.ID)
	pp.Seat2.PairedWithID = append(pp.Seat2.PairedWithID, pp.Seat1.ID)

	//a.Pairs = append(a.Pairs, Pair{
	//	Seat1: a.Attendees[keys[0]],
	//	Seat2: a.Attendees[keys[1]],
	//})
	a.Pairs = append(a.Pairs, pp)

	//addAttendeePairing(a.Attendees[0], a.Attendees[1], a.Attendees)
	return a, nil

}

//func (p *PairingRound) Generate(attendees []Attendee) ()

type Pair struct {
	Seat1 Attendee
	Seat2 Attendee
}

func NewPair(s1, s2 Attendee) Pair {
	return Pair{
		Seat1: s1,
		Seat2: s2,
	}
}

//func (a *PairingRound) arrayShift() Attendee {
//	t := a.Attendees[0]
//	a.Attendees = a.Attendees[1:]
//
//	return t
//}

func arrayPop(s []Attendee) []Attendee {
	index := len(s) - 1
	ret := make([]Attendee, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func (a *PairingRound) selectPartner(seat1 Attendee, c map[string]Attendee) Attendee {

	var seat2 Attendee
	//i := randomInt(0, len(d.Attendees))
	//Seat2 = d.peek(i)

	// get an array of the IDs so we can easily pick a random element using the index
	var keys []string
	for _, att := range a.Attendees {
		keys = append(keys, att.ID)
	}

	for {
		i := randomInt(0, len(keys))
		//seat2 = a.peek(i)
		seat2 = a.Attendees[keys[i]]

		if seat2.Industry != seat1.Industry {
			if seat1.PairedWith == nil {
				// remove from slice - swap to end and reslice

				//a.Attendees[i] = a.Attendees[len(a.Attendees)-1]
				//a.Attendees = a.Attendees[:len(a.Attendees)-1]

				// remove from attendees map
				delete(a.Attendees, seat2.ID)

				// remove from keys array
				keys[i] = keys[len(keys)-1]
				keys = keys[:len(keys)-1]

				return seat2
			}
			for _, v := range c {
				if v.ID == seat2.ID {
					if !arrayContains(seat1.ID, seat2.PairedWith) {
						// these two have not been paired before
						// remove from slice - swap to end and reslice
						//a.Attendees[i] = a.Attendees[len(a.Attendees)-1]
						//a.Attendees = a.Attendees[:len(a.Attendees)-1]

						// remove from attendees map
						delete(a.Attendees, seat2.ID)

						// remove from keys array
						keys[i] = keys[len(keys)-1]
						keys = keys[:len(keys)-1]

						//
						return seat2
					}

				}
			}

		}
	}
}

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

//func (a *PairingRound) peek(i int) Attendee {
//	return a.Attendees[i]
//}

func arrayContains(needle string, haystack []Attendee) bool {
	for _, v := range haystack {
		if needle == v.ID {
			return true
		}
	}

	return false
}

func addAttendeePairing(seat1 Attendee, seat2 Attendee, list []Attendee) {
	for i := range list {
		if list[i].ID == seat1.ID {
			list[i].PairedWith = append(list[i].PairedWith, seat2)
			//list[i].PairedWithName = append(list[i].PairedWithName, seat2.Name)
		}
		if list[i].ID == seat2.ID {
			list[i].PairedWith = append(list[i].PairedWith, seat1)
			//list[i].PairedWithName = append(list[i].PairedWithName, seat1.Name)
		}
	}
}
