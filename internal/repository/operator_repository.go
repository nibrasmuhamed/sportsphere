package repository

import (
	"context"
	"time"

	"github.com/nibrasmuhamed/sportsphere/db"
	models "github.com/nibrasmuhamed/sportsphere/pkg/model"
)

//go: go generate mockgen -source=internal/repository/operator_repository.go -destination=../../mocks/mock_operator_repository.go -package=mocks

type OperatorRepository interface {
	// uow can be nil if not required.
	CreateOperator(context.Context, models.Operator, db.UnitOfWork) error
	GetOperatorByName(string) (models.Operator, error)
}

type operatorRepository struct {
	db        db.DatabaseClient
	container db.DataContainer
}

// var OperatorCollection = db.Collection{
// 	CollectionName: "operator",
// 	PrimaryKeyName: "operator_id",
// }

func NewOperatorRepository() OperatorRepository {
	return &operatorRepository{
		db: db.GetDatabase(),
		container: db.Collection{
			CollectionName: "operator",
			PrimaryKeyName: "operator_id",
		},
	}
}

func (o *operatorRepository) CreateOperator(ctx context.Context, operator models.Operator, t db.UnitOfWork) error {
	operator.CreatedAt = time.Now()
	return o.db.Create(o.container, operator, t)
}

func (o *operatorRepository) GetOperatorByName(operatorName string) (models.Operator, error) {
	var operator models.Operator
	err := o.db.Get(o.container, "name", operatorName, &operator, "")
	return operator, err
}
