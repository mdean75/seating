package eventadapter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"seating/internal/app/domain"
	"seating/internal/app/ports"
	"seating/internal/handlers/attendeeadapter"
	"seating/internal/handlers/groupadapter"

	"github.com/gorilla/mux"
)

type HTTPHandler struct {
	eventService ports.EventService
}

func NewHTTPHandler(eventService ports.EventService) *HTTPHandler {
	return &HTTPHandler{eventService: eventService}
}

func (h *HTTPHandler) HandleGetPairingCount(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id := vars["id"]

		count, err := h.eventService.GetListCount(id)
		if err != nil {
			fmt.Println("error unable to get list count from db: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))

			next.ServeHTTP(w, r)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := make(map[string]interface{})
		resp["listCount"] = count

		json.NewEncoder(w).Encode(resp)

		next.ServeHTTP(w, r)
	}
}

func (h *HTTPHandler) HandleCreatingPairingRound(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var attendees []attendeeadapter.Attendee

		err := json.NewDecoder(r.Body).Decode(&attendees)
		if err != nil {
			fmt.Println("error unable to decode body: ", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}

		// convert the json attendee list to a domain array
		var domainAttendees []domain.Attendee
		for _, a := range attendees {
			domainAttendees = append(domainAttendees, domain.NewAttendee(a.ID, a.Name, a.CompanyName, a.Industry))
		}

		// convert the json attendde list to a domain attendee map
		domAttendees := make(map[string]domain.Attendee)
		for _, a := range attendees {
			domAttendees[a.ID] = domain.NewAttendee(a.ID, a.Name, a.CompanyName, a.Industry)
		}

		// get the event data which has all the attendees and their previous pairs
		event, err := h.eventService.GetEvent(id)
		if err != nil {
			fmt.Println("error unable to decode body: ", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}

		// convert event attendees to map
		dbEventAttendees := make(map[string]domain.Attendee)
		for _, d := range event.Attendees {
			dbEventAttendees[d.ID] = d
		}

		// add the 'paired with' data from db to attendees
		for _, d := range domAttendees {
			temp := domAttendees[d.ID]
			temp.PairedWith = dbEventAttendees[d.ID].PairedWith
			domAttendees[d.ID] = temp
		}

		// returns a domain attendee list
		pairing, err := domain.NewPairingRound(domAttendees)
		if err != nil {
			fmt.Println("error unable to decode body: ", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}

		var response []Pair

		for _, a := range pairing.Pairs {
			response = append(response, Pair{
				Seat1: attendeeadapter.Attendee{
					ID:          a.Seat1.ID,
					Name:        a.Seat1.Name,
					CompanyName: a.Seat1.CompanyName,
					Industry:    a.Seat1.Industry,
				},
				Seat2: attendeeadapter.Attendee{
					ID:          a.Seat2.ID,
					Name:        a.Seat2.Name,
					CompanyName: a.Seat2.CompanyName,
					Industry:    a.Seat2.Industry,
				}})
		}

		// convert response to domain object
		var domainPairs []domain.Pair
		for _, p := range response {
			domainPairs = append(domainPairs, domain.NewPair(
				domain.NewAttendee(p.Seat1.ID, p.Seat1.Name, p.Seat1.CompanyName, p.Seat1.Industry),
				domain.NewAttendee(p.Seat2.ID, p.Seat2.Name, p.Seat2.CompanyName, p.Seat2.Industry)))
		}

		err = h.eventService.CreatePairingRound(id, domainPairs)
		if err != nil {
			fmt.Println(err)
		}

		b, err := json.Marshal(response)
		if err != nil {
			fmt.Println("error unable to decode body: ", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)

		next.ServeHTTP(w, r)
	}
}

func (h *HTTPHandler) HandleCreateEvent(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var event Event

		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			fmt.Println("error unable to decode body: ", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			//return
			next.ServeHTTP(w, r)
			return
		}

		domainEvent, err := h.eventService.CreateEvent(event.GroupID, groupadapter.ConvertDomainGroupFromJSON(event.Group), event.Date)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			//return
			next.ServeHTTP(w, r)
			return
		}

		event.ID = domainEvent.ID
		event.Date = domainEvent.Date

		b, err := json.Marshal(event)
		if err != nil {
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(b)

		next.ServeHTTP(w, r)
	}
}

func (h *HTTPHandler) HandleGetEvent(next http.Handler) http.HandlerFunc {
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

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)

		next.ServeHTTP(w, r)
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

func (h *HTTPHandler) HandleGetGroupsEvents(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		w.Header().Set("Content-Type", "application/json")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			resp := map[string]string{"error": "Cannot get events for group, missing group id"}
			json.NewEncoder(w).Encode(resp)

			next.ServeHTTP(w, r)
			return
		}

		events, err := h.eventService.GetEventsForGroup(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp := map[string]string{"errorMsg": "Cannot get events for group", "err": err.Error()}
			json.NewEncoder(w).Encode(resp)

			next.ServeHTTP(w, r)
			return
		}

		var jsonEvents []Event
		for _, domEvent := range events {
			jsonEvents = append(jsonEvents, Event{
				ID:   domEvent.ID,
				Date: domEvent.Date,
			})
		}

		w.WriteHeader(http.StatusOK)
		resp := map[string]interface{}{"events": jsonEvents}

		json.NewEncoder(w).Encode(resp)

		next.ServeHTTP(w, r)
	}
}
