package attendeerepo

import (
	"context"
	"seating/internal/app/domain"
	"seating/internal/app/ports"
	"seating/internal/db"

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
	ID string `bson:"_id,omitempty"`
	Name string `bson:"name"`
	CompanyName string `bson:"companyName"`
	Industry string `bson:"industry"`
}

func NewAttendee(id, name, company, industry string) Attendee {
	return Attendee{
		ID: id,
		Name: name,
		CompanyName: company,
		Industry: industry,
	}
}

func convertDomainAttendeeToMongo(attendee domain.Attendee) Attendee {
	return NewAttendee(attendee.ID, attendee.Name, attendee.CompanyName, attendee.Industry)
}

func convertMongoAttendeeToDomain(attendee Attendee) domain.Attendee {
	return domain.NewAttendee(attendee.ID, attendee.Name, attendee.CompanyName, attendee.Industry)
}

func (m *MongoDataStore) Save(attendee domain.Attendee) (ports.ID, error) {
	a := convertDomainAttendeeToMongo(attendee)

	res, err := m.col.InsertOne(context.TODO(), a)
	if err != nil {
		return "", err
	}

	objID := res.InsertedID.(primitive.ObjectID)

	return ports.ID(objID.Hex()), nil
}