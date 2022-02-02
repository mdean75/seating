package eventrepo

import (
	"context"
	"seating/internal/app/ports"
	"seating/internal/db"
	eventadapter "seating/internal/handlers/event"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDataStore struct {
	*db.MongoConn
	db  *mongo.Database
	col *mongo.Collection
}

func NewDAO(dbconn *db.MongoConn, db, col string) ports.EventRepository {
	dbx := dbconn.Client.Database(db)
	conx := dbx.Collection(col)

	return &MongoDataStore{dbconn, dbx, conx}
}

func (m *MongoDataStore) CreateEvent(groupId ports.ID) (ports.ID, error) {
	event := eventadapter.NewEventRequest(eventadapter.ID(groupId))

	res, err := m.col.InsertOne(context.TODO(), event)
	if err != nil {
		return "", err
	}

	objID := res.InsertedID.(primitive.ObjectID)

	return ports.ID(objID.Hex()), nil
}