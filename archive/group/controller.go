package group

// import (
// 	"seating/internal/app/ports"
// 	"seating/internal/config"
// 	"seating/internal/db"
// )

// // type ID string

// // type Repository interface {
// // 	CreateGroup(displayName, shortName string) (ID, error)
// // 	CreateEvent(groupID ID, ) (ID, error)
// // }

// type Controller struct {
// 	Datastore         ports.Repository
// }

// func NewController(r ports.Repository) *Controller {
// 	return &Controller{Datastore: r}
// }

// func CreateController() (*Controller, error) {
// 	conf := config.EnvVar{}.LoadConfig()
// 	// if conf.DBConn() == "" {
// 	// 	return nil, fmt.Errorf("dbconn is not set")
// 	// }
	

// 	conf.MongoConfig.SetDBConn("mongodb://127.0.0.1:27017")
// 	mongoConn, err := db.NewMongoDatabase(conf.DBConn())
// 	if err != nil {
// 		return nil, err
// 	}

// 	dao := NewDAO(mongoConn, "testdb", "usergroup")

// 	return NewController(dao), nil
// }