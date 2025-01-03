package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mohdjishin/sportsphere/pkg/helper"
	models "github.com/mohdjishin/sportsphere/pkg/model"
	"github.com/mohdjishin/sportsphere/pkg/service"
)

type OperatorHandler interface {
	CreateOperator(w http.ResponseWriter, r *http.Request)
}

type OperatorController struct {
	operatorService service.OperatorService
}

func NewOperatorHandler(operatorService service.OperatorService) OperatorHandler {
	return &OperatorController{operatorService: operatorService}
}

func (o *OperatorController) CreateOperator(w http.ResponseWriter, r *http.Request) {
	operatorRequest := models.OperatorRequest{}
	err := json.NewDecoder(r.Body).Decode(&operatorRequest)
	if err != nil {
		helper.WriteErrorResponse(w, http.StatusBadRequest, helper.CreateErrorResponse("bad_request", err.Error()))
		return
	}
	operatorResponse, err := o.operatorService.CreateOperator(operatorRequest)
	if err != nil {
		helper.WriteErrorResponse(w, http.StatusInternalServerError, helper.CreateErrorResponse("internal_server_error", err.Error()))
		return
	}
	helper.WriteSuccessResponse(w, operatorResponse)
}
