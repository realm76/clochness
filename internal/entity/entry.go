package entity

import "time"

type Entry struct {
	ID          int32
	UserID      int32
	ProjectID   int32
	Description string
	StartDate   time.Time
	EndDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
