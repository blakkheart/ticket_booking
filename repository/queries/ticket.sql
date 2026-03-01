-- name: GetAllTickets :many
SELECT * FROM ticket;

-- name: CreateTicket :one
INSERT INTO ticket (name, description, price, quantity, event_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetTicket :one
SELECT sqlc.embed(ticket), sqlc.embed(event)
FROM ticket
JOIN event ON event.id == ticket.event_id
WHERE ticket.id = $1;
