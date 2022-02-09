package industryrepo

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

func NewDAO(dbconn *db.MongoConn, db, col string) ports.IndustryRepository {
	dbx := dbconn.Client.Database(db)
	conx := dbx.Collection(col)

	return &MongoDataStore{dbconn, dbx, conx}
}

type Industry struct {
	ID   string `bson:"_id,omitempty"`
	Name string `bson:"name"`
}

func NewIndustry(id, name string) Industry {
	return Industry{
		ID:   id,
		Name: name,
	}
}

func convertDomainIndustryToMongo(industry domain.Industry) Industry {
	return NewIndustry(industry.ID, industry.Name)
}

func convertMongoIndustryToDomain(industry Industry) domain.Industry {
	return domain.NewIndustry(industry.ID, industry.Name)
}

func (m *MongoDataStore) Save(industry domain.Industry) (ports.ID, error) {
	a := convertDomainIndustryToMongo(industry)

	res, err := m.col.InsertOne(context.TODO(), a)
	if err != nil {
		return "", err
	}

	objID := res.InsertedID.(primitive.ObjectID)

	return ports.ID(objID.Hex()), nil
}

func (m *MongoDataStore) Get(industryID string) (domain.Industry, error) {
	var industry Industry
	id, err := primitive.ObjectIDFromHex(industryID)
	if err != nil {
		return domain.Industry{}, err
	}

	err = m.col.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&industry)
	if err != nil {
		return domain.Industry{}, err
	}

	return convertMongoIndustryToDomain(industry), nil
}

var ErrUnableToDeleteResouce error = fmt.Errorf("unable to delete the specified resource")

func (m *MongoDataStore) Delete(industryID string) error {
	id, err := primitive.ObjectIDFromHex(industryID)
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
