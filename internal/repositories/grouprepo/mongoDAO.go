package grouprepo

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

func NewDAO(dbconn *db.MongoConn, db, col string) ports.GroupRepository {
	dbx := dbconn.Client.Database(db)
	conx := dbx.Collection(col)

	return &MongoDataStore{dbconn, dbx, conx}
}

type Group struct {
	ID string `bson:"_id,omitempty"`
	DisplayName string `bson:"displayName"`
	ShortName string `bson:"shortName"`
}

func NewMongoGroupFromDomain(group domain.Group) Group {
	return Group{
		ID: group.ID,
		DisplayName: group.DisplayName,
		ShortName: group.ShortName,
	}
}

func convertMongoGroupToDomain(group Group) domain.Group {
	return domain.NewGroup(group.ID, group.DisplayName, group.ShortName)
}

func (m *MongoDataStore) Save(group domain.Group) (ports.ID, error) {
	g := NewMongoGroupFromDomain(group)

	res, err := m.db.Collection(m.col.Name()).InsertOne(context.TODO(), g)
	if err != nil {
		return "", err
	}

	objId := res.InsertedID.(primitive.ObjectID)

	return ports.ID(objId.Hex()), nil
}

func (m *MongoDataStore) Get(groupID string) (domain.Group, error) {
	var group Group
	id, err := primitive.ObjectIDFromHex(groupID)
	if err != nil {
		return domain.Group{}, err
	}

	result := m.col.FindOne(context.TODO(), bson.M{"_id": id})//.Decode(&group)
	if result.Err() != nil {
		return domain.Group{}, result.Err()
	}

	if err := result.Decode(&group); err != nil {
		return domain.Group{}, err
	}

	return convertMongoGroupToDomain(group), nil
}


var ErrUnableToDeleteResouce error = fmt.Errorf("unable to delete the specified resource")

func (m *MongoDataStore) Delete(groupID string) error {
	id, err := primitive.ObjectIDFromHex(groupID)
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