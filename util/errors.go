package util

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse is the Response given in case of any Error
type ErrorResponse struct {
	Code        int    `json:"code" bson:"code"`
	Message     string `json:"message" bson:"message"`
	Description string `json:"description" bson:"description"`
}

// NewErrorResponse makes an error response object
func NewErrorResponse(code int, description string) ErrorResponse {
	var err ErrorResponse
	err.Code = code
	err.Message = http.StatusText(code)
	err.Description = description
	return err
}

// Response writes a response, error or no error
func Response(w http.ResponseWriter, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(payload)
}
