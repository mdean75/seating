package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type data struct {
	Industries []string
	Attendees  []Attendee
	Pairs      []pair
}

type Attendee struct {
	name       string
	id         int
	industry   string
	pairedWith []int
}

type pair struct {
	seat1 Attendee
	seat2 Attendee
}

func main() {
	//var Attendees []Attendee
	//var Industries []string
	var d data

	d.setIndustries(&d.Industries)
	scanner := bufio.NewScanner(os.Stdin)
	var text string
	for strings.ToUpper(text) != "Q" {
		displayMenu()
		scanner.Scan()
		text = scanner.Text()
		d.processInput(text)
	}

}

func (d *data) setIndustries(indust *[]string) {
	*indust = append(*indust, "Finance", "Construction", "Food and Beverage", "Entertainment")
}

func (d *data) processInput(action string) {
	switch action {
	case "1":
		d.addAttendee()

	case "2":
		d.displayAttendeeList()

	case "3":
		d.buildChart()
	}
}

func (d *data) addAttendee() {
	var att Attendee

	fmt.Println("Add Attendee")
	fmt.Println("\nName: \t\t")
	att.name = readInput()
	d.displayIndustries()
	att.industry = d.getIndustry()

	att.id = randomInt(1, 1000)
	d.Attendees = append(d.Attendees, att)
}

func (d *data) getIndustry() string {
	input := readInput()
	i, _ := strconv.Atoi(input)
	return d.Industries[i]
}

func (d *data) displayIndustries() {
	fmt.Println("Select Attendees' Industry")
	for k, v := range d.Industries {
		fmt.Println(k, ": ", v)
	}
}

func (d *data) displayAttendeeList() {
	for _, a := range d.Attendees {
		fmt.Println(a)
	}
}

func (d *data) buildChart() {
	// ensure an even pairing
	if len(d.Attendees)%2 != 0 {
		d.Attendees = append(d.Attendees, Attendee{name: "Placeholder"})
	}

	c := make([]Attendee, len(d.Attendees))
	copy(c, d.Attendees)

	for ok := true; ok; ok = len(d.Attendees) > 2 {
		var p pair

		p.seat1 = d.shiftArray()
		p.seat2 = d.selectPartner(p.seat1, c)

		d.Pairs = append(d.Pairs, p)
		addAttendeePairing(p.seat1.id, p.seat2.id, c)
	}

	// select last two no matter the match
	lastPair1 := d.shiftArray()
	lastPair2 := d.shiftArray()
	d.Pairs = append(d.Pairs, pair{
		seat1: lastPair1,
		seat2: lastPair2,
	})

	addAttendeePairing(lastPair1.id, lastPair2.id, c)

	fmt.Println(d.Pairs)
	fmt.Println("\n\n", c)

	// the slice should be nil at this point, reload original data
	d.Attendees = c
}

func addAttendeePairing(seat1 int, seat2 int, list []Attendee) {
	for i := range list {
		if list[i].id == seat1 {
			list[i].pairedWith = append(list[i].pairedWith, seat2)
		}
		if list[i].id == seat2 {
			list[i].pairedWith = append(list[i].pairedWith, seat1)
		}
	}
}

func (d *data) shiftArray() Attendee {
	t := d.Attendees[0]
	d.Attendees = d.Attendees[1:]

	return t
}

func (d *data) selectPartner(seat1 Attendee, c []Attendee) Attendee {
	var seat2 Attendee
	//i := randomInt(0, len(d.Attendees))
	//seat2 = d.peek(i)
	for {
		i := randomInt(0, len(d.Attendees))
		seat2 = d.peek(i)

		if seat2.industry != seat1.industry {
			if seat1.pairedWith == nil {
				// remove from slice - swap to end and reslice
				d.Attendees[i] = d.Attendees[len(d.Attendees)-1]
				d.Attendees = d.Attendees[:len(d.Attendees)-1]

				return seat2
			}
			for _, v := range c {
				if v.id == seat2.id {
					if !arrayContains(seat1.id, seat2.pairedWith) {
						// these two have not been paired before
						// remove from slice - swap to end and reslice
						d.Attendees[i] = d.Attendees[len(d.Attendees)-1]
						d.Attendees = d.Attendees[:len(d.Attendees)-1]

						return seat2
					}

				}
			}

		}
	}
	//if seat2.industry != seat1.industry {
	//	// remove from slice - swap to end and reslice
	//	d.Attendees[i] = d.Attendees[len(d.Attendees)-1]
	//	d.Attendees = d.Attendees[:len(d.Attendees)-1]
	//}
	return seat2

}

func arrayContains(needle int, haystack []int) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}

	return false
}

func (d *data) peek(i int) Attendee {
	return d.Attendees[i]
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func displayMenu() {
	fmt.Println("Select a menu option")
	fmt.Println("\nEnter Attendee \t\t\t ... 1")
	fmt.Println("\nDisplay list of Attendees \t ... 2")
	fmt.Println("\nCreate seating charts \t\t ... 3")
	fmt.Println("\nQuit \t\t\t\t ... Q")
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
