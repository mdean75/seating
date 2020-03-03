package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"seating/api"
	"time"
)

func main() {

	var a api.AppData
	a.Industries = api.SetIndustries()
	r := mux.NewRouter()

	// original handlers for the app
	r.HandleFunc("/", a.AttendeeEntry).Methods(http.MethodGet)             // display the form
	r.HandleFunc("/", a.ProcessAttendeeEntry).Methods(http.MethodPost)     // process form input
	r.HandleFunc("/attendees", a.DisplayAttendees).Methods(http.MethodGet) // display list of attendees
	r.HandleFunc("/seating", a.BuildChart).Methods(http.MethodGet)         // build and display the seating chart
	r.HandleFunc("/reset-attendees", a.ResetData).Methods(http.MethodGet)  // clear all data
	r.HandleFunc("/demo", a.Demo).Methods(http.MethodGet)                  // load the demo attendee data

	// api handlers
	r.HandleFunc("/api/reset", a.ResetAttendeesAPI).Methods(http.MethodGet)
	r.Handle("/api/attendee", a.AddAttendeeAPI()).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/attendees", a.DisplayAttendeesAPI).Methods(http.MethodGet)
	r.HandleFunc("/api/seating", a.BuildChartAPI).Methods(http.MethodGet)
	r.Handle("/api/appdata", a.GetAppData()).Methods(http.MethodGet)
	r.Handle("/api/count", a.GetListCount()).Methods(http.MethodGet)
	r.Handle("/api/industry", a.GetIndustries()).Methods(http.MethodGet)
	r.HandleFunc("/api/demo", a.DemoAPI).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:         "0.0.0.0:3000",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Println("api listening on port " + srv.Addr)
	log.Fatal(srv.ListenAndServe())

}
