package models

import "time"

type FamilyList struct {
	FLID       int
	CSTID      int
	FLRelation string
	FLName     string
	FLDOB      time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}
