package eventrepo

import (
	"context"
	"seating/internal/app/domain"
	"seating/internal/app/ports"
	"seating/internal/db"
	"time"

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

type Event struct {
	ID string `bson:"_id,omitempty"`
	Date time.Time `bson:"date"`
	GroupID string `bson:"groupId"`
}

func NewMongoEventFromDomain(domainEvent domain.Event) Event {
	return Event{
		ID: domainEvent.ID,
		Date: domainEvent.Date,
		GroupID: domainEvent.GroupID,
	}
}

func (m *MongoDataStore) Save(event domain.Event) (ports.ID, error) {
	// event := eventadapter.NewEventRequest(eventadapter.ID(groupId))
	e := NewMongoEventFromDomain(event)

	res, err := m.col.InsertOne(context.TODO(), e)
	if err != nil {
		return "", err
	}

	objID := res.InsertedID.(primitive.ObjectID)

	return ports.ID(objID.Hex()), nil
}