create table
  if not exists tag (id integer primary key, name text not null);

create table
  if not exists log_line (
    id integer primary key,
    value real not null,
    description text,
    created_at timestamp not null default current_timestamp,
    tag_id integer not null,
    foreign key (tag_id) references tag (id)
  );
