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
	Update(container DataContainer, keyName string, keyValue any, object any, operatorId string, ctx context.Context, t Transaction) error
	Create(container DataContainer, object any, t Transaction) error
	CreateMany(container DataContainer, object []any, t Transaction) error
	Delete(container DataContainer, keyName string, keyValue any, operatorId string, t Transaction) error
	StartSession() (any, error)
}

var db DatabaseClient

func GetDatabase() DatabaseClient {
	return db
}

func SetDatabase(database DatabaseClient) {
	db = database
}
