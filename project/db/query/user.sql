-- name: CreateUser :one
INSERT INTO "user"(
    user_name,
    email,
    hashed_password,
    password_changed_at,
    sso_identifer,
    cr_user
) VALUES(
    $1, $2, $3, $4, $5, $6
)   RETURNING *;

-- name: GetUser :one
SELECT * FROM "user"
WHERE user_id = $1 LIMIT 1;

-- name: GetUserForUpdateNoKey :one
SELECT * FROM "user"
WHERE user_id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: GetUserByEmail :one
SELECT * FROM "user"
WHERE email = $1 LIMIT 1;

-- name: Getusers :many
SELECT * FROM  "user"
ORDER BY user_id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE "user"
SET 
    user_name = COALESCE(sqlc.narg(user_name), user_name),  
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
    password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
    sso_identifer = COALESCE(sqlc.narg(sso_identifer), sso_identifer),
    up_date = COALESCE(sqlc.narg(up_date), up_date),
    up_user = COALESCE(sqlc.narg(up_user), up_user)
WHERE user_id = sqlc.arg(user_id)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE user_id = $1;
