-- name: GetAllAccounts :many
SELECT * FROM account;

-- name: CreateAccount :one
INSERT INTO account (name, email)
VALUES ($1, $2)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM account
WHERE id = $1;