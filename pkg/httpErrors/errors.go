package httpErrors

import "net/http"

var InternalServerError = NewRestError(http.StatusInternalServerError, "internal server error")
var UserNotFoundError = NewRestError(http.StatusNotFound, "user not found")
