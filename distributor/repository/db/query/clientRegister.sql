-- name: CreateClientRegister :one
INSERT INTO "client_register"(
    client_uid,
    stock_code,
    created_at,
    updated_at
) VALUES(
    $1, $2, $3, $4
)   RETURNING *;

-- name: GetClientRegisterByClientUID :many
SELECT * FROM "client_register"
WHERE client_uid = $1;

-- name: GetClientRegisters :many
SELECT * FROM  "client_register"
ORDER BY client_uid
LIMIT $1
OFFSET $2;

-- name: DeleteClientRegister :exec
DELETE FROM "client_register"
WHERE client_uid = sqlc.arg(client_uid)
    AND stock_code = sqlc.arg(stock_code);
