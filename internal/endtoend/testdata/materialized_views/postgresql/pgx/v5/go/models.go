// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package querytest

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Author struct {
	ID     int64
	Name   string
	Bio    pgtype.Text
	Gender pgtype.Int4
}

type AuthorsName struct {
	Name string
}
