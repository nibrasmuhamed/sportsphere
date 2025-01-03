package repository

import (
	"time"

	"github.com/mohdjishin/sportsphere/db"
	models "github.com/mohdjishin/sportsphere/pkg/model"
)

type OperatorRepository interface {
	CreateOperator(models.Operator) error
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

func (o *operatorRepository) CreateOperator(operator models.Operator) error {
	operator.CreatedAt = time.Now()
	return o.db.Create(OperatorCollection, operator)
}

func (o *operatorRepository) GetOperatorByName(operatorName string) (models.Operator, error) {
	var operator models.Operator
	err := o.db.Get(OperatorCollection, "name", operatorName, &operator, "")
	return operator, err
}
