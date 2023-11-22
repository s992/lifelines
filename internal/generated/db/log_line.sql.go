// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: log_line.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createLogLine = `-- name: CreateLogLine :one
insert into
  log_line (tag_id, value, description)
values
  (
    ?1,
    ?2,
    ?3
  ) returning id, value, description, created_at, tag_id
`

type CreateLogLineParams struct {
	TagID       int64   `json:"tagID"`
	Value       float64 `json:"value"`
	Description *string `json:"description"`
}

func (q *Queries) CreateLogLine(ctx context.Context, arg CreateLogLineParams) (LogLine, error) {
	row := q.db.QueryRowContext(ctx, createLogLine, arg.TagID, arg.Value, arg.Description)
	var i LogLine
	err := row.Scan(
		&i.ID,
		&i.Value,
		&i.Description,
		&i.CreatedAt,
		&i.TagID,
	)
	return i, err
}

const listLogLines = `-- name: ListLogLines :many
select
  ll.id, ll.value, ll.description, ll.created_at, ll.tag_id,
  t.name as tag_name
from
  log_line ll
  inner join tag t on t.id = ll.tag_id
where
  (
    ll.tag_id = ?1
    or ?1 is null
  )
  and (
    ll.created_at >= ?2
    or ?2 is null
  )
  and (
    ll.created_at <= ?3
    or ?3 is null
  )
`

type ListLogLinesParams struct {
	TagID         sql.NullInt64 `json:"tagID"`
	StartDateTime sql.NullTime  `json:"startDateTime"`
	EndDateTime   sql.NullTime  `json:"endDateTime"`
}

type ListLogLinesRow struct {
	ID          int64     `json:"id"`
	Value       float64   `json:"value"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	TagID       int64     `json:"tagId"`
	TagName     string    `json:"tagName"`
}

func (q *Queries) ListLogLines(ctx context.Context, arg ListLogLinesParams) ([]ListLogLinesRow, error) {
	rows, err := q.db.QueryContext(ctx, listLogLines, arg.TagID, arg.StartDateTime, arg.EndDateTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListLogLinesRow{}
	for rows.Next() {
		var i ListLogLinesRow
		if err := rows.Scan(
			&i.ID,
			&i.Value,
			&i.Description,
			&i.CreatedAt,
			&i.TagID,
			&i.TagName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}