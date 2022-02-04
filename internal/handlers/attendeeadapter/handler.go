package attendeeadapter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"seating/internal/app/ports"
)

type HTTPHandler struct {
	attendeeservice ports.AttendeeService
}

func NewHTTPHandler(attendeeService ports.AttendeeService) *HTTPHandler {
	return &HTTPHandler{attendeeservice: attendeeService}
}

func (h *HTTPHandler) HandleCreateAttendee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var attendee Attendee

		err := json.NewDecoder(r.Body).Decode(&attendee)
		if err != nil {
			fmt.Println("error unable to decode body: ", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
	
			return
		}

		domainAttendee, err := h.attendeeservice.CreateAttendee(attendee.Name, attendee.CompanyName, attendee.Industry)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}

		attendee.ID = domainAttendee.ID

		b, err := json.Marshal(attendee)
		if err != nil {
			return 
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(b)
	}
}