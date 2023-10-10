-- name: CreateStock :one
INSERT INTO stock(
    ticker_symbol,
    comp_name,
    current_price,
    market_cap,
    cr_user
) VALUES(
    $1, $2, $3, $4, $5
)   RETURNING *;

-- name: GetStock :one
SELECT * FROM stock
WHERE stock_id = $1 LIMIT 1;

-- name: GetstockByTS :many
SELECT * FROM stock
WHERE ticker_symbol = $1
LIMIT $2
OFFSET $3;

-- name: GetstockByCN :many
SELECT * FROM stock
WHERE comp_name = $1
LIMIT $2
OFFSET $3;

-- name: Getstocks :many
SELECT * FROM  stock
ORDER BY stock_id
LIMIT $1
OFFSET $2;

-- name: UpdateStock :one
UPDATE stock
SET current_price = $2,
up_date = $3,
up_user = $4
WHERE stock_id = $1
RETURNING *;

-- name: DeleteStock :exec
DELETE FROM stock
WHERE stock_id = $1;
