// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: tag.sql

package db

import (
	"context"
)

const createTag = `-- name: CreateTag :one
insert into
  tag (name)
values
  (?) returning id, name
`

func (q *Queries) CreateTag(ctx context.Context, name string) (Tag, error) {
	row := q.db.QueryRowContext(ctx, createTag, name)
	var i Tag
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getTag = `-- name: GetTag :one
select
  id, name
from
  tag
where
  id = ?
`

func (q *Queries) GetTag(ctx context.Context, id int64) (Tag, error) {
	row := q.db.QueryRowContext(ctx, getTag, id)
	var i Tag
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const listTags = `-- name: ListTags :many
select
  t.id, t.name
from
  tag t
  left join log_line ll on ll.tag_id = t.id
  and ll.created_at >= datetime ('now', '-7 days')
group by
  t.id
order by
  count(ll.id) desc
`

func (q *Queries) ListTags(ctx context.Context) ([]Tag, error) {
	rows, err := q.db.QueryContext(ctx, listTags)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Tag{}
	for rows.Next() {
		var i Tag
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
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

const searchTags = `-- name: SearchTags :many
select
  id, name
from
  tag
where
  lower(name) like '%' || lower(?1) || '%'
`

func (q *Queries) SearchTags(ctx context.Context, tagSearch string) ([]Tag, error) {
	rows, err := q.db.QueryContext(ctx, searchTags, tagSearch)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Tag{}
	for rows.Next() {
		var i Tag
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
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
