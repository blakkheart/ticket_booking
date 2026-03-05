-- name: GetAllTokens :many
SELECT * FROM refresh_token;

-- name: CreateToken :one
INSERT INTO refresh_token (user_id, token_hash, expires_at, revoked)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetTokenByUserID :one
SELECT * FROM refresh_token
WHERE user_id = $1;
