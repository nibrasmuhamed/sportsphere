package handlers

import (
	"net/http"

	"github.com/nibrasmuhamed/sportsphere/pkg/helper"
)

type HealthCheckController interface {
	HealthCheck(w http.ResponseWriter, r *http.Request)
}

type HealthCheck struct {
	Status string `json:"status"`
}

func NewHealthCheck() HealthCheckController {
	return &HealthCheck{Status: "ok"}
}
func (h *HealthCheck) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	helper.WriteSuccessResponse(w, h)
}
