package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/achmadardian/test-booking-api/errs"
	"github.com/achmadardian/test-booking-api/models"
	"github.com/achmadardian/test-booking-api/pagination"
	"github.com/achmadardian/test-booking-api/repositories"
	"github.com/achmadardian/test-booking-api/requests"
)

type FamilyListService struct {
	repo *repositories.FamilyListRepository
}

func NewFamilyService(repo *repositories.FamilyListRepository) *FamilyListService {
	return &FamilyListService{
		repo: repo,
	}
}

func (f *FamilyListService) GetAll(ctx context.Context, p pagination.Pagination) ([]models.FamilyList, pagination.Pagination, error) {
	families, page, err := f.repo.GetAll(ctx, p)
	if err != nil {
		return nil, pagination.Pagination{}, fmt.Errorf("FamilyListService.GetAll: failed to fetch: %w", err)
	}

	return families, page, nil
}

func (f *FamilyListService) GetByID(ctx context.Context, familyID int) (*models.FamilyList, error) {
	family, err := f.repo.GetByID(ctx, familyID)
	if err != nil {
		if errors.Is(err, errs.ErrDataNotFound) {
			return nil, errs.ErrDataNotFound
		}

		return nil, fmt.Errorf("FamilyListService.GetByID: failed to fetch: %w", err)
	}

	return family, nil
}

func (f *FamilyListService) Create(ctx context.Context, req requests.CreateFamilyRequest) (*models.FamilyList, error) {
	dob, err := time.Parse("2006-01-02", req.FLDOB)
	if err != nil {
		return nil, fmt.Errorf("FamilyListService.Create: failed to parse time: %w", err)
	}

	family := models.FamilyList{
		CSTID:      req.CstID,
		FLRelation: req.FLRelation,
		FLName:     req.FLName,
		FLDOB:      dob,
	}

	save, err := f.repo.Create(ctx, family)
	if err != nil {
		return nil, fmt.Errorf("FamilyListService.Create: failed to insert: %w", err)
	}

	return save, nil
}

func (f *FamilyListService) UpdateByID(ctx context.Context, familyID int, req requests.UpdateFamilyRequest) (*models.FamilyList, error) {
	var dob time.Time
	var err error

	_, err = f.GetByID(ctx, familyID)
	if err != nil {
		return nil, err
	}

	if req.FLDOB != "" {
		dob, err = time.Parse("2006-01-02", req.FLDOB)
		if err != nil {
			return nil, fmt.Errorf("FamilyListService.UpdateByID: failed to parse time: %w", err)
		}
	}

	family := models.FamilyList{
		CSTID:      req.CstID,
		FLRelation: req.FLRelation,
		FLName:     req.FLName,
		FLDOB:      dob,
	}

	save, err := f.repo.UpdateByID(ctx, familyID, family)
	if err != nil {
		return nil, fmt.Errorf("FamilyListService.UpdateByID: failed to update: %w", err)
	}

	return save, nil
}

func (f *FamilyListService) DeleteByID(ctx context.Context, familyID int) error {
	_, err := f.GetByID(ctx, familyID)
	if err != nil {
		return err
	}

	if err := f.repo.DeleteByID(ctx, familyID); err != nil {
		return fmt.Errorf("FamilyListService.DeleteByID: failed to delete: %w", err)
	}

	return nil
}
