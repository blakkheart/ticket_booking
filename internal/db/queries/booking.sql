-- name: GetAllBookings :many
SELECT * FROM booking;

-- name: CreateBooking :one
INSERT INTO booking (account_id, status, expires_at, paid_at)
VALUES ($1, $2, $3, $4)
RETURNING *;