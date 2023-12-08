-- name: CreateUserStock :one
INSERT INTO user_stock(
    user_id,
    stock_id,
    quantity,
    purchase_price_per_share,
    purchased_date,
    cr_user
) VALUES(
    $1, $2, $3, $4, $5, $6
)   RETURNING *;

-- name: GetUserStock :one
SELECT * FROM user_stock
WHERE user_stock_id = $1 LIMIT 1;

-- name: GetUserStocks :many
SELECT * FROM  user_stock
ORDER BY user_stock_id
LIMIT $1
OFFSET $2;

-- name: GetUserStocksByUserId :many
SELECT * FROM  user_stock
WHERE user_id = $1
ORDER BY user_stock_id
LIMIT $2
OFFSET $3;

-- name: GetUserStocksByStockId :many
SELECT * FROM  user_stock
WHERE stock_id = $1
ORDER BY user_stock_id
LIMIT $2
OFFSET $3;

-- name: GetUserStocksByPDate :many
SELECT * FROM  user_stock
WHERE purchased_date = $1
ORDER BY user_stock_id
LIMIT $2
OFFSET $3;

-- name: GetUserStocksByUserAStock :many
SELECT * FROM  user_stock
WHERE purchased_date = $1
AND stock_id = $2
ORDER BY user_stock_id
LIMIT $3
OFFSET $4;

-- name: UpdateUserStock :one
UPDATE user_stock
SET quantity = $3,
purchase_price_per_share = $4,
purchased_date = $5,
up_date = $6,
up_user = $7
WHERE user_id = $1
AND stock_id = $2
RETURNING *;

-- name: DeleteUserStock :exec
DELETE FROM user_stock
WHERE user_stock_id = $1;


-- name: GetserStockByUidandSid :one
SELECT * FROM user_stock
WHERE user_id = $1
AND stock_id = $2;


-- name: GetUserStockByUidandSidForUpdateNoK :one
SELECT * FROM user_stock
WHERE user_id = $1
AND stock_id = $2
FOR NO KEY UPDATE;

