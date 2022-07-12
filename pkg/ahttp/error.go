package ahttp

import (
	"net/http"
)

// Standard Errors
var ErrBadRequest = ErrorResponse{
	Status:  http.StatusBadRequest,
	Code:    "bad_request",
	Message: "Bad Request",
}
var ErrStatusNoContent = ErrorResponse{
	Status:  http.StatusNoContent,
	Code:    "no_content",
	Message: "No Content",
}
var ErrUnauthorized = ErrorResponse{
	Status:  http.StatusUnauthorized,
	Code:    "unauthorized",
	Message: "Unauthorized",
}
var ErrForbidden = ErrorResponse{
	Status:  http.StatusForbidden,
	Code:    "forbidden",
	Message: "Forbidden",
}
var ErrNotAcceptable = ErrorResponse{
	Status:  http.StatusNotAcceptable,
	Code:    "not_acceptable",
	Message: "Not Acceptable",
}
var ErrInternalServer = ErrorResponse{
	Status:  http.StatusInternalServerError,
	Code:    "internal_error",
	Message: "Internal Error",
}
