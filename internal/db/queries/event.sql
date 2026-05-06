-- name: GetAllEvents :many
SELECT * FROM event;

-- name: CreateEvent :one
INSERT INTO event (id, title, description, location, starts_at, ends_at, status)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetEvent :one
SELECT * FROM event
WHERE id = $1;

-- name: GetEvents :many
SELECT *
FROM event
WHERE
    ($1::text IS NULL OR location = $1)
AND ($2::timestamptz IS NULL OR starts_at >= $2)
AND ($3::timestamptz IS NULL OR ends_at IS NULL OR ends_at <= $3)
AND ($4::text IS NULL OR status = $4)
ORDER BY starts_at DESC
LIMIT $5 OFFSET $6;

--name: UpdateEventByIdPatch :one
UPDATE event
SET
    title = COALESCE($2, title),
    description = COALESCE($3, description),
    location = COALESCE($4, location),
    starts_at = COALESCE($5, starts_at),
    ends_at = COALESCE($6, ends_at),
    status = COALESCE($7, status)
WHERE id = $1
RETURNING *;