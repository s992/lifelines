-- name: ListLogLines :many
select
  ll.*,
  t.name as tag_name
from
  log_line ll
  inner join tag t on t.id = ll.tag_id
where
  (
    ll.tag_id = sqlc.narg ('tagID')
    or sqlc.narg ('tagID') is null
  )
  and (
    ll.created_at >= sqlc.narg ('startDateTime')
    or sqlc.narg ('startDateTime') is null
  )
  and (
    ll.created_at <= sqlc.narg ('endDateTime')
    or sqlc.narg ('endDateTime') is null
  )
order by ll.created_at desc;

-- name: CreateLogLine :one
insert into
  log_line (tag_id, value, description)
values
  (
    sqlc.arg ('tagID'),
    sqlc.arg ('value'),
    sqlc.narg ('description')
  ) returning *;
