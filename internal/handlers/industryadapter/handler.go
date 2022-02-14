package industryadapter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"seating/internal/app/ports"

	"github.com/gorilla/mux"
)

type HTTPHandler struct {
	industryservice ports.IndustryService
}

func NewHTTPHandler(industryService ports.IndustryService) *HTTPHandler {
	return &HTTPHandler{industryservice: industryService}
}

func (h *HTTPHandler) HandleCreateIndustry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var industry Industry

		err := json.NewDecoder(r.Body).Decode(&industry)
		if err != nil {
			fmt.Println("error unable to decode body: ", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}

		domainIndustry, err := h.industryservice.CreateIndustry(industry.Name)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}

		industry.ID = domainIndustry.ID

		b, err := json.Marshal(industry)
		if err != nil {
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(b)
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

		industry, err := h.industryservice.GetIndustry(id)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		eventResponse := convertJSONIndustryFromDomain(industry)

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

		err := h.industryservice.DeleteIndustry(id)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte("resource deleted"))
	}
}
