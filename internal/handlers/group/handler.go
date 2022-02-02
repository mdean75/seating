package groupadapter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"seating/internal/app/ports"
)

type HTTPHandler struct {
	groupService ports.GroupService
}

func NewHTTPHandler(groupService ports.GroupService) *HTTPHandler {
	return &HTTPHandler{groupService: groupService}
}

func (h *HTTPHandler) HandleCreateGroup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		
		var group Group
		err := json.NewDecoder(r.Body).Decode(&group)
		if err != nil {
			fmt.Println("error unable to decode body: ", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		
			return
		}

		id, err := h.groupService.CreateGroup(group.DisplayName, group.ShortName)

		// id, err := c.Datastore.CreateGroup(group.DisplayName, group.ShortName)
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