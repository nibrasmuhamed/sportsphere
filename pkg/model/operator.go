package models

import "time"

type OperatorResponse struct {
	OperatorId string `json:"operator_id,omitempty"`
	Name       string `json:"name,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
}

type OperatorRequest struct {
	Name string `json:"name"`
}

type Operator struct {
	OperatorID    string    `json:"operatorId" bson:"operator_id"`
	Name          string    `json:"name" bson:"name"`
	CreatedAt     time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt     time.Time `json:"updatedAt" bson:"updated_at"`
	IsDeactivated bool      `json:"isDeactivated" bson:"is_deactivated"`
}
