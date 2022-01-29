package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"seating/api"
	"seating/internal/repository"
)

func NewRouterWithCors() http.Handler {
	methods := []string{http.MethodPost, http.MethodGet, http.MethodOptions}
	origins := []string{"http://localhost:4200", "https://letters2lostlovedones.com", "http://127.0.0.1:4200"}
	headers := []string{"Content-Type"}

	opts := cors.Options{
		AllowedMethods:     methods,
		AllowedOrigins:     origins,
		AllowedHeaders:     headers,
		OptionsPassthrough: true,
		Debug:              true,
	}

	r := mux.NewRouter()

	addRoutes(r)
	

	crs := cors.New(opts)
	return crs.Handler(r)
}

func addRoutes(r *mux.Router) {
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	m, err := repository.NewMongoDatabase("mongodb://127.0.0.1:27017")
	if err != nil {
		fmt.Println(err)
	}
	
	var a api.AppData
	a.Industries = api.SetIndustries()
	a.Conn = m

	r.HandleFunc("/api/reset", a.ResetAttendeesAPI).Methods(http.MethodGet)
	r.Handle("/api/attendee", a.AddAttendeeAPI()).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/attendees", a.DisplayAttendeesAPI).Methods(http.MethodGet)
	r.HandleFunc("/api/seating", a.BuildChartAPI).Methods(http.MethodGet)
	r.Handle("/api/appdata", a.GetAppData()).Methods(http.MethodGet)
	r.Handle("/api/count", a.GetListCount()).Methods(http.MethodGet)
	r.Handle("/api/industry", a.GetIndustries()).Methods(http.MethodGet)
	r.HandleFunc("/api/demo", a.DemoAPI).Methods(http.MethodGet)
}

