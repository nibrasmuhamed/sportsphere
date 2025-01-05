package db

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

var _ UnitOfWork = (*MongoUOW)(nil) // Compile-time check to ensure MongoUOW implements UnitOfWork

// MongoUOW wraps a MongoDB session for transactional operations
type MongoUOW struct {
	session    mongo.Session
	ctx        context.Context
	db         DatabaseClient
	operations []func() error // Queue to store operations
}

// StartTransaction initiates a new transaction
func NewMongoUOW(ctx context.Context) (UnitOfWork, error) {
	sessionType, err := GetDatabase().StartSession()
	if err != nil {
		return nil, err
	}

	session, ok := sessionType.(mongo.Session)
	if !ok {
		return nil, errors.New("invalid session type")
	}

	sessionCtx := mongo.NewSessionContext(ctx, session)
	return &MongoUOW{
		session:    session,
		ctx:        sessionCtx,
		db:         GetDatabase(),
		operations: []func() error{}, // Initialize empty operations list
	}, nil
}

// QueueCreate adds a create operation to the transaction queue
func (tx *MongoUOW) QueueCreate(container DataContainer, object any) {
	tx.operations = append(tx.operations, func() error {
		err := tx.db.Create(container, object, nil)
		// if err != nil {
		// 	logger.Error("error while creating object", zap.Any("error", err), zap.Any("from", c.Value("stage")))
		// }
		return err
	})
}

// QueueUpdate adds an update operation to the transaction queue
func (tx *MongoUOW) QueueUpdate(container DataContainer, keyName string, keyValue any, object any, operatorId string) {
	tx.operations = append(tx.operations, func() error {
		err := tx.db.Update(container, keyName, keyValue, object, operatorId, tx.ctx, nil)
		return err
	})
}

// QueueDelete adds a delete operation to the transaction queue
func (tx *MongoUOW) QueueDelete(container DataContainer, keyName string, keyValue any, operatorId string) {
	tx.operations = append(tx.operations, func() error {
		err := tx.db.Delete(container, keyName, keyValue, operatorId, nil)
		return err
	})
}
func (tx *MongoUOW) QueueCreateMany(container DataContainer, objects []any) {
	tx.operations = append(tx.operations, func() error {
		err := tx.db.CreateMany(container, objects, nil)
		return err
	})
}

// Commit executes all queued operations in sequence, rolling back on any error
func (tx *MongoUOW) Commit() error {
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
func (tx *MongoUOW) Abort() error {
	return tx.session.AbortTransaction(tx.ctx)
}
