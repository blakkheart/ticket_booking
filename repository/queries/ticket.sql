-- name: GetAllTickets :many
SELECT * FROM ticket;

-- name: CreateTicket :one
INSERT INTO ticket (name, description, price, quantity, event_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;