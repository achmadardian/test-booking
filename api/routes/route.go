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
	// customer
	cstRepo := repositories.NewCustomerRepository(DB)
	cstSvc := services.NewCustomerService(cstRepo)
	cstHandl := handlers.NewCustomerHandler(logger, cstSvc)

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
	// customer
	customer := api.PathPrefix("/customers").Subrouter().StrictSlash(true)
	customer.HandleFunc("", cstHandl.GetAllCustomerWithRelations).Methods(http.MethodGet)
	customer.HandleFunc("/{customer_id}", cstHandl.GetCustomerByIDWithRelations).Methods(http.MethodGet)
	customer.HandleFunc("", cstHandl.CreateCustomer).Methods(http.MethodPost)
	customer.HandleFunc("/{customer_id}", cstHandl.UpdateCustomerByID).Methods(http.MethodPatch)
	customer.HandleFunc("/{customer_id}", cstHandl.DeleteCustomerByID).Methods(http.MethodDelete)

	return r
}
