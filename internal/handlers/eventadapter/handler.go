package eventadapter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"seating/internal/app/ports"

	"github.com/gorilla/mux"
)

type HTTPHandler struct {
	eventService ports.EventService
}

func NewHTTPHandler(eventService ports.EventService) *HTTPHandler {
	return &HTTPHandler{eventService: eventService}
}

func (h *HTTPHandler) HandleCreateEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var event Event

		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			fmt.Println("error unable to decode body: ", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
	
			return
		}

		domainEvent, err := h.eventService.CreateEvent(event.GroupID)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}

		event.ID = domainEvent.ID
		event.Date = domainEvent.Date

		b, err := json.Marshal(event)
		if err != nil {
			return 
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(b)
	}
}

func (h *HTTPHandler) HandleGetEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if id == "" {
			// TODO: handle this much better
			w.Write([]byte("id value is empty"))
			return
		}

		event, err := h.eventService.GetEvent(id)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		eventResponse := ConvertJSONEventFromDomain(event)

		b, err := json.Marshal(eventResponse)
		if err != nil {
			return 
		}

		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func (h *HTTPHandler) HandleDeleteEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if id == "" {
			// TODO: handle this much better
			w.Write([]byte("id value is empty"))
			return
		}

		err := h.eventService.DeleteEvent(id)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte("resource deleted"))
	}
}