package try

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// This is where I will try to work on the database connection by running the NewMongo client

const (
	db_name             = "Log_db"
	db_collection_name  = "Log_ollection"
	internalServerError = "internal server error"
)

type MongoDbConnection struct {
	client *mongo.Client
}

func NewMongoClient(cl *mongo.Client) *MongoDbConnection {
	return &MongoDbConnection{
		client: cl,
	}
}

func (m *MongoDbConnection) Add(dataLog Log) (string, error) {
	fmt.Println(dataLog)
	collection := m.client.Database(db_name).Collection(db_collection_name)
	_, err := collection.InsertOne(context.Background(), dataLog)

	if err != nil {
		return "", err
	}
	return "It has been inserted already", nil
}

func (m *MongoDbConnection) List() (Logs, error) {
	collection := m.client.Database(db_name).Collection(db_collection_name)
	query := bson.M{}
	curs, err := collection.Find(context.Background(), query)
	if err != nil {
		return nil, err
	}
	var values Logs
	//  I could have done it by using
	// err := cursor.All(ctx, &res); err != nil
	// Just wanted tolearn how to work with cursor
	for curs.Next(context.Background()) {
		var value Log
		if err := curs.Decode(&value); err != nil {
			return nil, err
		}
		values = append(values, value)
		if err := curs.Err(); err != nil {
			return nil, err
		}
	}
	return values, nil
}
