-- name: GetTag :one
select
  *
from
  tag
where
  id = ?;

-- name: ListTags :many
select
  *
from
  tag;

-- name: SearchTags :many
select
  *
from
  tag
where
  lower(name) like '%' || lower(sqlc.arg ('tag_search')) || '%';

-- name: CreateTag :one
insert into
  tag (name)
values
  (?) returning *;
