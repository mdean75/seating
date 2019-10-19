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
	Attendees  []attendee
}

func main() {
	//var Attendees []attendee
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
	//3displayMenu()
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
	}
}

func (d *data) addAttendee() {
	var att attendee

	fmt.Println("Add Attendee")
	fmt.Println("\nName: \t\t")
	att.name = readInput()
	d.displayIndustries()
	att.industry = d.getIndustry()
	fmt.Println("\n")
	att.id = randomInt(1, 100)
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
