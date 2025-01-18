package db

import (
	"context"
)

type DatabaseClient interface {
	Connect(uri, dbName string) error
	Close() error
	Ping() error
	Get(container DataContainer, keyName string, keyValue any, object any, operatorId string) error
	GetMany(container DataContainer, keyName string, keyValue string, object any, operatorId string) error
	// Transaction can be nil if not required.
	Update(container DataContainer, keyName string, keyValue any, object any, operatorId string, ctx context.Context, t UnitOfWork) error
	// Transaction can be nil if not required.
	Create(container DataContainer, object any, t UnitOfWork) error
	// Transaction can be nil if not required.
	CreateMany(container DataContainer, object []any, t UnitOfWork) error
	// Transaction can be nil if not required.
	Delete(container DataContainer, keyName string, keyValue any, operatorId string, t UnitOfWork) error
	Count(container DataContainer, keyName string, keyValue any) (int, error)
	StartSession() (any, error)
}

var db DatabaseClient

func GetDatabase() DatabaseClient {
	return db
}

func SetDatabase(database DatabaseClient) {
	db = database
}
