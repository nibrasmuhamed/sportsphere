package service

import (
	"context"

	"github.com/nibrasmuhamed/sportsphere/mocks"
	models "github.com/nibrasmuhamed/sportsphere/pkg/model"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type UserServiceTestSuite struct {
	suite.Suite
	UserService UserService
	uow         *mocks.MockUnitOfWork
	repository  *mocks.MockUserRepository
	mockCtrl    *gomock.Controller
}

func (suite *UserServiceTestSuite) SetupTest() {
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

	// Mock the repository method to expect a call and return nil (indicating success).
	suite.repository.EXPECT().CreateUser(gomock.Eq(userRequest), gomock.AssignableToTypeOf(models.User{}), suite.uow).Return(nil).Times(1)

	ctx := context.Background()
	response, err := suite.UserService.RegisterUser(ctx, userRequest)

	suite.T().Logf("Response: %+v", response)

	suite.NoError(err, "Expected no error, but got: %v", err)

	// Assert the message in the response.
	suite.Equal("User registered successfully", response.Message, "Expected success message but got: %v", response.Message)
}
