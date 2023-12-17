-- name: CreateFrontendClient :one
INSERT INTO "frontend_client"(
    ip,
    region,
    updated_at
) VALUES(
    $1, $2, $3
)   RETURNING *;

-- name: GetFrontendClientByID :one
SELECT * FROM "frontend_client"
WHERE client_uid = sqlc.arg(client_uid);

-- name: GetFrontendClientByIP :one
SELECT * FROM "frontend_client"
WHERE Ip = sqlc.arg(Ip);

-- name: GetFrontendClients :many
SELECT * FROM  "frontend_client"
ORDER BY client_uid
LIMIT $1
OFFSET $2;

-- name: DeleteFrontendClient :exec
DELETE FROM "frontend_client"
WHERE client_uid = sqlc.arg(client_uid) CASCADE;