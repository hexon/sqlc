// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package querytest

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Job struct {
	TaskName string
	LastRun  pgtype.Timestamptz
}
