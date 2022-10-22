-- name: Get :one
select * from person where id = ?;

-- name: List :many
select id, first_name, last_name 
from person
order by first_name;

-- name: Create :execresult
insert into person (
    first_name, last_name, created_at
) 
values(
    ?, ?, now()
);

-- name: Delete :exec
delete from person 
where id = ?;

-- name: Update :exec
update person 
set first_name = ?, last_name = ?, updated_at = now() 
where id = ?;

-- name: Search :many
select id, first_name, last_name from person 
where first_name like ? or last_name like ?;
