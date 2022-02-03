package eventrepo

import (
	"context"
	"fmt"
	"seating/internal/app/domain"
	"seating/internal/app/ports"
	"seating/internal/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

func convertMongoEventToDomain(event Event) domain.Event {
	return domain.NewEvent(event.ID, event.GroupID, event.Date)
}

func (m *MongoDataStore) Save(event domain.Event) (ports.ID, error) {
	e := NewMongoEventFromDomain(event)

	res, err := m.col.InsertOne(context.TODO(), e)
	if err != nil {
		return "", err
	}

	objID := res.InsertedID.(primitive.ObjectID)

	return ports.ID(objID.Hex()), nil
}

func (m *MongoDataStore) Get(eventID string) (domain.Event, error) {
	var event Event
	id, err := primitive.ObjectIDFromHex(eventID)
	if err != nil {
		return domain.Event{}, err
	}

	err = m.col.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&event)
	if err != nil {
		return domain.Event{}, err
	}

	return convertMongoEventToDomain(event), nil
	
}

var ErrUnableToDeleteResouce error = fmt.Errorf("unable to delete the specified resource")

func (m *MongoDataStore) Delete(eventID string) error {
	id, err := primitive.ObjectIDFromHex(eventID)
	if err != nil {
		return err
	}

	result, err := m.col.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return ErrUnableToDeleteResouce
	}

	return nil
}