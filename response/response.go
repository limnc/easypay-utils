package response

// This will covered the topic of return a standardize HTTP response for the API.
// The following format wii be ErrorResponse and SuccessResponse

import (
	"net/http"
	"time"
)

//VER : 0.2.0

type Response struct {
	Success    bool         `json:"success"`
	Data       interface{}  `json:"data"`
	Error      *ErrorDetail `json:"error"`
	StatusCode int          `json:"status_code"`
	RequestID  string       `json:"request_id"`
	Timestamp  time.Time    `json:"timestamp"`
	Size       int          `json:"size"`
}

type ErrorDetail struct {
	Code    string `json:"success"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

// Constructor of Success Response
func NewSuccessResponse(data interface{}) Response {
	//Generate the explanation below about using the interface in Go
	//The interface is a type that defines a set of methods
	//The methods are the methods that the type must implement
	//The interface is used to define the type of the data that the function will return
	return Response{
		Success:    true,
		Data:       data,
		StatusCode: http.StatusOK,
		Timestamp:  time.Now(),
	}

}

//Constructor for the Error Response
/*
@author: NC
*/
func NewErrorResponse(statusCode int, code string, message string, detail string) Response {
	return Response{
		Success:    false,
		StatusCode: statusCode,
		Error: &ErrorDetail{
			Code:    code,
			Message: message,
			Detail:  detail,
		},
		Timestamp: time.Now(),
	}
}

// Common error codes
const (
	ErrCodeValidation   = "VALIDATION_ERROR"
	ErrCodeNotFound     = "NOT_FOUND"
	ErrCodeUnauthorized = "UNAUTHORIZED"
	ErrCodeInternal     = "INTERNAL_ERROR"
	ErrCodeBadRequest   = "BAD_REQUEST"
	// Add more as needed
)
