package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nibrasmuhamed/sportsphere/db"
	"github.com/nibrasmuhamed/sportsphere/internal/repository"
	"github.com/nibrasmuhamed/sportsphere/pkg/logger"
	models "github.com/nibrasmuhamed/sportsphere/pkg/model"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(context.Context, models.RegisterUserRequest) (*models.RegisterUserResponse, error)
}

type userService struct {
	unitOfWork db.UnitOfWork
	userRepo   repository.UserRepository
}

func NewUserService(userRepoRepo repository.UserRepository, uow db.UnitOfWork) UserService {
	return &userService{userRepo: userRepoRepo, unitOfWork: uow}
}

func (u *userService) RegisterUser(ctx context.Context, r models.RegisterUserRequest) (*models.RegisterUserResponse, error) {
	if u.userRepo.UserExistsByEmail(r.Email) {
		logger.GetLogger().Error("user with email already exists", zap.String("email", r.Email))
		return nil, fmt.Errorf("user with email %s already exists", r.Email)
	}

	if u.userRepo.UserExistsByUsername(r.UserName) {
		logger.GetLogger().Error("user with username already exists", zap.String("username", r.UserName))
		return nil, fmt.Errorf("user with username %s already exists", r.UserName)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.GetLogger().Error("error while hashing password", zap.Error(err))
		return nil, err
	}
	user := models.User{
		ID:        uuid.New().String(),
		Email:     r.Email,
		CreatedAt: time.Now(),
		UserName:  r.UserName,
		Password:  string(hashedPassword),
		Phone:     r.Phone,
	}

	if err := u.userRepo.CreateUser(ctx, user, nil); err != nil {
		logger.GetLogger().Error("error while creating user", zap.Error(err))
		return nil, err
	}
	logger.GetLogger().Info("user registered successfully", zap.String("email", r.Email))
	return &models.RegisterUserResponse{
		Message: "User registered successfully",
	}, nil
}
