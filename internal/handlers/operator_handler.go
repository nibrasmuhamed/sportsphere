package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nibrasmuhamed/sportsphere/pkg/helper"
	models "github.com/nibrasmuhamed/sportsphere/pkg/model"
	"github.com/nibrasmuhamed/sportsphere/pkg/service"
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

// CreateOperator godoc
// @Summary Create a new operator
// @Description Add a new operator with details
// @Tags Operator
// @Accept json
// @Produce json
// @Param operator body models.OperatorRequest true "Operator Details"
// @Success 201 {object} models.OperatorResponse
// @Failure 400 {object} helper.BaseResponse
// @Failure 500 {object} helper.BaseResponse
// @Router /api/v1/operator [post]
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
