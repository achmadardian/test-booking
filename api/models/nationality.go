package models

import "time"

type Nationality struct {
	NationalityID   int
	NationalityName string
	NationalityCode string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}
