/**
filename:	seating-app.go
purpose:	provides functions for handling the go html based app
update:		3/3/2020
comments:	all of the functions specific to the original go html template app have been moved to this file and
			are considered deprecated.  All future work will focus on converting this app to an api only.
*/

package api

import (
	"context"
	"fmt"
	"github.com/rivo/sessions"
	"html/template"
	"log"
	"net/http"
	"seating/app"
	"time"
)

// AttendeeEntry is the handler to display the user input form.
func (a *AppData) AttendeeEntry(w http.ResponseWriter, r *http.Request) {

	// start the session but do not create a new session if one does not exist as this could indicate either someone mistakenly routed to this endpoint or some kind of attack.
	session, err := sessions.Start(w, r, false)

	var tData inputData
	// check if there is a session, then get the session message and immediately delete it, if not present set default message
	if session != nil {
		msg := session.GetAndDelete("successMsg", "")
		err = session.Destroy(w, r)
		//if err != nil {
		//	// there was an error deleting the session, redirect to main page
		//	http.Redirect(w, r, "/", http.StatusSeeOther)
		//
		//	log.Error(err.Error())
		//	return
		//}
		tData.SuccessMsg = fmt.Sprintf("%v", msg)

	}

	tData.Industries = a.Industries
	tData.ListCount = a.ListCount

	err = loadForm(app.InputForm, w, tData)
	if err != nil {
		//log.Debug("error in AttendeeEntry() received from LoadForm()", err.Error())
	}

}

// loadForm is a helper function to handling parsing and displaying the forms.
func loadForm(file string, w http.ResponseWriter, data interface{}) error {

	// parse the template file
	t, err := template.New("html").Parse(file)
	if err != nil {
		log.Println("template parsing error: ", err)
		return err
	}

	// load the form
	err = t.Execute(w, data)
	if err != nil {
		log.Print("template executing error: ", err)
		return err
	}

	return nil

}

func (a *AppData) ProcessAttendeeEntry(w http.ResponseWriter, r *http.Request) {

	// start new session and create cookie
	session, err := sessions.Start(w, r, true)

	err = r.ParseForm()
	if err != nil {
		// do something
	}

	name := r.Form.Get("name")
	business := r.Form.Get("business")
	industry := r.Form.Get("industry")

	attendee := Attendee{
		Name:     name,
		Business: business,
		ID:       a.generateID(),
		Industry: industry,
	}

	if len(a.Attendees) > 0 {
		if a.Attendees[len(a.Attendees)-1].Name == "Placeholder" {
			a.Attendees = a.Attendees[:len(a.Attendees)-1]
		}
	}

	a.Attendees = append(a.Attendees, attendee)

	err = session.Set("successMsg", fmt.Sprintf("Added %s to meeting", attendee.Name))

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

// DisplayAttendees writes the attendees directly to the browser
func (a *AppData) DisplayAttendees(w http.ResponseWriter, r *http.Request) {

	var s string
	for _, att := range a.Attendees {
		//s = s + att.Name + att.Business + "\n"
		name := att.Name
		if len(name) < 20 {
			numSpaces := 20 - len(name)
			i := 0
			for i < numSpaces {
				name = name + " "
				i++
			}
		}
		s = s + name + att.Business + "\n"
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	w.Write([]byte(s))
}

func (a *AppData) BuildChart(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	go func(ctx context.Context) {
		defer cancel()
		// ensure enough attendees have been entered
		if len(a.Attendees) < 5 {
			// start new session and create cookie
			msg := "Unable to build seating charts, not enough attendees!"
			setSessionError(msg, "/", w, r)

			return
		}

		// add a placeholder member if odd number registered
		if len(a.Attendees)%2 != 0 {
			a.Attendees = append(a.Attendees, Attendee{Name: "Placeholder"})
		}

		c := make([]Attendee, len(a.Attendees))
		num := copy(c, a.Attendees)
		fmt.Printf("number copied: %d \n", num)

		var s string

		for ok := true; ok; ok = len(a.Attendees) > 2 {
			var p Pair

			p.Seat1 = a.arrayShift()
			p.Seat2 = a.selectPartner(p.Seat1, c)

			a.Pairs = append(a.Pairs, p)
			addAttendeePairing(p.Seat1, p.Seat2, c)
		}

		// select last two no matter the match
		lastPair1 := a.arrayShift()
		lastPair2 := a.arrayShift()
		a.Pairs = append(a.Pairs, Pair{
			Seat1: lastPair1,
			Seat2: lastPair2,
		})

		addAttendeePairing(lastPair1, lastPair2, c)

		s += printPairs(a.Pairs)

		a.Pairs = []Pair{}
		// the slice should be nil at this point, reload original data
		a.Attendees = make([]Attendee, len(c))
		num = copy(a.Attendees, c)
		fmt.Printf("number copied: %d \n", num)

		a.ListCount++
		a.Pairs = []Pair{}

		// remove the placeholder attendee
		a.Attendees = arrayPop(a.Attendees)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-type", "application/json")
		w.Write([]byte(s))

	}(ctx)

	select {
	case <-ctx.Done():
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println(ctx.Err())
			setSessionError("unable to generate report, cannot find unique pairing", "/", w, r)
		}
	}

}

func printPairs(Pairs []Pair) string {
	var s string
	for _, p := range Pairs {

		p1 := fmt.Sprintf("%s (%s)", p.Seat1.Name, p.Seat1.Industry)
		if len(p1) < 60 {
			numSpaces := 60 - len(p1)
			i := 0
			for i < numSpaces {
				p1 = p1 + " "
				i++
			}
		}
		p2 := fmt.Sprintf("%s (%s)", p.Seat2.Name, p.Seat2.Industry)
		s = s + p1 + p2 + "\n"

	}
	s = s + "\n"

	return s
}

func (a *AppData) ResetData(w http.ResponseWriter, r *http.Request) {

	// start new session and create cookie
	session, _ := sessions.Start(w, r, true)
	session.Set("successMsg", "Meeting attendees have been reset")

	a.Attendees = []Attendee{}
	a.ListCount = 0

	http.Redirect(w, r, "/", http.StatusFound)
}

func setSessionError(msg, url string, w http.ResponseWriter, r *http.Request) {
	session, err := sessions.Start(w, r, true)
	if err != nil {

	}
	err = session.Set("successMsg", msg)
	http.Redirect(w, r, url, http.StatusSeeOther)

}
