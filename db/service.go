package db

import (
	"context"
)

var _ DatabaseClient = (*MongoDB)(nil) // Compile-time check to ensure DB implements DatabaseClient

type Collection struct {
	Name           string
	PrimaryKeyName string
	IndexKeyNames  []string
}

type DatabaseClient interface {
	Connect(uri, dbName string) error
	Close() error
	Ping() error
	Get(collection Collection, keyName string, keyValue any, object any, operatorId string) error
	GetMany(collection Collection, keyName string, keyValue string, object any, operatorId string) error
	Update(collection Collection, keyName string, keyValue any, object any, operatorId string, ctx context.Context) error
	Create(collection Collection, keyValue any, object any) error
	CreateMany(collection Collection, keyValue any, object []any) error
	Delete(collection Collection, keyName string, keyValue any, operatorId string) error
}

var db DatabaseClient

func GetDatabase() DatabaseClient {
	return db
}

func SetDatabase(database DatabaseClient) {
	db = database
}
