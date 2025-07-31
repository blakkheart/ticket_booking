-- name: GetAllEvents :many
SELECT * FROM event;

-- name: CreateEvent :one
INSERT INTO event (title, description, date, location)
VALUES ($1, $2, $3, $4)
RETURNING *;