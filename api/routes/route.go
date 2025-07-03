package routes

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/achmadardian/test-booking-api/handlers"
	"github.com/achmadardian/test-booking-api/middlewares"
	"github.com/achmadardian/test-booking-api/repositories"
	"github.com/achmadardian/test-booking-api/services"
	"github.com/gorilla/mux"
)

func NewRoute(logger *slog.Logger, DB *sql.DB) http.Handler {
	// healthcheck
	healthcheckHandl := handlers.NewHealthcheckHandler()
	// nationality
	nationRepo := repositories.NewNationalityRepository(DB)
	nationSvc := services.NewNationalityService(nationRepo)
	nationHandl := handlers.NewNationalityHandler(logger, nationSvc)

	// middleware
	middlewares.SetLogger(logger)

	// router
	r := mux.NewRouter()
	r.Use(middlewares.Logger)

	api := r.PathPrefix("/api").Subrouter()

	// routes
	// healthcheck
	api.HandleFunc("", healthcheckHandl.Healthcheck).Methods(http.MethodGet)
	// nationality
	nationality := api.PathPrefix("/nationalities").Subrouter()
	nationality.HandleFunc("", nationHandl.GetAllNationality).Methods(http.MethodGet)

	return r
}
