-- name: CreateStockTransaction :one
INSERT INTO stock_transaction(
    user_id,
    stock_id,
    transaction_type,
    transaction_date,
    transation_amt,
    transation_proce_per_share,
    cr_user
) VALUES(
    $1, $2, $3, $4,$5,$6,$7
)   RETURNING *;

-- name: GetStockTransaction :one
SELECT * FROM stock_transaction
WHERE "TransationId" = $1 LIMIT 1;

-- name: GetStockTransactionsByUserId :many
SELECT * FROM stock_transaction
WHERE user_id = $1
ORDER BY "TransationId"
LIMIT $2
OFFSET $3;

-- name: GetStockTransactionsByStockId :many
SELECT * FROM stock_transaction
WHERE stock_id = $1
ORDER BY "TransationId"
LIMIT $2
OFFSET $3;

-- name: GetStockTransactionsByDate :many
SELECT * FROM stock_transaction
WHERE transaction_date = $1 
ORDER BY transaction_date
LIMIT $2
OFFSET $3;


-- name: GetStockTransactions :many
SELECT * FROM  stock_transaction
ORDER BY "TransationId"
LIMIT $1
OFFSET $2;


-- name: DeleteStockTransaction :exec
DELETE FROM stock_transaction
WHERE "TransationId" = $1;
