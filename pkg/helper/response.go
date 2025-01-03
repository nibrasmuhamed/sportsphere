package helper

import (
	"encoding/json"
	"net/http"
)

type BaseResponse struct {
	ErrorCode    string `json:"error_code,omitempty"`
	ErrorType    string `json:"error_type,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

func CreateErrorResponse(code string, message string) BaseResponse {
	return BaseResponse{
		ErrorCode:    code,
		ErrorMessage: message,
	}
}

func WriteSuccessResponse(w http.ResponseWriter, response any) error {
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, error BaseResponse) error {
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(error)
}

func WriteCreatedResponse(w http.ResponseWriter, response any) error {
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(response)
}

func WriteNoContentResponse(w http.ResponseWriter, response any) error {
	w.WriteHeader(http.StatusNoContent)
	return json.NewEncoder(w).Encode(response)
}
