package service

import (
	"context"
	"testing"

	"github.com/nibrasmuhamed/sportsphere/mocks"
	models "github.com/nibrasmuhamed/sportsphere/pkg/model"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type UserServiceTestSuite struct {
	suite.Suite
	context     context.Context
	UserService UserService
	uow         *mocks.MockUnitOfWork
	repository  *mocks.MockUserRepository
	mockCtrl    *gomock.Controller
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}

func (suite *UserServiceTestSuite) SetupTest() {
	suite.context = context.Background()
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.repository = mocks.NewMockUserRepository(suite.mockCtrl)
	suite.uow = mocks.NewMockUnitOfWork(suite.mockCtrl)
	suite.UserService = NewUserService(suite.repository, suite.uow)
}

func (suite *UserServiceTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite *UserServiceTestSuite) TestRegisterUserSuccess() {
	// Arrange: Set up the user request and mock the repository's CreateUser method.
	userRequest := models.RegisterUserRequest{
		UserName: "test",
		Email:    "test@sportsphere.com",
		Password: "password",
		Phone:    "1234567890",
	}

	suite.repository.EXPECT().
		CreateUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	suite.repository.EXPECT().UserExistsByEmail(userRequest.Email).Return(false)
	suite.repository.EXPECT().UserExistsByUsername(userRequest.UserName).Return(false)
	response, err := suite.UserService.RegisterUser(suite.context, userRequest)

	suite.T().Logf("Response: %+v", response)

	suite.NoError(err, "Expected no error, but got: %v", err)

	// Assert the message in the response.
	suite.Equal("User registered successfully", response.Message, "Expected success message but got: %v", response.Message)
}
