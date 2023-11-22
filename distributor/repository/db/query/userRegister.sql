-- name: CreateUserRegister :one
INSERT INTO "user_register"(
    user_id,
    stock_code,
    updated_at
) VALUES(
    $1, $2, $3
)   RETURNING *;

-- name: GetUserRegisterByUserID :many
SELECT * FROM "user_register"
WHERE user_id = sqlc.arg(user_id);

-- name: GetUserRegisters :many
SELECT * FROM  "user_register"
ORDER BY user_id
LIMIT $1
OFFSET $2;

-- name: DeleteUserRegister :exec
DELETE FROM "user_register"
WHERE user_id = sqlc.arg(user_id)
    AND stock_code = sqlc.arg(stock_code);
