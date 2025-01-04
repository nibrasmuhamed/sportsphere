package repository

import (
	"context"
	"time"

	"github.com/mohdjishin/sportsphere/db"
	models "github.com/mohdjishin/sportsphere/pkg/model"
)

type OperatorRepository interface {
	// Transaction can be nil if not required.
	CreateOperator(context.Context, models.Operator, db.Transaction) error
	GetOperatorByName(string) (models.Operator, error)
}

type operatorRepository struct {
	db db.DatabaseClient
}

var OperatorCollection = db.Collection{
	CollectionName: "operator",
	PrimaryKeyName: "operator_id",
}

func NewOperatorRepository() OperatorRepository {
	return &operatorRepository{db: db.GetDatabase()}
}

func (o *operatorRepository) CreateOperator(ctx context.Context, operator models.Operator, t db.Transaction) error {
	operator.CreatedAt = time.Now()
	return o.db.Create(OperatorCollection, operator, t)
}

// func (o *operatorRepository) CreateOperatorWithTransaction(ctx context.Context, operator models.Operator) error {
// 	return o.db.WithTransaction(ctx, func(tx db.Transaction) error {
// 		operator.CreatedAt = time.Now()
// 		return tx.Create(OperatorCollection, operator)
// 	})
// }

func (o *operatorRepository) GetOperatorByName(operatorName string) (models.Operator, error) {
	var operator models.Operator
	err := o.db.Get(OperatorCollection, "name", operatorName, &operator, "")
	return operator, err
}
