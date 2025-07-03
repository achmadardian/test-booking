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

type FamilyListHandler struct {
	logger  *slog.Logger
	service *services.FamilyListService
}

func NewFamilyListHandler(l *slog.Logger, s *services.FamilyListService) *FamilyListHandler {
	return &FamilyListHandler{
		logger:  l,
		service: s,
	}
}

func (f *FamilyListHandler) GetAllFamilies(w http.ResponseWriter, r *http.Request) {
	pagination := pagination.GetPagination(r)

	families, page, err := f.service.GetAll(r.Context(), pagination)
	if err != nil {
		f.logger.Error("failed to fetch all families",
			slog.String("handler", "FamilyListHandler.GetAllFamilies"),
			slog.Any("pagination", pagination),
			slog.Any("error", err),
		)

		responses.InternalServerError(w)
		return
	}

	res := responses.FamilyListResponse{}
	resMap := res.Map(families)

	responses.OkPage(w, resMap, page)
}

func (f *FamilyListHandler) GetFamilyByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	if param["family_id"] == "" {
		responses.BadRequest(w, "missing family_id param")
		return
	}

	familyID, err := strconv.Atoi(param["family_id"])
	if err != nil {
		responses.BadRequest(w, "invalid family_id param")
		return
	}

	family, err := f.service.GetByID(r.Context(), familyID)
	if err != nil {
		if errors.Is(err, errs.ErrDataNotFound) {
			responses.NotFound(w)
			return
		}

		f.logger.Error("failed to fetch family by id",
			slog.String("handler", "FamilyListHandler.GetFamilyByID"),
			slog.Int("param", familyID),
			slog.Any("error", err),
		)

		responses.InternalServerError(w)
		return
	}

	res := responses.FamilyListResponse{}
	resMap := res.MapRow(family)

	responses.Ok(w, resMap)
}

func (f *FamilyListHandler) CreateFamily(w http.ResponseWriter, r *http.Request) {
	var req requests.CreateFamilyRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.BadRequest(w, "invalid JSON format")
		return
	}

	family, err := f.service.Create(r.Context(), req)
	if err != nil {
		f.logger.Error("failed to create family",
			slog.String("handler", "FamilyListHandler.CreateFamily"),
			slog.Any("body", req),
			slog.Any("error", err),
		)

		responses.InternalServerError(w)
		return
	}

	res := responses.FamilyListResponse{}
	resMap := res.MapRow(family)

	responses.Created(w, resMap)
}

func (f *FamilyListHandler) UpdateFamilyByID(w http.ResponseWriter, r *http.Request) {
	var req requests.UpdateFamilyRequest

	param := mux.Vars(r)
	if param["family_id"] == "" {
		responses.BadRequest(w, "missing family_id param")
		return
	}

	familyID, err := strconv.Atoi(param["family_id"])
	if err != nil {
		responses.BadRequest(w, "invalid family_id param")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.BadRequest(w, "invalid JSON format")
		return
	}

	family, err := f.service.UpdateByID(r.Context(), familyID, req)
	if err != nil {
		if errors.Is(err, errs.ErrDataNotFound) {
			responses.NotFound(w)
			return
		}

		f.logger.Error("failed to update family by id",
			slog.String("handler", "FamilyListHandler.UpdateFamilyByID"),
			slog.Int("param", familyID),
			slog.Any("body", req),
			slog.Any("error", err),
		)

		responses.InternalServerError(w)
		return
	}

	res := responses.FamilyListResponse{}
	resMap := res.MapRow(family)

	responses.Updated(w, resMap)
}

func (f *FamilyListHandler) DeleteFamilyByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	if param["family_id"] == "" {
		responses.BadRequest(w, "missing family_id param")
		return
	}

	familyID, err := strconv.Atoi(param["family_id"])
	if err != nil {
		responses.BadRequest(w, "invalid family_id param")
		return
	}

	if err := f.service.DeleteByID(r.Context(), familyID); err != nil {
		if errors.Is(err, errs.ErrDataNotFound) {
			responses.NotFound(w)
			return
		}

		f.logger.Error("failed to delete family by id",
			slog.String("handler", "FamilyListHandler.DeleteFamilyByID"),
			slog.Int("param", familyID),
			slog.Any("error", err),
		)

		responses.InternalServerError(w)
		return
	}

	responses.Deleted(w)
}
