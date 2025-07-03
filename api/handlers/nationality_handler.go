package handlers

import (
	"log/slog"
	"net/http"

	"github.com/achmadardian/test-booking-api/pagination"
	"github.com/achmadardian/test-booking-api/responses"
	"github.com/achmadardian/test-booking-api/services"
)

type NationalityHandler struct {
	logger  *slog.Logger
	service *services.NationalityService
}

func NewNationalityHandler(l *slog.Logger, s *services.NationalityService) *NationalityHandler {
	return &NationalityHandler{
		logger:  l,
		service: s,
	}
}

func (n *NationalityHandler) GetAllNationality(w http.ResponseWriter, r *http.Request) {
	pagination := pagination.GetPagination(r)

	nationalities, page, err := n.service.GetAll(r.Context(), pagination)
	if err != nil {
		n.logger.Error("failed to fetch all nationality",
			slog.String("handler", "NationalityHandler.GetAllNationality"),
			slog.Any("error", err),
		)

		responses.InternalServerError(w)
		return
	}

	res := responses.NationalityResponse{}
	mapRes := res.Map(nationalities)

	responses.OkPage(w, mapRes, page)
}
