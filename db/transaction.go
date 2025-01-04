package db

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

type Transaction interface {
	QueueCreate(container DataContainer, object any)
	QueueCreateMany(container DataContainer, object []any)
	QueueUpdate(container DataContainer, keyName string, keyValue any, object any, operatorId string)
	QueueDelete(container DataContainer, keyName string, keyValue any, operatorId string)
	Commit() error
	Abort() error
}

// MongoTransaction wraps a MongoDB session for transactional operations
type MongoTransaction struct {
	session    mongo.Session
	ctx        context.Context
	db         DatabaseClient
	operations []func() error // Queue to store operations
}

// StartTransaction initiates a new transaction
func NewTransaction(ctx context.Context) (Transaction, error) {
	sessionType, err := GetDatabase().StartSession()
	if err != nil {
		return nil, err
	}

	session, ok := sessionType.(mongo.Session)
	if !ok {
		return nil, errors.New("invalid session type")
	}

	sessionCtx := mongo.NewSessionContext(ctx, session)
	return &MongoTransaction{
		session:    session,
		ctx:        sessionCtx,
		db:         GetDatabase(),
		operations: []func() error{}, // Initialize empty operations list
	}, nil
}

// QueueCreate adds a create operation to the transaction queue
func (tx *MongoTransaction) QueueCreate(container DataContainer, object any) {
	tx.operations = append(tx.operations, func() error {
		err := tx.db.Create(container, object, nil)
		return err
	})
}

// QueueUpdate adds an update operation to the transaction queue
func (tx *MongoTransaction) QueueUpdate(container DataContainer, keyName string, keyValue any, object any, operatorId string) {
	tx.operations = append(tx.operations, func() error {
		err := tx.db.Update(container, keyName, keyValue, object, operatorId, tx.ctx, nil)
		return err
	})
}

// QueueDelete adds a delete operation to the transaction queue
func (tx *MongoTransaction) QueueDelete(container DataContainer, keyName string, keyValue any, operatorId string) {
	tx.operations = append(tx.operations, func() error {
		err := tx.db.Delete(container, keyName, keyValue, operatorId, nil)
		return err
	})
}
func (tx *MongoTransaction) QueueCreateMany(container DataContainer, objects []any) {
	tx.operations = append(tx.operations, func() error {
		err := tx.db.CreateMany(container, objects, nil)
		return err
	})
}

// Commit executes all queued operations in sequence, rolling back on any error
func (tx *MongoTransaction) Commit() error {
	// Start the transaction
	err := tx.session.StartTransaction()
	if err != nil {
		return err
	}

	// Execute all queued operations
	for _, operation := range tx.operations {
		if err := operation(); err != nil {
			// If any operation fails, abort the transaction and return the error
			tx.Abort()
			return err
		}
	}

	// Commit if all operations succeed
	return tx.session.CommitTransaction(tx.ctx)
}

// Abort rolls back the transaction
func (tx *MongoTransaction) Abort() error {
	return tx.session.AbortTransaction(tx.ctx)
}
