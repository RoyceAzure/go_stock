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

-- name: GetUserStocksByUserId :many
SELECT user_stock.*, stock.stock_name, stock_code FROM  user_stock
LEFT JOIN stock 
ON  user_stock.stock_id = stock.stock_id
WHERE 
    (sqlc.narg(user_id)::bigint IS NULL OR user_stock.user_id = sqlc.narg(user_id))
    AND (sqlc.narg(stock_id)::bigint IS NULL OR user_stock.stock_id = sqlc.narg(stock_id))
    AND (sqlc.narg(stock_id)::bigint IS NULL OR user_stock.stock_id = sqlc.narg(stock_id))
    AND (sqlc.narg(purchased_date_start)::timestamptz IS NULL OR user_stock.purchased_date > sqlc.narg(purchased_date_start))
    AND (sqlc.narg(purchased_date_end)::timestamptz IS NULL OR user_stock.purchased_date < sqlc.narg(purchased_date_end))
ORDER BY user_stock_id
LIMIT sqlc.arg(limits)
OFFSET sqlc.arg(offsets);

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

