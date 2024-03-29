package attendeeadapter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"seating/internal/app/ports"

	"github.com/gorilla/mux"
)

type HTTPHandler struct {
	attendeeservice ports.AttendeeService
}

func NewHTTPHandler(attendeeService ports.AttendeeService) *HTTPHandler {
	return &HTTPHandler{attendeeservice: attendeeService}
}

//var demoAttendees []domain.Attendee

func (h *HTTPHandler) HandleInitDemo(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["eventId"]

		for _, attendee := range demoAttendees {
			_, err := h.attendeeservice.CreateAttendee(attendee.Name, attendee.CompanyName, attendee.Industry, id)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))

				return
			}
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("demo attendees added"))
		next.ServeHTTP(w, r)
	}
}
func (h *HTTPHandler) HandleCreateAttendee(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["eventId"]

		var attendee Attendee

		err := json.NewDecoder(r.Body).Decode(&attendee)
		if err != nil {
			fmt.Println("error unable to decode body: ", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}

		domainAttendee, err := h.attendeeservice.CreateAttendee(attendee.Name, attendee.CompanyName, attendee.Industry, id)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}

		attendee.ID = domainAttendee.ID // TODO: kinda questioning this?

		b, err := json.Marshal(attendee)
		if err != nil {
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(b)

		next.ServeHTTP(w, r)
	}
}

func (h *HTTPHandler) HandleGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if id == "" {
			// TODO: handle this much better
			w.Write([]byte("id value is empty"))
			return
		}

		attendee, err := h.attendeeservice.GetAttendee(id)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		eventResponse := convertJSONAttendeeFromDomain(attendee)

		b, err := json.Marshal(eventResponse)
		if err != nil {
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func (h *HTTPHandler) HandleDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if id == "" {
			// TODO: handle this much better
			w.Write([]byte("id value is empty"))
			return
		}

		err := h.attendeeservice.DeleteAttendee(id)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte("resource deleted"))
	}
}
