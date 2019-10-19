package main

type attendee struct {
	name       string
	id         int
	industry   string
	pairedWith []int
}

type pair struct {
	seat1 attendee
	seat2 attendee
}
