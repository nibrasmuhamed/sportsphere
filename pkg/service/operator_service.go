package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mohdjishin/sportsphere/db"
	"github.com/mohdjishin/sportsphere/internal/repository"
	"github.com/mohdjishin/sportsphere/pkg/logger"
	models "github.com/mohdjishin/sportsphere/pkg/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type OperatorService interface {
	CreateOperator(models.OperatorRequest) (models.OperatorResponse, error)
}

type operatorService struct {
	unitOfWork   db.UnitOfWork
	operatorRepo repository.OperatorRepository
}

func NewOperatorService(operatorRepo repository.OperatorRepository, uow db.UnitOfWork) OperatorService {
	return &operatorService{operatorRepo: operatorRepo, unitOfWork: uow}
}

func (o *operatorService) CreateOperator(operatorRequest models.OperatorRequest) (models.OperatorResponse, error) {
	// uow, err := db.NewUnitOfWork(context.Background())
	// if err != nil {
	// 	logger.Error("error while creating transaction", zap.Error(err))
	// 	return models.OperatorResponse{}, err
	// }

	operatorModel := models.Operator{
		OperatorID: uuid.New().String(),
		Name:       operatorRequest.Name,
	}

	dbOperator, err := o.operatorRepo.GetOperatorByName(operatorRequest.Name)
	logger.Info("dbOperator", zap.Any("dbOperator", dbOperator))
	if err == nil && operatorModel.Name == dbOperator.Name {
		logger.Error("operator with name already exists", zap.String("name", operatorRequest.Name))
		return models.OperatorResponse{}, fmt.Errorf("operator with name %s already exists", operatorRequest.Name)
	}
	if err != nil && err != mongo.ErrNoDocuments {
		logger.Error("error while fetching operator", zap.Error(err))
		return models.OperatorResponse{}, err
	}

	err = o.operatorRepo.CreateOperator(context.Background(), operatorModel, o.unitOfWork)
	if err != nil {
		logger.Error("error while creating operator", zap.Error(err))
		return models.OperatorResponse{}, err
	}

	err = o.unitOfWork.Commit()
	if err != nil {
		logger.Error("error while committing transaction", zap.Error(err))
		return models.OperatorResponse{}, err
	}
	return models.OperatorResponse{
		OperatorId: operatorModel.OperatorID,
		Name:       operatorModel.Name,
	}, nil
}
