package models

import "time"

type Customer struct {
	CstID         int
	NationalityId int
	CstName       string
	CstDOB        time.Time
	CstPhoneNum   string
	CstEmail      string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	Nationality   Nationality
}
