package handlers

import (
	"net/http"

	"github.com/achmadardian/test-booking-api/responses"
)

type HealthcheckHandler struct{}

func NewHealthcheckHandler() *HealthcheckHandler {
	return &HealthcheckHandler{}
}

func (h *HealthcheckHandler) Healthcheck(w http.ResponseWriter, r *http.Request) {
	responses.Ok(w, nil, "app is healthy")
}
