// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const insertStatus = `-- name: InsertStatus :exec
INSERT INTO status VALUES ($1, $2)
`

type InsertStatusParams struct {
	Column1 pgtype.Int4
	Column2 pgtype.Text
}

func (q *Queries) InsertStatus(ctx context.Context, arg InsertStatusParams) error {
	_, err := q.db.Exec(ctx, insertStatus, arg.Column1, arg.Column2)
	return err
}
