-- name: GetAllBookings :many
SELECT * FROM booking;

-- name: CreateBooking :one
INSERT INTO booking (account_id, event_id, quantity)
VALUES ($1, $2, $3)
RETURNING *;