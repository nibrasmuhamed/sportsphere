package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nibrasmuhamed/sportsphere/pkg/helper"
	models "github.com/nibrasmuhamed/sportsphere/pkg/model"
	"github.com/nibrasmuhamed/sportsphere/pkg/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

// RegisterUser godoc
// @Summary Create a new operator
// @Description Register a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param operator body models.RegisterUserRequest true "Operator Details"
// @Success 201 {object} models.RegisterUserResponse
// @Failure 400 {object} helper.BaseResponse
// @Router /api/v1/user [post]
func (c *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.WriteErrorResponse(w, http.StatusBadRequest, helper.CreateErrorResponse("bad_request", err.Error()))
		return
	}

	resp, err := c.userService.RegisterUser(r.Context(), req)
	if err != nil {
		helper.WriteErrorResponse(w, http.StatusInternalServerError, helper.CreateErrorResponse("internal_server_error", err.Error()))
		return
	}

	helper.WriteSuccessResponse(w, resp)
}
