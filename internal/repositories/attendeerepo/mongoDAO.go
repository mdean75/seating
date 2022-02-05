package attendeerepo

import (
	"context"
	"fmt"
	"seating/internal/app/domain"
	"seating/internal/app/ports"
	"seating/internal/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDataStore struct {
	*db.MongoConn
	db  *mongo.Database
	col *mongo.Collection
}

func NewDAO(dbconn *db.MongoConn, db, col string) ports.AttendeeRepository {
	dbx := dbconn.Client.Database(db)
	conx := dbx.Collection(col)

	return &MongoDataStore{dbconn, dbx, conx}
}

type Attendee struct {
	ID          string `bson:"_id,omitempty"`
	Name        string `bson:"name"`
	CompanyName string `bson:"companyName"`
	Industry    string `bson:"industry"`
}

func NewAttendee(id, name, company, industry string) Attendee {
	return Attendee{
		ID:          id,
		Name:        name,
		CompanyName: company,
		Industry:    industry,
	}
}

func convertDomainAttendeeToMongo(attendee domain.Attendee) Attendee {
	return NewAttendee(attendee.ID, attendee.Name, attendee.CompanyName, attendee.Industry)
}

func convertMongoAttendeeToDomain(attendee Attendee) domain.Attendee {
	return domain.NewAttendee(attendee.ID, attendee.Name, attendee.CompanyName, attendee.Industry)
}

func (m *MongoDataStore) Save(attendee domain.Attendee, eventID string) (ports.ID, error) {
	a := convertDomainAttendeeToMongo(attendee)
	objID, err := primitive.ObjectIDFromHex(eventID)
	if err != nil {
		return "", err
	}

	//res, err := m.col.InsertOne(context.TODO(), a)
	result, err := m.col.UpdateByID(context.TODO(), objID, bson.M{"$addToSet": bson.M{"attendees": a}})
	if err != nil {
		return "", err
	}

	if result.ModifiedCount == 0 {
		return "", fmt.Errorf("record was not updated")
	}

	//objID := res.InsertedID.(primitive.ObjectID)

	return ports.ID(objID.Hex()), nil
}

func (m *MongoDataStore) Get(attendeeID string) (domain.Attendee, error) {
	var attendee Attendee
	id, err := primitive.ObjectIDFromHex(attendeeID)
	if err != nil {
		return domain.Attendee{}, err
	}

	err = m.col.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&attendee)
	if err != nil {
		return domain.Attendee{}, err
	}

	return convertMongoAttendeeToDomain(attendee), nil
}

var ErrUnableToDeleteResouce error = fmt.Errorf("unable to delete the specified resource")

func (m *MongoDataStore) Delete(attendeeID string) error {
	id, err := primitive.ObjectIDFromHex(attendeeID)
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
