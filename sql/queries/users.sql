-- name: CreateUser :one
INSERT INTO users (
    id,
    created_at,
    updated_at,
    email,
    hashed_password
)
values (
    gen_random_uuid(),
    now(),
    now(),
    $1,
    $2
)
returning *;


-- name: UpdateUser :one
UPDATE users SET email = $1, hashed_password = $2, updated_at = NOW()
WHERE id = $3
RETURNING *;
