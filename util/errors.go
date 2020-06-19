package util

import (
	"encoding/json"
	"net/http"
)

// Status is the Response given in case of any Error
type Status struct {
	Code        int    `json:"code" bson:"code"`
	Message     string `json:"message" bson:"message"`
	Description string `json:"description" bson:"description"`
}

// NewStatus makes an error response object
func NewStatus(code int, description string) Status {
	var status Status
	status.Code = code
	status.Message = http.StatusText(code)
	status.Description = description
	return status
}

// Response writes a response, error or no error
func Response(w http.ResponseWriter, payload interface{}, status Status) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(status.Code)

	type Resp struct {
		Payload interface{} `json:"payload" bson:"payload"`
		Status  Status      `json:"status" bson:"status"`
	}

	resp := Resp{Payload: payload, Status: status}
	json.NewEncoder(w).Encode(resp)
}
