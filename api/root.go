package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/danesparza/Dashboard-service/data"
)

// SystemResponse is a response for a system request
type SystemResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Service encapsulates API service operations
type Service struct {
	DB        *data.Manager
	StartTime time.Time
}

// sendErrorResponse is used to send back an error:
func sendErrorResponse(rw http.ResponseWriter, err error, code int) {
	//	Our return value
	response := SystemResponse{
		Status:  code,
		Message: "Error: " + err.Error()}

	//	Serialize to JSON & return the response:
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(code)
	json.NewEncoder(rw).Encode(response)
}

// sendDataResponse is used to send back data:
func sendDataResponse(rw http.ResponseWriter, msg string, data interface{}) {
	//	Our return value
	response := SystemResponse{
		Status:  http.StatusOK,
		Message: msg,
		Data:    data,
	}

	//	Serialize to JSON & return the response:
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(response)
}

// getWSResponse gets a JSON formatted WebSocket event response
func getWSResponse(messageType string, item data.ConfigItem) string {
	//	Our default return value:
	retval := ""

	//	Our WebSocket return value
	response := data.WebSocketResponse{
		Data: item,
		Type: messageType}

	//	Serialize to JSON and return as a string:
	responseBytes := new(bytes.Buffer)
	if err := json.NewEncoder(responseBytes).Encode(&response); err == nil {
		retval = responseBytes.String()
	}

	return retval
}
