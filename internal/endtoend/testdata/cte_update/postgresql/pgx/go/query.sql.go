// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const updateAttribute = `-- name: UpdateAttribute :one
with updated_attribute as (UPDATE attribute_value
    SET
        val = CASE WHEN $1::bool THEN $2 ELSE val END
    WHERE attribute_value.id = $3
    RETURNING id,attribute,val)
select updated_attribute.id, val, name
from updated_attribute
         left join attribute on updated_attribute.attribute = attribute.id
`

type UpdateAttributeParams struct {
	FilterValue pgtype.Bool
	Value       pgtype.Text
	ID          pgtype.Int8
}

type UpdateAttributeRow struct {
	ID   int64
	Val  string
	Name pgtype.Text
}

func (q *Queries) UpdateAttribute(ctx context.Context, arg UpdateAttributeParams) (UpdateAttributeRow, error) {
	row := q.db.QueryRow(ctx, updateAttribute, arg.FilterValue, arg.Value, arg.ID)
	var i UpdateAttributeRow
	err := row.Scan(&i.ID, &i.Val, &i.Name)
	return i, err
}
