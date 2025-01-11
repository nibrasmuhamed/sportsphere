package db

import (
	"context"
	"errors"

	"github.com/nibrasmuhamed/sportsphere/config"
	"github.com/nibrasmuhamed/sportsphere/pkg/constants"
)

type UnitOfWork interface {
	QueueCreate(container DataContainer, object any)
	QueueCreateMany(container DataContainer, object []any)
	QueueUpdate(container DataContainer, keyName string, keyValue any, object any, operatorId string)
	QueueDelete(container DataContainer, keyName string, keyValue any, operatorId string)
	Commit() error
	Abort() error
}

func NewUnitOfWork(ctx context.Context) (UnitOfWork, error) {
	switch config.Get().DatabaseType {
	case string(constants.MONGODB):
		return NewMongoUOW(ctx)
	default:
		return nil, errors.New("invalid database type")
	}
}
