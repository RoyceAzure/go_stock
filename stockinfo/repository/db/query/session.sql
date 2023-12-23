-- name: CreateSession :one
INSERT INTO session(
  id,
  user_id,
  refresh_token,
  user_agent,
  client_ip,
  is_blocked,
  expired_at
) VALUES(
    $1, $2, $3, $4, $5, $6, $7
)   RETURNING *;

-- name: GetSession :one
SELECT * FROM session
WHERE id = $1 LIMIT 1;

-- name: DeleteSession :exec
DELETE FROM session
WHERE id = $1;


-- name: GetSessionByUserId :one
SELECT * FROM session
WHERE user_id = $1 LIMIT 1;