package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/achmadardian/test-booking-api/models"
	"github.com/achmadardian/test-booking-api/pagination"
)

type TX interface {
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type FamilyListRepository struct {
	DB *sql.DB
}

func NewFamilyListRepository(DB *sql.DB) *FamilyListRepository {
	return &FamilyListRepository{
		DB: DB,
	}
}

func (f *FamilyListRepository) GetAll(ctx context.Context, p pagination.Pagination) ([]models.FamilyList, pagination.Pagination, error) {
	var families []models.FamilyList
	limit, offset := p.GetLimitOffset(p)

	query := `SELECT fl_id, cst_id, fl_relation, fl_name, fl_dob 
	FROM family_list 
	WHERE deleted_at IS NULL 
	ORDER BY fl_id LIMIT $1 OFFSET $2`

	rows, err := f.DB.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, pagination.Pagination{}, fmt.Errorf("FamilyListRepository.GetAll: failed to query select: %w", err)
	}

	for rows.Next() {
		var family models.FamilyList

		if err := rows.Scan(&family.FLID, &family.CSTID, &family.FLRelation, &family.FLName, &family.FLDOB); err != nil {
			return nil, pagination.Pagination{}, fmt.Errorf("FamilyListRepository.GetAll: failed to scan: %w", err)
		}

		families = append(families, family)
	}

	hasNext := false
	if len(families) > p.PageSize {
		hasNext = true
		families = families[:p.PageSize]
	}

	page := pagination.Pagination{
		Page:     p.Page,
		PageSize: p.PageSize,
		HasNext:  hasNext,
	}

	return families, page, nil
}

func (f *FamilyListRepository) GetAllByCustomerID(ctx context.Context, cstID int) ([]models.FamilyList, error) {
	var families []models.FamilyList

	query := `SELECT fl_id, cst_id, fl_relation, fl_name, fl_dob 
	FROM family_list 
	WHERE cst_id = $1 AND deleted_at IS NULL 
	ORDER BY fl_id ASC`

	rows, err := f.DB.QueryContext(ctx, query, cstID)
	if err != nil {
		return nil, fmt.Errorf("FamilyListRepository.GetAllByCustomerID: failed to query select: %w", err)
	}

	for rows.Next() {
		var family models.FamilyList

		if err := rows.Scan(&family.FLID, &family.CSTID, &family.FLRelation, &family.FLName, &family.FLDOB); err != nil {
			return nil, fmt.Errorf("FamilyListRepository.GetAllByCustomerID: failed to scan: %w", err)
		}

		families = append(families, family)
	}

	return families, nil
}

func (f *FamilyListRepository) GetByID(ctx context.Context, familyID int) (*models.FamilyList, error) {
	var family models.FamilyList

	query := `SELECT fl_id, cst_id, fl_relation, fl_name, fl_dob FROM family_list WHERE fl_id = $1 AND deleted_at IS NULL`

	row := f.DB.QueryRowContext(ctx, query, familyID)
	if err := row.Scan(&family.FLID, &family.CSTID, &family.FLRelation, &family.FLName, &family.FLDOB); err != nil {
		return nil, fmt.Errorf("FamilyListRepository.GetByID: failed to scan: %w", err)
	}

	return &family, nil
}

func (f *FamilyListRepository) Create(ctx context.Context, tx TX, fl models.FamilyList) (*models.FamilyList, error) {
	DB := tx
	if DB == nil {
		DB = f.DB
	}

	query := `INSERT INTO family_list (cst_id, fl_relation, fl_name, fl_dob)
	VALUES ($1, $2, $3, $4)
	RETURNING fl_id, cst_id, fl_relation, fl_name, fl_dob`

	row := DB.QueryRowContext(ctx, query, fl.CSTID, fl.FLRelation, fl.FLName, fl.FLDOB)
	if err := row.Scan(&fl.FLID, &fl.CSTID, &fl.FLRelation, &fl.FLName, &fl.FLDOB); err != nil {
		return nil, fmt.Errorf("FamilyListRepository.Create: failed to insert: %w", err)
	}

	return &fl, nil
}

func (f *FamilyListRepository) UpdateByID(ctx context.Context, familyID int, fl models.FamilyList) (*models.FamilyList, error) {
	query := `UPDATE family_list SET updated_at = CURRENT_TIMESTAMP`

	args := []any{}
	argIndex := 1

	if fl.CSTID != 0 {
		query += fmt.Sprintf(", cst_id = $%d", argIndex)
		args = append(args, fl.CSTID)
		argIndex++
	}

	if fl.FLRelation != "" {
		query += fmt.Sprintf(", fl_relation = $%d", argIndex)
		args = append(args, fl.FLRelation)
		argIndex++
	}

	if fl.FLName != "" {
		query += fmt.Sprintf(", fl_name = $%d", argIndex)
		args = append(args, fl.FLName)
		argIndex++
	}

	if !fl.FLDOB.IsZero() {
		query += fmt.Sprintf(", fl_dob = $%d", argIndex)
		args = append(args, fl.FLDOB)
		argIndex++
	}

	query += fmt.Sprintf(" WHERE fl_id = $%d RETURNING fl_id, cst_id, fl_relation, fl_name, fl_dob", argIndex)
	args = append(args, familyID)

	row := f.DB.QueryRowContext(ctx, query, args...)
	if err := row.Scan(&fl.FLID, &fl.CSTID, &fl.FLRelation, &fl.FLName, &fl.FLDOB); err != nil {
		return nil, fmt.Errorf("FamilyListRepository.UpdateByID: failed to update: %w", err)
	}

	return &fl, nil
}

func (f *FamilyListRepository) DeleteByID(ctx context.Context, familyID int) error {
	query := `UPDATE family_list SET deleted_at = CURRENT_TIMESTAMP WHERE fl_id = $1`

	_, err := f.DB.ExecContext(ctx, query, familyID)
	if err != nil {
		return fmt.Errorf("FamilyListRepository.DeleteByID: failed to delete: %w", err)
	}

	return nil
}
