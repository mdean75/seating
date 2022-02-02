package grouprepo

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

func (m *MongoDataStore) Save(group domain.Group) (ports.ID, error) {
	g := NewMongoGroupFromDomain(group)

	res, err := m.db.Collection(m.col.Name()).InsertOne(context.TODO(), g)
	if err != nil {
		return "", err
	}

	objId := res.InsertedID.(primitive.ObjectID)

	return ports.ID(objId.Hex()), nil
}