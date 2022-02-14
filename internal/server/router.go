package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"seating/api"
	"seating/internal/app"
	"seating/internal/config"
)

func NewRouterWithCors(controller *app.Controller, conf config.Configuration) http.Handler {
	methods := []string{http.MethodPost, http.MethodGet, http.MethodOptions}
	origins := []string{"http://localhost:8080", "http://192.168.1.211:8080", "http://localhost:4200", "https://seating-ui.bedaring.me", "http://127.0.0.1:4200"}
	headers := []string{"Content-Type"}

	opts := cors.Options{
		AllowedMethods:     methods,
		AllowedOrigins:     origins,
		AllowedHeaders:     headers,
		OptionsPassthrough: true,
		Debug:              conf.DebugCors,
	}

	r := mux.NewRouter()

	addRoutes(r, controller)

	crs := cors.New(opts)
	return crs.Handler(r)
}

func addRoutes(r *mux.Router, controller *app.Controller) {
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
	r.HandleFunc("/api/demo", a.DemoAPI).Methods(http.MethodGet, http.MethodOptions)

	// group routes
	r.HandleFunc("/group", controller.GroupHandler.HandleCreateGroup()).Methods(http.MethodPost)
	r.HandleFunc("/group/{id}", controller.GroupHandler.HandleGetGroup()).Methods(http.MethodGet)
	r.HandleFunc("/group", controller.GroupHandler.HandleGetAllGroups()).Methods(http.MethodGet)
	r.HandleFunc("/group/{id}", controller.GroupHandler.HandleDeleteGroup()).Methods(http.MethodDelete)

	// get events for group by group id
	r.HandleFunc("/group/{id}/event", HandlePreFlight(
		controller.EventHandler.HandleGetGroupsEvents(
			logRequest()))).Methods(http.MethodGet, http.MethodOptions)

	// event routes
	// create event
	r.HandleFunc("/event", HandlePreFlight(
		controller.EventHandler.HandleCreateEvent(
			logRequest()))).Methods(http.MethodPost, http.MethodOptions)

	// get event by id
	r.HandleFunc("/event/{id}", HandlePreFlight(
		controller.EventHandler.HandleGetEvent(
			logRequest()))).Methods(http.MethodGet, http.MethodOptions)

	// delete event by id
	r.HandleFunc("/event/{id}", controller.EventHandler.HandleDeleteEvent()).Methods(http.MethodDelete)

	// create pairing round for event
	r.HandleFunc("/event/{id}/pairing", HandlePreFlight(
		controller.EventHandler.HandleCreatingPairingRound(
			logRequest()))).Methods(http.MethodPost, http.MethodOptions)

	// get count of pairing rounds for event
	r.HandleFunc("/event/{id}/pairing/count", HandlePreFlight(
		controller.EventHandler.HandleGetPairingCount(
			logRequest()))).Methods(http.MethodGet, http.MethodOptions)

	// create / add attendee for event
	r.HandleFunc("/event/{eventId}/attendee", HandlePreFlight(
		controller.AttendeeHandler.HandleCreateAttendee(
			logRequest()))).Methods(http.MethodPost, http.MethodOptions)

	// add demo attendees to an event
	r.HandleFunc("/event/{eventId}/demo", HandlePreFlight(
		controller.AttendeeHandler.HandleInitDemo(
			logRequest()))).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/attendee/{id}", controller.AttendeeHandler.HandleGet()).Methods(http.MethodGet)
	r.HandleFunc("/attendee/{id}", controller.AttendeeHandler.HandleDelete()).Methods(http.MethodDelete)

	r.HandleFunc("/industry", controller.Industryhandler.HandleCreateIndustry()).Methods(http.MethodPost)
	r.HandleFunc("/industry/{id}", controller.Industryhandler.HandleGet()).Methods(http.MethodGet)
	r.HandleFunc("/industry/{id}", controller.Industryhandler.HandleDelete()).Methods(http.MethodDelete)
}
