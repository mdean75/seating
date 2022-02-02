package eventadapter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"seating/internal/app/ports"
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

		// id, err := c.Datastore.CreateEvent(ports.ID(event.GroupID))
		id, err := h.eventService.CreateEvent(event.GroupID)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(id))
	}
}