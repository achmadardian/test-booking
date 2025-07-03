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

type CustomerService struct {
	repo *repositories.CustomerRepository
}

func NewCustomerService(repo *repositories.CustomerRepository) *CustomerService {
	return &CustomerService{
		repo: repo,
	}
}

func (c *CustomerService) GetAll(ctx context.Context, p pagination.Pagination) ([]models.Customer, pagination.Pagination, error) {
	customers, page, err := c.repo.GetAll(ctx, p)
	if err != nil {
		return nil, pagination.Pagination{}, fmt.Errorf("CustomerService.GetAll: failed to fetch: %w", err)
	}

	return customers, page, nil
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

func (c *CustomerService) Create(ctx context.Context, req requests.CreateCustomerRequest) (*models.Customer, error) {
	time, err := time.Parse("2006-01-02", req.CstDOB)
	if err != nil {
		return nil, fmt.Errorf("CustomerService.Create: failed to convert date of birth: %w", err)
	}

	cst := models.Customer{
		NationalityId: req.NationalityID,
		CstName:       req.CstName,
		CstDOB:        time,
		CstPhoneNum:   req.CstPhoneNum,
		CstEmail:      req.CstEmail,
	}

	saveCst, err := c.repo.Create(ctx, cst)
	if err != nil {
		return nil, fmt.Errorf("CustomerService.Create: failed to insert: %w", err)
	}

	return saveCst, nil
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
