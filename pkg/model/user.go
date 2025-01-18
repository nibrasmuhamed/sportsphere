package models

import "time"

type User struct {
	ID            string    `json:"id" bson:"_id"`
	UserName      string    `json:"userName" bson:"user_name"`
	Email         string    `json:"email" bson:"email"`
	Phone         string    `json:"phone" bson:"phone"`
	Password      string    `json:"password" bson:"password"`
	OperatorID    string    `json:"operatorId,omitempty" bson:"operator_id,omitempty"`
	CreatedAt     time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt     time.Time `json:"updatedAt" bson:"updated_at"`
	IsDeactivated bool      `json:"isDeactivated" bson:"is_deactivated"`
}

type RegisterUserRequest struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}
type RegisterUserResponse struct {
	Message string `json:"message"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Token   string `json:"token"`
	IdToken string `json:"idToken"`
}
