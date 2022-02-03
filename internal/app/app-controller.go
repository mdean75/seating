package app

// import "seating/internal/app/group"

// // "fmt"
// // "seating/internal/app/group"
// // "seating/internal/config"
// // "seating/internal/db"

// type ID string

// type Repository interface {
	
// 	CreateGroup(displayName, shortName string) (ID, error)
// 	// StartEvent(group interface{}) (ID, error) // interface will be replaced with group type
	
// 	// AddAttendee(eventID ID, name, businessName, industry string) error

// 	// GetAttendees(eventID ID) ([]interface{}, error) // interface{} will be replaced with attenddee type
	
// 	// CreatePairing(eventID ID) error

// 	// GetPairing(eventID ID, pairingNum int) (interface{}, error) // interface{} will be replaced with pairing type
// 	// GetAllMeetingPairings(eventID ID) ([]interface{}, error) // interface{} will be replaced with pairing type

// 	// GetGroupsEvents(groupID ID) ([]interface{}, error) // interface{} will be replaced with event type - also just meta

// }

// type Controller struct {
// 	// Datastore         Repository
// 	GroupController *group.Controller
	
	
// }

// func NewController(groupController *group.Controller) *Controller {
// 	return &Controller{GroupController: groupController}
// }

// // func NewController(r Repository) *Controller {
// // 	return &Controller{Datastore: r}
// // }

// // func CreateController() (*Controller, error) {
// // 	conf := config.EnvVar{}.LoadConfig()
// // 	// if conf.DBConn() == "" {
// // 	// 	return nil, fmt.Errorf("dbconn is not set")
// // 	// }
	

// // 	conf.MongoConfig.SetDBConn("mongodb://127.0.0.1:27017")
// // 	mongoConn, err := db.NewMongoDatabase(conf.DBConn())
// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	dao := NewDAO(mongoConn, "testdb", "usergroup")

// // 	return NewController(dao), nil
// // }