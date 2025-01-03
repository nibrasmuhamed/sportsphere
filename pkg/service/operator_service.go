package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mohdjishin/sportsphere/internal/repository"
	models "github.com/mohdjishin/sportsphere/pkg/model"
)

type OperatorService interface {
	CreateOperator(models.OperatorRequest) (models.OperatorResponse, error)
}

type operatorService struct {
	operatorRepo repository.OperatorRepository
}

func NewOperatorService(operatorRepo repository.OperatorRepository) OperatorService {
	return &operatorService{operatorRepo: operatorRepo}
}
func (o *operatorService) CreateOperator(operatorRequest models.OperatorRequest) (models.OperatorResponse, error) {
	operatorModel := models.Operator{
		OperatorID: uuid.New().String(),
		Name:       operatorRequest.Name,
	}
	_, err := o.operatorRepo.GetOperatorByName(operatorRequest.Name)
	if err == nil {
		return models.OperatorResponse{}, fmt.Errorf("operator with name %s already exists", operatorRequest.Name)
	}
	err = o.operatorRepo.CreateOperator(operatorModel)
	if err != nil {
		return models.OperatorResponse{}, err
	}
	return models.OperatorResponse{
		OperatorId: operatorModel.OperatorID,
		Name:       operatorModel.Name,
	}, nil
}
