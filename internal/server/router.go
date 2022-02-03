package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"seating/api"
	"seating/internal/config"
	"seating/internal/handlers/eventadapter"
	"seating/internal/handlers/groupadapter"
)

func NewRouterWithCors(groupService *groupadapter.HTTPHandler, eventService *eventadapter.HTTPHandler, conf config.Configuration) http.Handler {
	methods := []string{http.MethodPost, http.MethodGet, http.MethodOptions}
	origins := []string{"http://localhost:4200", "https://letters2lostlovedones.com", "http://127.0.0.1:4200"}
	headers := []string{"Content-Type"}

	opts := cors.Options{
		AllowedMethods:     methods,
		AllowedOrigins:     origins,
		AllowedHeaders:     headers,
		OptionsPassthrough: true,
		Debug:              conf.DebugCors,
	}

	r := mux.NewRouter()

	addRoutes(r, groupService, eventService)
	

	crs := cors.New(opts)
	return crs.Handler(r)
}

func addRoutes(r *mux.Router, groupService *groupadapter.HTTPHandler, eventService *eventadapter.HTTPHandler) {
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	
	var a api.AppData
	a.Industries = api.SetIndustries()

	r.HandleFunc("/api/reset", a.ResetAttendeesAPI).Methods(http.MethodGet)
	r.Handle("/api/attendee", a.AddAttendeeAPI()).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/api/attendees", a.DisplayAttendeesAPI).Methods(http.MethodGet)
	r.HandleFunc("/api/seating", a.BuildChartAPI).Methods(http.MethodGet)
	r.Handle("/api/appdata", a.GetAppData()).Methods(http.MethodGet)
	r.Handle("/api/count", a.GetListCount()).Methods(http.MethodGet)
	r.Handle("/api/industry", a.GetIndustries()).Methods(http.MethodGet)
	r.HandleFunc("/api/demo", a.DemoAPI).Methods(http.MethodGet)

	r.HandleFunc("/group", groupService.HandleCreateGroup()).Methods(http.MethodPost)
	r.HandleFunc("/group/{id}", groupService.HandleGetGroup()).Methods(http.MethodGet)

	r.HandleFunc("/event", eventService.HandleCreateEvent()).Methods(http.MethodPost)
	r.HandleFunc("/event/{id}", eventService.HandleGetEvent()).Methods(http.MethodGet)
	r.HandleFunc("/event/{id}", eventService.HandleDeleteEvent()).Methods(http.MethodDelete)
}

