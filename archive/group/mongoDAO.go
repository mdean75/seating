package group

// import (
// 	"context"
// 	"seating/internal/app/ports"
// 	"seating/internal/db"
// 	groupadapter "seating/internal/handlers/group"

// 	// "seating/internal/app/group"

// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// type MongoDataStore struct {
// 	*db.MongoConn
// 	db  *mongo.Database
// 	col *mongo.Collection
// }

// func NewDAO(dbconn *db.MongoConn, db, col string) ports.GroupRepository {
// 	dbx := dbconn.Client.Database(db)
// 	conx := dbx.Collection(col)

// 	return &MongoDataStore{dbconn, dbx, conx}
// }

// func NewGroup(displayName, shortName string) domain.Group {
// 	return domain.Group{
// 		DisplayName: displayName,
// 		ShortName: shortName,
// 	}
// }

// func NewEvent(groupID ports.ID) domain.Event {
// 	return domain.Event{
// 		GroupID: string(groupID),
// 		Date: time.Now(),
// 	}
// }

// func (m *MongoDataStore) CreateGroup(displayName, shortName string) (ports.ID, error) {
// 	group := groupadapter.NewGroupRequest(displayName, shortName)

// 	res, err := m.db.Collection(m.col.Name()).InsertOne(context.TODO(), group)
// 	if err != nil {
// 		return "", err
// 	}

// 	objId := res.InsertedID.(primitive.ObjectID)

// 	return ports.ID(objId.Hex()), nil
// }

// func (m *MongoDataStore) CreateEvent(groupId ports.ID) (ports.ID, error) {
// 	event := eventadapter.NewEventRequest(eventadapter.ID(groupId))

// 	res, err := m.col.InsertOne(context.TODO(), event)
// 	if err != nil {
// 		return "", err
// 	}

// 	objID := res.InsertedID.(primitive.ObjectID)

// 	return ports.ID(objID.Hex()), nil
// }