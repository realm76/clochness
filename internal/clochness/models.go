// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package clochness

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Entry struct {
	ID          int32
	UserID      int32
	ProjectID   pgtype.Int4
	Description pgtype.Text
	StartDate   pgtype.Date
	EndDate     pgtype.Date
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type Organization struct {
	ID        int32
	Name      string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Project struct {
	ID             int32
	Name           string
	Description    pgtype.Text
	OrganizationID int32
	StartDate      pgtype.Date
	EndDate        pgtype.Date
	CreatedAt      pgtype.Timestamp
	UpdatedAt      pgtype.Timestamp
}

type User struct {
	ID       int32
	Email    string
	Password string
}

type UserOrganization struct {
	UserID         int32
	OrganizationID int32
	Role           pgtype.Text
	JoinedAt       pgtype.Timestamp
}
