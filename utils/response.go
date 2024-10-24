package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    interface{}    `json:"data,omitempty"`
}

// global response
// status for is ok or not ok or whatever
// message given to client either success or not
// status code for response http
// data is optional
func JSONResponse(w http.ResponseWriter, r *http.Request, status string,message string, data interface{},statusCode int) {
	response := Response {
		Status:  status,
		Message: message,
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}