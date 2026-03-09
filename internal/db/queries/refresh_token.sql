-- name: GetAllTokens :many
SELECT * FROM refresh_token;

-- name: CreateToken :one
INSERT INTO refresh_token (user_id, token_hash, expires_at, revoked, replaced_by)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetTokenByUserID :one
SELECT * FROM refresh_token
WHERE user_id = $1;


-- name: UpdateTokenByUserID :exec
UPDATE refresh_token
SET token_hash = $2, expires_at = $3, revoked = $4, replaced_by = $5
WHERE user_id = $1;


-- name: GetTokenByHash :one
SELECT * FROM refresh_token
WHERE token_hash = $1;