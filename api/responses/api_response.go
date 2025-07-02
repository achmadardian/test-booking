package responses

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type ApiResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}

const (
	MsgOk                  = "success"
	MsgCreated             = "created"
	MsgUpdated             = "updated"
	MsgDeleted             = "deleted"
	MsgBadRequest          = "bad request"
	MsgNotFound            = "not found"
	MsgUnprocessableEntity = "unprocessable entity"
	MsgInternalServerError = "internal server error"
)

var logger *slog.Logger

func SetLogger(l *slog.Logger) {
	logger = l
}

func writeJSON(w http.ResponseWriter, statusCode int, res ApiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		logger.Error("failed to encode JSON",
			slog.Any("error", err),
		)

		http.Error(w, MsgInternalServerError, http.StatusInternalServerError)
	}
}

func Ok(w http.ResponseWriter, data any, message ...string) {
	msg := MsgOk
	if len(message) > 0 {
		msg = message[0]
	}

	res := ApiResponse{
		Message: msg,
		Data:    data,
	}

	writeJSON(w, http.StatusOK, res)
}

func Created(w http.ResponseWriter, data any, message ...string) {
	msg := MsgCreated
	if len(message) > 0 {
		msg = message[0]
	}

	res := ApiResponse{
		Message: msg,
		Data:    data,
	}

	writeJSON(w, http.StatusCreated, res)
}

func Updated(w http.ResponseWriter, data any, message ...string) {
	msg := MsgUpdated
	if len(message) > 0 {
		msg = message[0]
	}

	res := ApiResponse{
		Message: msg,
		Data:    data,
	}

	writeJSON(w, http.StatusOK, res)
}

func BadRequest(w http.ResponseWriter, message ...string) {
	msg := MsgBadRequest
	if len(message) > 0 {
		msg = message[0]
	}

	res := ApiResponse{
		Message: msg,
	}

	writeJSON(w, http.StatusBadRequest, res)
}

func NotFound(w http.ResponseWriter, message ...string) {
	msg := MsgNotFound
	if len(message) > 0 {
		msg = message[0]
	}

	res := ApiResponse{
		Message: msg,
	}

	writeJSON(w, http.StatusNotFound, res)
}

func UnprocessableEntity(w http.ResponseWriter, errors any, message ...string) {
	msg := MsgUnprocessableEntity
	if len(message) > 0 {
		msg = message[0]
	}

	res := ApiResponse{
		Message: msg,
		Errors:  errors,
	}

	writeJSON(w, http.StatusUnprocessableEntity, res)
}

func Deleted(w http.ResponseWriter, message ...string) {
	msg := MsgDeleted
	if len(message) > 0 {
		msg = message[0]
	}

	res := ApiResponse{
		Message: msg,
	}

	writeJSON(w, http.StatusOK, res)
}

func InternalServerError(w http.ResponseWriter) {
	res := ApiResponse{
		Message: MsgInternalServerError,
	}

	writeJSON(w, http.StatusInternalServerError, res)
}
