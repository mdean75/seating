package repository

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"go.mongodb.org/mongo-driver/mongo/readpref"
// )

// type MongoConn struct {
// 	Client *mongo.Client
// }

// func NewMongoDatabase(conn string) (*mongo.Client, error) {
// 	fmt.Println("connecting to mongo at ", conn)
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conn))
// 	if err != nil {
// 		return nil, err
// 	}

// 	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	err = client.Ping(ctx, readpref.Primary())
// 	if err != nil {
// 		return nil, err
// 	}
// 	cl := client

// 	// return &MongoConn{
// 	// 	Client: cl,
// 	// }, nil
// 	return cl, nil
// }
