-- name: GetAllBookingItems :many
SELECT * FROM booking_items;

-- name: CreateBookingItem :one
INSERT INTO booking_items (booking_id, ticket_type_id, quantity, price_at_booking)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetBookingItem :one
SELECT sqlc.embed(booking_items), sqlc.embed(booking), sqlc.embed(ticket_type)
FROM booking_items
JOIN booking ON booking.id == booking_items.booking_id
JOIN ticket_type ON ticket_type.id == booking_items.ticket_type_id
WHERE booking_items.id = $1;
