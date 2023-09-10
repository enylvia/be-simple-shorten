package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Status  uint16      `json:"status"`
	Data    interface{} `json:"data"`
}

func ReturnJSON(w http.ResponseWriter, response Response) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func SuccessResponse(payload interface{}) Response {
	return Response{
		Message: "Request successfully!",
		Status:  200,
		Data:    payload,
	}
}
func CreatedResponse(payload interface{}) Response {
	return Response{
		Message: "Request successfully!",
		Status:  201,
		Data:    payload,
	}
}
func InternalServerErrorResponse(payload interface{}) Response {
	return Response{
		Message: "Something went wrong!",
		Status:  500,
		Data:    payload,
	}
}
func BadRequestResponse(message interface{}) Response {
	return Response{
		Message: "Bad request!",
		Status:  403,
		Data:    message,
	}
}
