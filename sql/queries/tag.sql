-- name: GetTag :one
select
  *
from
  tag
where
  id = ?;

-- name: ListTags :many
select
  t.*
from
  tag t
  left join log_line ll on ll.tag_id = t.id
  and ll.created_at >= datetime ('now', '-7 days')
group by
  t.id
order by
  count(ll.id) desc;

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
