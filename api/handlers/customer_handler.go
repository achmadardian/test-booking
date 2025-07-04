package handlers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/achmadardian/test-booking-api/errs"
	"github.com/achmadardian/test-booking-api/pagination"
	"github.com/achmadardian/test-booking-api/requests"
	"github.com/achmadardian/test-booking-api/responses"
	"github.com/achmadardian/test-booking-api/services"
	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	logger  *slog.Logger
	service *services.CustomerService
}

func NewCustomerHandler(l *slog.Logger, s *services.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		logger:  l,
		service: s,
	}
}

func (c *CustomerHandler) GetAllCustomerWithRelations(w http.ResponseWriter, r *http.Request) {
	pagination := pagination.GetPagination(r)

	customers, page, err := c.service.GetAll(r.Context(), pagination)
	if err != nil {
		c.logger.Error("failed to fetch all customer",
			slog.String("handler", "CustomerHandler.GetAllCustomerWithRelations"),
			slog.Any("error", err),
		)

		responses.InternalServerError(w)
		return
	}

	res := responses.CustomerResponse{}
	mapRes := res.Map(customers)

	responses.OkPage(w, mapRes, page)
}

func (c *CustomerHandler) GetAllFamiliesByCustomerID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	if param["customer_id"] == "" {
		responses.BadRequest(w)
		return
	}

	cstID, err := strconv.Atoi(param["customer_id"])
	if err != nil {
		responses.BadRequest(w, "invalid customer_id param")
		return
	}

	families, err := c.service.GetAllFamiliesByCustomerID(r.Context(), cstID)
	if err != nil {
		c.logger.Error("failed to fetch all families by customer id",
			slog.String("handler", "CustomerHandler.GetAllFamiliesByCustomerID"),
			slog.Any("error", err),
		)

		responses.InternalServerError(w)
		return
	}

	res := responses.FamilyListResponse{}
	mapRes := res.Map(families)

	responses.Ok(w, mapRes)

}

func (c *CustomerHandler) GetCustomerByIDWithRelations(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	if param["customer_id"] == "" {
		responses.BadRequest(w)
		return
	}

	cstID, err := strconv.Atoi(param["customer_id"])
	if err != nil {
		responses.BadRequest(w, "invalid customer_id param")
		return
	}

	customer, err := c.service.GetByID(r.Context(), cstID)
	if err != nil {
		if errors.Is(err, errs.ErrDataNotFound) {
			responses.NotFound(w)
			return
		}

		c.logger.Error("failed to fetch customer by id",
			slog.String("handler", "CustomerHandler.GetCustomerByIDWithRelations"),
			slog.Any("error", err),
		)

		responses.InternalServerError(w)
		return
	}

	res := responses.CustomerResponse{}
	resMap := res.MapRow(customer)

	responses.Ok(w, resMap)
}

func (c *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var req requests.CreateCustomerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.BadRequest(w, "invalid JSON format")
		return
	}

	customer, families, err := c.service.Create(r.Context(), req)
	if err != nil {
		c.logger.Error("failed to create customer",
			slog.String("handler", "CustomerHandler.CreateCustomer"),
			slog.Any("body", req),
			slog.Any("error", err),
		)

		responses.InternalServerError(w)
		return
	}

	var familiesRes []responses.FamilyListResponse
	for _, f := range families {
		family := responses.FamilyListResponse{
			FLID:       f.FLID,
			CSTID:      f.CSTID,
			FLRelation: f.FLRelation,
			FLName:     f.FLName,
			FLDOB:      f.FLDOB.Format("2006-01-02"),
		}

		familiesRes = append(familiesRes, family)
	}

	resMap := responses.CreateCustomerResponse{
		CstID:       customer.CstID,
		CstName:     customer.CstName,
		CstDOB:      customer.CstDOB.Format("2006-01-02"),
		CstPhoneNum: customer.CstPhoneNum,
		CstEmail:    customer.CstEmail,
		Nationality: responses.NationalityResponse{
			NationalityID:   customer.Nationality.NationalityID,
			NationalityName: customer.Nationality.NationalityName,
			NationalityCode: customer.Nationality.NationalityCode,
		},
		Families: familiesRes,
	}

	responses.Created(w, resMap)
}

func (c *CustomerHandler) UpdateCustomerByID(w http.ResponseWriter, r *http.Request) {
	var req requests.UpdateCustomerRequest

	param := mux.Vars(r)
	if param["customer_id"] == "" {
		responses.BadRequest(w, "missing customer_id param")
		return
	}

	cstID, err := strconv.Atoi(param["customer_id"])
	if err != nil {
		responses.BadRequest(w, "invalid customer_id param")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.BadRequest(w, "invalid JSON format")
		return
	}

	update, err := c.service.UpdateByID(r.Context(), cstID, req)
	if err != nil {
		if errors.Is(err, errs.ErrDataNotFound) {
			responses.NotFound(w)
			return
		}

		c.logger.Error("failed to update customer by id",
			slog.String("handler", "CustomerHandler.UpdateCustomerByID"),
			slog.Any("error", err),
		)

		responses.InternalServerError(w)
		return
	}

	res := responses.CustomerResponse{}
	mapRes := res.MapRow(update)

	responses.Updated(w, mapRes)
}

func (c *CustomerHandler) DeleteCustomerByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	if param["customer_id"] == "" {
		responses.BadRequest(w, "missing customer_id param")
		return
	}

	cstID, err := strconv.Atoi(param["customer_id"])
	if err != nil {
		responses.BadRequest(w, "invalid customer_id param")
		return
	}

	if err := c.service.DeleteByID(r.Context(), cstID); err != nil {
		if errors.Is(err, errs.ErrDataNotFound) {
			responses.NotFound(w)
			return
		}

		c.logger.Error("failed to delete customer by id",
			slog.String("handler", "CustomerHandler.DeleteCustomerByID"),
			slog.Any("error", err),
		)

		responses.InternalServerError(w)
		return
	}

	responses.Deleted(w)
}
