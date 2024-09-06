package httpErrors

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Err    string `json:"error"`
	Status int    `json:"status"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func ReturnError(w http.ResponseWriter, err error) {
	errorValue, ok := err.(RestErr)
	if !ok {
		errResponse := ErrorResponse{
			Err:    "internal server error",
			Status: http.StatusInternalServerError,
		}
		_ = json.NewEncoder(w).Encode(errResponse)
		return
	}
	w.WriteHeader(errorValue.Status())
	_ = json.NewEncoder(w).Encode(ErrorResponse{
		Err:    errorValue.ErrorValue(),
		Status: errorValue.Status(),
	})
}

func ReturnSuccess(w http.ResponseWriter, message any) {
	w.WriteHeader(http.StatusOK)
	if message != nil {
		json.NewEncoder(w).Encode(message)
		return
	}
	json.NewEncoder(w).Encode(SuccessResponse{
		Message: "success",
		Status:  http.StatusOK,
	})
}
