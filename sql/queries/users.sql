-- name: CreateUser :one
INSERT INTO users (
    id,
    created_at,
    updated_at,
    email
)
values (
    gen_random_uuid(),
    now(),
    now(),
    $1
)
returning *;
