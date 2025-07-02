package routes

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/achmadardian/test-booking-api/handlers"
	"github.com/gorilla/mux"
)

func NewRoute(logger *slog.Logger, DB *sql.DB) http.Handler {
	// healthcheck
	healthcheckHandl := handlers.NewHealthcheckHandler()

	// router
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	// routes
	// healthcheck
	api.HandleFunc("", healthcheckHandl.Healthcheck).Methods(http.MethodGet)

	return r
}
