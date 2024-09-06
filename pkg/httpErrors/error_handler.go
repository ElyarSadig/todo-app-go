package httpErrors

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Err    string `json:"error"`
	Status int    `json:"status"`
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		errorValue, ok := err.(RestErr)
	if !ok {
		errResponse := Error{
			Err:    "internal server error",
			Status: http.StatusInternalServerError,
		}
		_ = json.NewEncoder(w).Encode(errResponse)
		return
	}
	w.WriteHeader(errorValue.Status())
	_ = json.NewEncoder(w).Encode(Error{
		Err:    errorValue.ErrorValue(),
		Status: errorValue.Status(),
	})
	}
	
}
