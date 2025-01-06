package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mohdjishin/sportsphere/mocks"
	models "github.com/mohdjishin/sportsphere/pkg/model"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type OperatorServiceTestSuite struct {
	suite.Suite
	OperatorService OperatorService
	uow             *mocks.MockUnitOfWork
	repository      *mocks.MockOperatorRepository
	mockCtrl        *gomock.Controller
}

func TestOperatorService(t *testing.T) {
	suite.Run(t, new(OperatorServiceTestSuite))
}

func (suite *OperatorServiceTestSuite) SetupTest() {
	// Initialize gomock.Controller and assign to mockCtrl
	suite.mockCtrl = gomock.NewController(suite.T())

	// Create mocks using the same controller
	suite.repository = mocks.NewMockOperatorRepository(suite.mockCtrl)
	suite.uow = mocks.NewMockUnitOfWork(suite.mockCtrl)

	// Initialize the service with the mocked repository
	suite.OperatorService = NewOperatorService(suite.repository, suite.uow)
}

func (suite *OperatorServiceTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}
func (suite *OperatorServiceTestSuite) TestCreateOperatorSuccess() {
	operatorRequest := models.OperatorRequest{
		Name: "test",
	}

	// Mock GetOperatorByName to simulate operator not existing
	suite.repository.EXPECT().
		GetOperatorByName(operatorRequest.Name).
		Return(models.Operator{}, mongo.ErrNoDocuments).Times(1)

	// Mock CreateOperator
	suite.repository.EXPECT().
		CreateOperator(gomock.Any(), gomock.Any(), suite.uow).
		DoAndReturn(func(ctx interface{}, operator models.Operator, uow interface{}) error {
			suite.T().Logf("CreateOperator called with operator=%v, uow=%v", operator, uow)
			return nil
		})

	// Mock Commit
	suite.uow.EXPECT().
		Commit().
		Return(nil)

	// Act
	response, err := suite.OperatorService.CreateOperator(operatorRequest)

	// Assert
	suite.Nil(err)
	suite.Equal(operatorRequest.Name, response.Name)
	suite.NotEmpty(response.OperatorId)
}
