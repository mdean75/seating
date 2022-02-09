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
	r.HandleFunc("/api/demo", a.DemoAPI).Methods(http.MethodGet)

	r.HandleFunc("/group", controller.GroupHandler.HandleCreateGroup()).Methods(http.MethodPost)
	r.HandleFunc("/group/{id}", controller.GroupHandler.HandleGetGroup()).Methods(http.MethodGet)
	r.HandleFunc("/group", controller.GroupHandler.HandleGetAllGroups()).Methods(http.MethodGet)
	r.HandleFunc("/group/{id}", controller.GroupHandler.HandleDeleteGroup()).Methods(http.MethodDelete)

	r.HandleFunc("/event", HandlePreFlight(
		controller.EventHandler.HandleCreateEvent(
			logRequest()))).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/event/{id}", controller.EventHandler.HandleGetEvent()).Methods(http.MethodGet)
	r.HandleFunc("/event/{id}", controller.EventHandler.HandleDeleteEvent()).Methods(http.MethodDelete)
	r.HandleFunc("/event/{id}/pairing", controller.EventHandler.HandleCreatingPairingRound()).Methods(http.MethodPost)

	r.HandleFunc("/event/{eventId}/attendee", HandlePreFlight(
		controller.AttendeeHandler.HandleCreateAttendee(
			logRequest()))).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/event/{eventId}/demo", controller.AttendeeHandler.HandleInitDemo()).Methods(http.MethodPost)
	r.HandleFunc("/attendee/{id}", controller.AttendeeHandler.HandleGet()).Methods(http.MethodGet)
	r.HandleFunc("/attendee/{id}", controller.AttendeeHandler.HandleDelete()).Methods(http.MethodDelete)

	r.HandleFunc("/industry", controller.Industryhandler.HandleCreateIndustry()).Methods(http.MethodPost)
	r.HandleFunc("/industry/{id}", controller.Industryhandler.HandleGet()).Methods(http.MethodGet)
	r.HandleFunc("/industry/{id}", controller.Industryhandler.HandleDelete()).Methods(http.MethodDelete)
}
