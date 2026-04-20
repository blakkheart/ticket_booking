-- name: GetAllAccounts :many
SELECT * FROM account;

-- name: CreateAccount :one
INSERT INTO account (id, name, email, password, role)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM account
WHERE id = $1;

-- name: GetAccountByEmail :one
SELECT * FROM account
WHERE email = $1;