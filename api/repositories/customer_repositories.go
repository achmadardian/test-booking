package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/achmadardian/test-booking-api/errs"
	"github.com/achmadardian/test-booking-api/models"
	"github.com/achmadardian/test-booking-api/pagination"
)

type CustomerRepository struct {
	DB *sql.DB
}

func NewCustomerRepository(DB *sql.DB) *CustomerRepository {
	return &CustomerRepository{
		DB: DB,
	}
}

func (c *CustomerRepository) GetAll(ctx context.Context, p pagination.Pagination) ([]models.Customer, pagination.Pagination, error) {
	var customers []models.Customer
	limit, offset := p.GetLimitOffset(p)

	query := `
	SELECT
		c.cst_id,
		c.nationality_id,
		c.cst_name,
		c.cst_dob,
		c.cst_phone_num,
		c.cst_email,
		n.nationality_id,
		n.nationality_name,
		n.nationality_code
	FROM customer c
	JOIN nationality n on c.nationality_id=n.nationality_id
	AND c.deleted_at IS NULL
	ORDER BY c.cst_id DESC
	LIMIT $1 OFFSET $2`

	rows, err := c.DB.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, pagination.Pagination{}, fmt.Errorf("CustomerRepository.GetAll: failed to select query: %w", err)
	}

	for rows.Next() {
		var cst models.Customer

		scan := rows.Scan(&cst.CstID, &cst.NationalityId, &cst.CstName, &cst.CstDOB, &cst.CstPhoneNum, &cst.CstEmail,
			&cst.Nationality.NationalityID, &cst.Nationality.NationalityName, &cst.Nationality.NationalityCode)

		if scan != nil {
			return nil, pagination.Pagination{}, fmt.Errorf("CustomerRepository.GetAll: failed to scan: %w", err)
		}

		customers = append(customers, cst)
	}

	hasNext := false
	if len(customers) > p.PageSize {
		hasNext = true
		customers = customers[:p.PageSize]
	}

	page := pagination.Pagination{
		Page:     p.Page,
		PageSize: p.PageSize,
		HasNext:  hasNext,
	}

	return customers, page, nil
}

func (c *CustomerRepository) GetByID(ctx context.Context, cstID int) (*models.Customer, error) {
	var cst models.Customer

	query := `
	SELECT
		c.cst_id,
		c.nationality_id,
		c.cst_name,
		c.cst_dob,
		c.cst_phone_num,
		c.cst_email,
		n.nationality_id,
		n.nationality_name,
		n.nationality_code
	FROM customer c
	JOIN nationality n on c.nationality_id=n.nationality_id
	WHERE c.cst_id = $1
	AND c.deleted_at IS NULL`

	row := c.DB.QueryRowContext(ctx, query, cstID)
	err := row.Scan(&cst.CstID, &cst.NationalityId, &cst.CstName, &cst.CstDOB, &cst.CstPhoneNum, &cst.CstEmail,
		&cst.Nationality.NationalityID, &cst.Nationality.NationalityName, &cst.Nationality.NationalityCode)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrDataNotFound
		}

		return nil, fmt.Errorf("CustomerRepository.GetByID: failed to scan: %w", err)
	}

	return &cst, nil
}

func (c *CustomerRepository) Create(ctx context.Context, tx TX, cst models.Customer) (*models.Customer, error) {
	DB := tx
	if DB == nil {
		DB = c.DB
	}

	query := `INSERT INTO customer (nationality_id, cst_name, cst_dob, cst_phone_num, cst_email)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING cst_id, nationality_id, cst_name, cst_dob, cst_phone_num, cst_email`

	row := c.DB.QueryRowContext(ctx, query, cst.NationalityId, cst.CstName, cst.CstDOB, cst.CstPhoneNum, cst.CstEmail)
	err := row.Scan(&cst.CstID, &cst.NationalityId, &cst.CstName, &cst.CstDOB, &cst.CstPhoneNum, &cst.CstEmail)
	if err != nil {
		return nil, fmt.Errorf("CustomerRepository.Create: failed to insert: %w", err)
	}

	return &cst, nil
}

func (c *CustomerRepository) UpdateByID(ctx context.Context, cstID int, cst models.Customer) (*models.Customer, error) {
	query := `UPDATE customer SET updated_at = CURRENT_TIMESTAMP`

	args := []any{}
	argIndex := 1

	if cst.CstName != "" {
		query += fmt.Sprintf(" , cst_name = $%d", argIndex)
		args = append(args, cst.CstName)
		argIndex++
	}

	if cst.NationalityId != 0 {
		query += fmt.Sprintf(", nationality_id = $%d", argIndex)
		args = append(args, cst.NationalityId)
		argIndex++
	}

	if !cst.CstDOB.IsZero() {
		query += fmt.Sprintf(", cst_dob = $%d", argIndex)
		args = append(args, cst.CstDOB)
		argIndex++
	}

	if cst.CstPhoneNum != "" {
		query += fmt.Sprintf(", cst_phone_num = $%d", argIndex)
		args = append(args, cst.CstPhoneNum)
		argIndex++
	}

	if cst.CstEmail != "" {
		query += fmt.Sprintf(", cst_email = $%d", argIndex)
		args = append(args, cst.CstEmail)
		argIndex++
	}

	query += fmt.Sprintf(" WHERE cst_id = $%d RETURNING cst_id, nationality_id, cst_name, cst_dob, cst_phone_num, cst_email", argIndex)
	args = append(args, cstID)

	row := c.DB.QueryRowContext(ctx, query, args...)
	err := row.Scan(&cst.CstID, &cst.NationalityId, &cst.CstName, &cst.CstDOB, &cst.CstPhoneNum, &cst.CstEmail)
	if err != nil {
		return nil, fmt.Errorf("CustomerRepository.UpdateByID: failed to update: %w", err)
	}

	return &cst, nil
}

func (c *CustomerRepository) DeleteByID(ctx context.Context, cstID int) error {
	query := `UPDATE customer SET deleted_at = CURRENT_TIMESTAMP WHERE cst_id = $1`
	_, err := c.DB.ExecContext(ctx, query, cstID)
	if err != nil {
		return fmt.Errorf("CustomerRepository.Delete: failed to delete by id: %w", err)
	}

	return nil
}
