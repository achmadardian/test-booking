package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/achmadardian/test-booking-api/errs"
	"github.com/achmadardian/test-booking-api/models"
	"github.com/achmadardian/test-booking-api/pagination"
	"github.com/achmadardian/test-booking-api/repositories"
	"github.com/achmadardian/test-booking-api/requests"
)

type CustomerService struct {
	repo  *repositories.CustomerRepository
	fRepo *repositories.FamilyListRepository
	DB    *sql.DB
}

func NewCustomerService(repo *repositories.CustomerRepository, f *repositories.FamilyListRepository, DB *sql.DB) *CustomerService {
	return &CustomerService{
		repo:  repo,
		fRepo: f,
		DB:    DB,
	}
}

func (c *CustomerService) GetAll(ctx context.Context, p pagination.Pagination) ([]models.Customer, pagination.Pagination, error) {
	customers, page, err := c.repo.GetAll(ctx, p)
	if err != nil {
		return nil, pagination.Pagination{}, fmt.Errorf("CustomerService.GetAll: failed to fetch: %w", err)
	}

	return customers, page, nil
}

func (c *CustomerService) GetAllFamiliesByCustomerID(ctx context.Context, cstID int) ([]models.FamilyList, error) {
	families, err := c.fRepo.GetAllByCustomerID(ctx, cstID)
	if err != nil {
		return nil, fmt.Errorf("CustomerService.GetAllFamiliesByCustomerID: failed to fetch: %w", err)
	}

	return families, nil
}

func (c *CustomerService) GetByID(ctx context.Context, cstID int) (*models.Customer, error) {
	customer, err := c.repo.GetByID(ctx, cstID)
	if err != nil {
		if errors.Is(err, errs.ErrDataNotFound) {
			return nil, errs.ErrDataNotFound
		}

		return nil, fmt.Errorf("CustomerService.GetByID: failed to fetch: %w", err)
	}

	return customer, nil
}

func (c *CustomerService) Create(ctx context.Context, req requests.CreateCustomerRequest) (*models.Customer, []models.FamilyList, error) {
	tx, err := c.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("CustomerService.Create: failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	cstDOB, err := time.Parse("2006-01-02", req.CstDOB)
	if err != nil {
		return nil, nil, fmt.Errorf("CustomerService.Create: failed to convert date of birth customer: %w", err)
	}

	cst := models.Customer{
		NationalityId: req.NationalityID,
		CstName:       req.CstName,
		CstDOB:        cstDOB,
		CstPhoneNum:   req.CstPhoneNum,
		CstEmail:      req.CstEmail,
	}

	saveCst, err := c.repo.Create(ctx, tx, cst)
	if err != nil {
		return nil, nil, fmt.Errorf("CustomerService.Create: failed to insert: %w", err)
	}

	var families []models.FamilyList
	for _, f := range req.Families {
		fmlDOB, err := time.Parse("2006-01-02", f.FLDOB)
		if err != nil {
			return nil, nil, fmt.Errorf("CustomerService.Create: failed to convert date of birth family: %w", err)
		}

		fml := models.FamilyList{
			CSTID:      saveCst.CstID,
			FLRelation: f.FLRelation,
			FLName:     f.FLName,
			FLDOB:      fmlDOB,
		}

		saveFml, err := c.fRepo.Create(ctx, tx, fml)
		if err != nil {
			return nil, nil, fmt.Errorf("CustomerService.Create: failed to insert family: %w", err)
		}

		families = append(families, *saveFml)
	}

	tx.Commit()

	return saveCst, families, nil
}

func (c *CustomerService) UpdateByID(ctx context.Context, cstID int, req requests.UpdateCustomerRequest) (*models.Customer, error) {
	var dob time.Time
	var err error

	_, err = c.GetByID(ctx, cstID)
	if err != nil {
		return nil, err
	}

	if req.CstDOB != "" {
		dob, err = time.Parse("2006-01-02", req.CstDOB)
		if err != nil {
			return nil, fmt.Errorf("CustomerService.UpdateByID: failed to convert date of birth: %w", err)
		}
	}

	cst := models.Customer{
		NationalityId: req.NationalityID,
		CstName:       req.CstName,
		CstDOB:        dob,
		CstPhoneNum:   req.CstPhoneNum,
		CstEmail:      req.CstEmail,
	}

	_, err = c.repo.UpdateByID(ctx, cstID, cst)
	if err != nil {
		return nil, fmt.Errorf("CustomerService.UpdateByID: failed to update: %w", err)
	}

	updatedCst, err := c.GetByID(ctx, cstID)
	if err != nil {
		return nil, err
	}

	return updatedCst, nil
}

func (c *CustomerService) DeleteByID(ctx context.Context, cstID int) error {
	_, err := c.GetByID(ctx, cstID)
	if err != nil {
		return err
	}

	if err = c.repo.DeleteByID(ctx, cstID); err != nil {
		return fmt.Errorf("CustomerService.DeleteByID: failed to delete: %w", err)
	}

	return nil
}
