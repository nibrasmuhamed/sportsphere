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

var _ DatabaseClient = (*MongoDB)(nil)   // Compile-time check to ensure DB implements DatabaseClient
var _ DataContainer = (*Collection)(nil) // Compile-time check to ensure Collection implements DataContainer

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

type Collection struct {
	CollectionName string
	PrimaryKeyName string
	IndexKeyNames  []string
}

func (c Collection) Name() string {
	return c.CollectionName
}
func NewMongoDB() *MongoDB {
	return &MongoDB{}
}

func (db *MongoDB) StartSession() (any, error) {
	return db.client.StartSession()
}

// // QueueCreate adds a create operation to the transaction queue
// func (tx *MongoDB) QueueCreate(container DataContainer, object any) {
// 	tx.operations = append(tx.operations, func() error {
// 		collection := tx.database.Collection(container.Name())
// 		_, err := collection.InsertOne(tx.mongoCtx, object)
// 		return err
// 	})
// }

// // QueueUpdate adds an update operation to the transaction queue
// func (tx *MongoDB) QueueUpdate(container DataContainer, keyName string, keyValue any, object any, operatorId string) {
// 	tx.operations = append(tx.operations, func() error {
// 		collection := tx.database.Collection(container.Name())
// 		filter := bson.M{keyName: keyValue}
// 		if operatorId != "" {
// 			filter["operatorid"] = operatorId
// 		}
// 		_, err := collection.UpdateOne(tx.mongoCtx, filter, bson.M{"$set": object})
// 		return err
// 	})
// }

// // QueueDelete adds a delete operation to the transaction queue
// func (tx *MongoDB) QueueDelete(container DataContainer, keyName string, keyValue any, operatorId string) {
// 	tx.operations = append(tx.operations, func() error {
// 		collection := tx.database.Collection(container.Name())
// 		filter := bson.M{keyName: keyValue}
// 		if operatorId != "" {
// 			filter["operatorid"] = operatorId
// 		}
// 		_, err := collection.DeleteOne(tx.mongoCtx, filter)
// 		return err
// 	})
// }

func (tx *MongoDB) Commit(f []func() error) error {
	// Start the transaction
	session, err := tx.client.StartSession()
	if err != nil {
		return err
	}

	// Execute all queued operations
	for _, operation := range f {
		if err := operation(); err != nil {
			// If any operation fails, abort the transaction and return the error
			session.AbortTransaction(tx.mongoCtx)
			return err
		}
	}

	// Commit if all operations succeed
	return session.CommitTransaction(tx.mongoCtx)
}

func (db *MongoDB) Connect(uri, dbName string) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

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
	container DataContainer,
	keyName string,
	keyValue any,
	object any,
	operatorId string) error {

	dbCollection := db.database.Collection(container.Name())
	filter := bson.M{}
	filter[strings.ToLower(keyName)] = keyValue
	if operatorId != "" {
		filter["operatorid"] = operatorId
	}
	result := dbCollection.FindOne(context.Background(), filter)
	if result != nil {
		err := result.Decode(object)
		return err
	}
	return fmt.Errorf("no db entry for key: %v and value: %v", keyName, keyValue)
}

func (db *MongoDB) GetMany(
	container DataContainer,
	keyName,
	keyValue string,
	object any,
	operatorId string) error {
	filter := bson.M{}
	filter[strings.ToLower(keyName)] = keyValue
	dbCollection := db.database.Collection(container.Name())

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

// func (db *MongoDB) GetAll(container DataContainer, object any, operatorId string) error {

// }

func (db *MongoDB) Update(
	container DataContainer,
	keyName string,
	keyValue any,
	object any,
	operatorId string,
	ctx context.Context, t UnitOfWork) error {
	if t != nil {
		t.QueueUpdate(container, keyName, keyValue, object, operatorId)
		return nil
	}
	dbCollection := db.database.Collection(container.Name())
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
	container DataContainer,
	object any, t UnitOfWork) error {
	if t != nil {
		t.QueueCreate(container, object)
		return nil
	}
	dbCollection := db.database.Collection(container.Name())
	_, err := dbCollection.InsertOne(context.Background(), object)
	return err
}

func (db *MongoDB) CreateMany(container DataContainer, object []any, t UnitOfWork) error {
	if t != nil {
		t.QueueCreateMany(container, object)
		return nil
	}
	dbCollection := db.database.Collection(container.Name())
	_, err := dbCollection.InsertMany(db.mongoCtx, object)
	return err
}

func (db *MongoDB) Delete(
	container DataContainer,
	keyName string,
	keyValue any,
	operatorId string, t UnitOfWork) error {
	if t != nil {
		t.QueueDelete(container, keyName, keyValue, operatorId)
		return nil
	}
	dbCollection := db.database.Collection(container.Name())
	filter := bson.M{}
	filter[strings.ToLower(keyName)] = keyValue
	if operatorId != "" {
		filter["operatorid"] = operatorId
	}
	_, err := dbCollection.DeleteOne(db.mongoCtx, filter)
	return err
}
