package group

import (
// "encoding/json"
// "fmt"
// "net/http"
// "seating/internal/app/domain"
// "seating/internal/app/ports"
// "seating/internal/app"
// "seating/internal/app/group"
)

// type Group struct {
// 	ID string `bson:"_id,omitempty"`
// 	DisplayName string `bson:"displayName"`
// 	ShortName string `bson:"shortName"`
// }

// func HandleCreateGroup(c *Controller) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		var group domain.Group
// 		err := json.NewDecoder(r.Body).Decode(&group)
// 		if err != nil {
// 			fmt.Println("error unable to decode body: ", err)
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write([]byte(err.Error()))

// 			return
// 		}

// 		id, err := c.Datastore.CreateGroup(group.DisplayName, group.ShortName)
// 		if err != nil {
// 			fmt.Println(err)
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write([]byte(err.Error()))

// 			return
// 		}

// 		w.WriteHeader(http.StatusCreated)
// 		w.Write([]byte(id))

// 	}
// }

// func HandleCreateEvent(c *Controller) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var event domain.Event

// 		err := json.NewDecoder(r.Body).Decode(&event)
// 		if err != nil {
// 			fmt.Println("error unable to decode body: ", err)
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write([]byte(err.Error()))

// 			return
// 		}

// 		id, err := c.Datastore.CreateEvent(ports.ID(event.GroupID))

// 		if err != nil {
// 			fmt.Println(err)
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write([]byte(err.Error()))

// 			return
// 		}

// 		w.WriteHeader(http.StatusCreated)
// 		w.Write([]byte(id))
// 	}
// }
