package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/achmadardian/test-booking-api/models"
	"github.com/achmadardian/test-booking-api/pagination"
)

type NationalityRepository struct {
	DB *sql.DB
}

func NewNationalityRepository(DB *sql.DB) *NationalityRepository {
	return &NationalityRepository{
		DB: DB,
	}
}

func (n *NationalityRepository) GetAll(ctx context.Context, p pagination.Pagination) ([]models.Nationality, pagination.Pagination, error) {
	var Nationalities []models.Nationality
	limit, offset := p.GetLimitOffset(p)

	query := `SELECT nationality_id, nationality_name, nationality_code FROM nationality ORDER BY nationality_id DESC LIMIT $1 OFFSET $2`

	rows, err := n.DB.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, pagination.Pagination{}, fmt.Errorf("NationalityRepository.GetAll: failed to query select nationality: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var nationality models.Nationality

		if err := rows.Scan(&nationality.NationalityID, &nationality.NationalityName, &nationality.NationalityCode); err != nil {
			return nil, pagination.Pagination{}, fmt.Errorf("NationalityRepository.GetAll: failed to scan nationality: %w", err)
		}

		Nationalities = append(Nationalities, nationality)
	}

	hasNext := false
	if len(Nationalities) > p.PageSize {
		hasNext = true
		Nationalities = Nationalities[:p.PageSize]
	}

	page := pagination.Pagination{
		Page:     p.Page,
		PageSize: p.PageSize,
		HasNext:  hasNext,
	}

	return Nationalities, page, nil
}
