-- name: GetAllEvents :many
SELECT * FROM event;

-- name: CreateEvent :one
INSERT INTO event (title, description, location, starts_at, ends_at, status)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetEvent :one
SELECT * FROM event
WHERE id = $1;