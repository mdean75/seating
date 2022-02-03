package groupadapter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"seating/internal/app/ports"

	"github.com/gorilla/mux"
)

type HTTPHandler struct {
	groupService ports.GroupService
}

func NewHTTPHandler(groupService ports.GroupService) *HTTPHandler {
	return &HTTPHandler{groupService: groupService}
}

// handler will decode body into json group, 
// call group service which creates a domain group
// then call group repository to convert to dao object and persist
// which returns an id and error to the group service
// that id is added to the domain group object in the service
// that domain group is then returned to this function
// the domain group id is added to the json group object and response 
// is sent back to the called with the completed group
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

		domainGroup, err := h.groupService.CreateGroup(group.DisplayName, group.ShortName)

		// id, err := c.Datastore.CreateGroup(group.DisplayName, group.ShortName)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

			return
		}

		group.ID = domainGroup.ID
		b, err := json.Marshal(group)
		if err != nil {
			return // maybe write just the id
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(b)
		
	}
}

func (h *HTTPHandler) HandleGetGroup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if id == "" {
			// TODO: handle this much better
			w.Write([]byte("id value is empty"))
			return
		}

		group, err := h.groupService.GetGroup(id)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		groupResponse := ConvertJSONGroupFromDomain(group)

		b, err := json.Marshal(groupResponse)
		if err != nil {
			return 
		}

		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}