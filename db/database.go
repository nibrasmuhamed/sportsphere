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
	Update(container DataContainer, keyName string, keyValue any, object any, operatorId string, ctx context.Context) error
	Create(container DataContainer, object any) error
	CreateMany(container DataContainer, keyValue any, object []any) error
	Delete(container DataContainer, keyName string, keyValue any, operatorId string) error
}

var db DatabaseClient

func GetDatabase() DatabaseClient {
	return db
}

func SetDatabase(database DatabaseClient) {
	db = database
}
