/**
filename:	seating-api.go
purpose:	provides functions specific to the api
update:		3/3/2020
comments:	all of the functions that are specific to the api only functionality.
*/

package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// AddAttendeeAPI handles the request to add a new meeting attendee.
func (a *AppData) AddAttendeeAPI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// todo: convert to post method
		name := r.URL.Query().Get("name")
		bus := r.URL.Query().Get("business")
		ind := r.URL.Query().Get("industry")

		var att Attendee
		//var i interface{}
		//err := json.NewDecoder(r.Body).Decode(&i)
		//if err != nil {
		//	resp := map[string]string{"error": err.Error()}
		//
		//	//resp := map[string]string{"response": "successfully added attendee"}
		//	b, _ := json.Marshal(resp)
		//
		//	w.Header().Set("Content-Type", "application/json")
		//	w.Header().Set("Access-Control-Allow-Origin", "*")
		//	w.Header().Set("Access-Control-Allow-Origin", "Origin, Content-Type, X-Auth-Token")
		//	w.Header().Set("Access-Control-Allow-Methods", "PUT, GET, POST, DELETE, OPTIONS");
		//	w.Header().Set("Cache-Control", "no-store")
		//	w.WriteHeader(http.StatusBadRequest)
		//	w.Write(b)
		//
		//	return
		//}

		att.Name = name
		att.Business = bus
		att.Industry = ind
		att.ID = a.generateID() //randomInt(1, 1000)
		a.Attendees = append(a.Attendees, att)

		resp := map[string]string{"response": "successfully added attendee"}
		b, _ := json.Marshal(resp)

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Cache-Control", "no-store")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}

}

// DisplayAttendeesAPI handles the request to fetch the list of meeting attendees.
func (a *AppData) DisplayAttendeesAPI(w http.ResponseWriter, r *http.Request) {

	b, _ := json.Marshal(a.Attendees)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// ResetAttendeesAPI handles the request to reset the meeting attendees and the pairing list data.
func (a *AppData) ResetAttendeesAPI(w http.ResponseWriter, r *http.Request) {

	a.Attendees = []Attendee{}
	a.ListCount = 0

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(http.StatusOK)
	w.Write(nil)
}

// DisplayPairsAPI handles the request to get the current meeting pairings.
// todo: not sure if this is used in the current implementation
func (a *AppData) DisplayPairsAPI(w http.ResponseWriter, r *http.Request) {

	m := struct {
		Pairs []Pair
	}{Pairs: a.Pairs}

	b, _ := json.Marshal(m)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// GetAppData handles the request to get the full appData details
func (a *AppData) GetAppData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Cache-Control", "no-store")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(a)
	}
}

// GetListCount handles the request to get only the number of pairing lists that have been generated.
func (a *AppData) GetListCount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Cache-Control", "no-store")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(a.ListCount)
	}
}

// GetIndustries handles the request to get the list of industries to display in the option select.
func (a *AppData) GetIndustries() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		m := struct {
			Industries []string
		}{Industries: a.Industries}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(m)
	}
}

// BuildChartAPI handles the request to build a pairing list.  The logic is handled in a go routine with a context
//timeout to prevent an infinite loop in the case that a unique pairing cannot be generated.
func (a *AppData) BuildChartAPI(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	go func(ctx context.Context) {
		defer cancel()
		// ensure enough attendees have been entered
		if len(a.Attendees) < 5 {
			m := map[string]string{"error": "Unable to build seating charts, not enough attendees!"}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(m)

			return
		}

		// add a placeholder member if odd number registered
		if len(a.Attendees)%2 != 0 {
			a.Attendees = append(a.Attendees, Attendee{Name: "Placeholder"})
		}

		c := make([]Attendee, len(a.Attendees))
		num := copy(c, a.Attendees)
		fmt.Printf("number copied: %d \n", num)

		//var s string

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

		//m := struct {
		//	Pairs []Pair
		//}{Pairs: a.Pairs}

		b, err := json.Marshal(a.Pairs)
		if err != nil {
			fmt.Println(err)
		}

		a.Pairs = []Pair{}
		// the slice should be nil at this point, reload original data
		a.Attendees = make([]Attendee, len(c))
		num = copy(a.Attendees, c)
		fmt.Printf("number copied: %d \n", num)

		a.ListCount++
		a.Pairs = []Pair{}

		// remove the placeholder attendee
		a.Attendees = arrayPop(a.Attendees)

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Cache-Control", "no-store")
		w.WriteHeader(http.StatusOK)
		w.Write(b)

		_, err = a.Conn.Database("testdb").Collection("testcol").InsertOne(context.TODO(), map[string]interface{}{"attendees": a.Attendees, "pairs": a.Pairs})
		if err != nil {
			fmt.Println("error: ", err)
		}

	}(ctx)

	select {
	case <-ctx.Done():
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println(ctx.Err())
			fmt.Println("unable to generate report, cannot find unique pairing")
		}
	}

}
