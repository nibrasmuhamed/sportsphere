package service

import (
	"testing"

	"github.com/nibrasmuhamed/sportsphere/mocks"
	models "github.com/nibrasmuhamed/sportsphere/pkg/model"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/mock/gomock"
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
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.repository = mocks.NewMockOperatorRepository(suite.mockCtrl)
	suite.uow = mocks.NewMockUnitOfWork(suite.mockCtrl)
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
		GetOperatorByName(gomock.Eq(operatorRequest.Name)).
		Return(models.Operator{}, mongo.ErrNoDocuments).Times(1)

	// Mock CreateOperator
	suite.repository.EXPECT().
		CreateOperator(gomock.Any(), gomock.AssignableToTypeOf(models.Operator{}), suite.uow).
		Return(nil).Times(1)

	// Mock Commit
	suite.uow.EXPECT().
		Commit().
		Return(nil).Times(1)

	// Act
	response, err := suite.OperatorService.CreateOperator(operatorRequest)

	// Assert
	suite.Require().NoError(err)
	suite.Equal(operatorRequest.Name, response.Name)
	suite.NotEmpty(response.OperatorId)
}
