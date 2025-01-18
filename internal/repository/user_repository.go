package repository

import (
	"context"
	"time"

	"github.com/nibrasmuhamed/sportsphere/db"
	models "github.com/nibrasmuhamed/sportsphere/pkg/model"
)

//go:generate mockgen -source=internal/repository/user_repository.go -destination=../../mocks/mock_user_repository.go -package=mocks

type UserRepository interface {
	CreateUser(context.Context, models.User, db.UnitOfWork) error
	GetUserByEmail(string) (*models.User, error)
	UserExistsByEmail(string) bool
	UserExistsByUsername(string) bool
}

type userRepository struct {
	db        db.DatabaseClient
	container db.DataContainer
}

func NewUserRepository() UserRepository {
	return &userRepository{
		db: db.GetDatabase(),
		container: db.Collection{
			CollectionName: "user",
			PrimaryKeyName: "user_id",
		},
	}
}

func (u *userRepository) CreateUser(ctx context.Context, user models.User, t db.UnitOfWork) error {
	user.CreatedAt = time.Now()
	return u.db.Create(u.container, user, t)
}

func (u *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := u.db.Get(u.container, "email", email, &user, "")
	return &user, err
}

func (u *userRepository) UserExistsByEmail(email string) bool {
	count, err := u.db.Count(u.container, "email", email)
	return err == nil && count > 0
}

func (u *userRepository) UserExistsByUsername(username string) bool {
	count, err := u.db.Count(u.container, "user_name", username)
	return err == nil && count > 0
}
