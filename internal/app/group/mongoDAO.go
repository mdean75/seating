package group

import (
	"context"
	"seating/internal/db"

	// "seating/internal/app/group"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDataStore struct {
	*db.MongoConn
	db  *mongo.Database
	col *mongo.Collection
}

func NewDAO(dbconn *db.MongoConn, db, col string) Repository {
	dbx := dbconn.Client.Database(db)
	conx := dbx.Collection(col)

	return &MongoDataStore{dbconn, dbx, conx}
}

func NewGroup(displayName, shortName string) Group {
	return Group{
		DisplayName: displayName,
		ShortName: shortName,
	}
}

func (m *MongoDataStore) CreateGroup(displayName, shortName string) (ID, error) {
	group := NewGroup(displayName, shortName)

	res, err := m.db.Collection(m.col.Name()).InsertOne(context.TODO(), group)
	if err != nil {
		return "", err
	}

	objId := res.InsertedID.(primitive.ObjectID)

	return ID(objId.Hex()), nil
}