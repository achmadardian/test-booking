package services

import (
	"context"
	"fmt"

	"github.com/achmadardian/test-booking-api/models"
	"github.com/achmadardian/test-booking-api/pagination"
	"github.com/achmadardian/test-booking-api/repositories"
)

type NationalityService struct {
	repo *repositories.NationalityRepository
}

func NewNationalityService(r *repositories.NationalityRepository) *NationalityService {
	return &NationalityService{
		repo: r,
	}
}

func (n *NationalityService) GetAll(ctx context.Context, p pagination.Pagination) ([]models.Nationality, pagination.Pagination, error) {
	nationalities, page, err := n.repo.GetAll(ctx, p)
	if err != nil {
		return nil, pagination.Pagination{}, fmt.Errorf("NationalityService.GetAll: failed to fetch: %w", err)
	}

	return nationalities, page, nil
}
