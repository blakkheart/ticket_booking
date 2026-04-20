-- name: GetAllTickets :many
SELECT * FROM ticket_type;

-- name: CreateTicket :one
INSERT INTO ticket_type (id, name, description, price, available_quantity, event_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetTicketByID :one
SELECT sqlc.embed(ticket_type), sqlc.embed(event)
FROM ticket_type
JOIN event ON event.id == ticket_type.event_id
WHERE ticket_type.id = $1;


-- name: UpdateTicketQuantityByID :one
UPDATE ticket_type
SET available_quantity = $2
WHERE id = $1 AND available_quantity >= $2
RETURNING *;
