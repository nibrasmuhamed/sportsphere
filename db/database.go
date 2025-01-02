package db

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	mongoCtx                 context.Context
	client                   *mongo.Client
	database                 *mongo.Database
	ConnectionCreated        int
	ConnectionPoolCreated    int
	ConnectionClosed         int
	ConnectionReady          int
	ConnectionCheckOutFailed int
	ConnectionCheckedOut     int
	ConnectionCheckedIn      int
	ConnectionPoolCleared    int
	ConnectionPoolClosed     int
	checkedOut               []uint64
}

func (db *MongoDB) Connect(uri, dbName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	db.client = client
	db.mongoCtx = ctx
	db.database = client.Database(dbName)
	db.ConnectionReady = 1
	fmt.Println("Connected to MongoDB!")
	return nil
}

func (db *MongoDB) Close() error {
	if db.client != nil {
		err := db.client.Disconnect(db.mongoCtx)
		if err != nil {
			return fmt.Errorf("failed to disconnect MongoDB: %w", err)
		}
		db.ConnectionClosed = 1
		fmt.Println("Disconnected from MongoDB!")
	}
	return nil
}

func (db *MongoDB) Ping() error {
	err := db.client.Ping(db.mongoCtx, nil)
	if err != nil {
		return fmt.Errorf("MongoDB connection ping failed: %w", err)
	}
	fmt.Println("MongoDB is reachable!")
	return nil
}

func (db *MongoDB) Get(
	collection Collection,
	keyName string,
	keyValue any,
	object any,
	operatorId string) error {
	dbCollection := db.database.Collection(collection.Name)
	filter := bson.M{}
	filter[strings.ToLower(keyName)] = keyValue
	if operatorId != "" {
		filter["operatorid"] = operatorId
	}
	result := dbCollection.FindOne(db.mongoCtx, filter)
	if result != nil {
		err := result.Decode(object)
		return err
	}
	return fmt.Errorf("no db entry for key: %v and value: %v", keyName, keyValue)
}

func (db *MongoDB) GetMany(
	collection Collection,
	keyName,
	keyValue string,
	object any,
	operatorId string) error {
	filter := bson.M{}
	filter[strings.ToLower(keyName)] = keyValue
	dbCollection := db.database.Collection(collection.Name)

	// Add operatorId to the filter if provided
	if operatorId != "" {
		filter["operatorid"] = operatorId
	}

	result, err := dbCollection.Find(db.mongoCtx, filter)
	if err != nil {
		return fmt.Errorf("failed to fetch documents: %w", err)
	}
	defer result.Close(db.mongoCtx)

	if result != nil {
		err = result.All(db.mongoCtx, object)
	}
	return err
}

// func (db *MongoDB) GetAll(collection Collection, object any, operatorId string) error {

// }

func (db *MongoDB) Update(
	collection Collection,
	keyName string,
	keyValue any,
	object any,
	operatorId string,
	ctx context.Context) error {

	dbCollection := db.database.Collection(collection.Name)
	pk := bson.M{strings.ToLower(keyName): keyValue}
	conditionArr := []bson.M{}
	conditionArr = append(conditionArr, pk)
	filter := bson.M{}
	if operatorId != "" {
		filter["operatorid"] = operatorId
	}
	conditionArr = append(conditionArr, filter)
	_, err := dbCollection.UpdateOne(db.mongoCtx,
		bson.M{"$and": conditionArr},
		bson.M{"$set": object},
	)
	return err
}

func (db *MongoDB) Create(
	collection Collection,
	keyValue any,
	object any) error {
	dbCollection := db.database.Collection(collection.Name)
	_, err := dbCollection.InsertOne(db.mongoCtx, object)
	return err
}

func (db *MongoDB) CreateMany(collection Collection, keyValue any, object []any) error {
	dbCollection := db.database.Collection(collection.Name)
	_, err := dbCollection.InsertMany(db.mongoCtx, object)
	return err
}

func (db *MongoDB) Delete(
	collection Collection,
	keyName string,
	keyValue any,
	operatorId string) error {
	dbCollection := db.database.Collection(collection.Name)
	filter := bson.M{}
	filter[strings.ToLower(keyName)] = keyValue
	if operatorId != "" {
		filter["operatorid"] = operatorId
	}
	_, err := dbCollection.DeleteOne(db.mongoCtx, filter)
	return err
}
