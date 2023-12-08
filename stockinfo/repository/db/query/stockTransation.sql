-- name: CreateStockTransaction :one
INSERT INTO stock_transaction(
    user_id,
    stock_id,
    fund_id,
    transaction_type,
    transaction_date,
    transation_amt,
    transation_price_per_share,
    cr_user
) VALUES(
    $1, $2, $3, $4,$5,$6,$7,$8
)   RETURNING *;

-- name: GetStockTransaction :one
SELECT * FROM stock_transaction
WHERE "transation_id" = $1 LIMIT 1;

-- name: GetStockTransactionsByUserId :many
SELECT * FROM stock_transaction
WHERE user_id = $1
ORDER BY "transation_id"
LIMIT $2
OFFSET $3;

-- name: GetStockTransactionsByStockId :many
SELECT * FROM stock_transaction
WHERE stock_id = $1
ORDER BY "transation_id"
LIMIT $2
OFFSET $3;

-- name: GetStockTransactionsByDate :many
SELECT * FROM stock_transaction
WHERE transaction_date = $1 
ORDER BY transaction_date
LIMIT $2
OFFSET $3;


-- name: GetStockTransactionsFilter :many
SELECT * FROM  stock_transaction
WHERE 
    user_id =  COALESCE(NULLIF(sqlc.arg(user_id),0), user_id)
    AND stock_id = COALESCE(NULLIF(sqlc.arg(stock_id),0), stock_id)
    AND transaction_type = COALESCE(sqlc.narg(transaction_type), transaction_type)
ORDER BY "transation_id"
LIMIT sqlc.arg(limits)
OFFSET sqlc.arg(offsets);


-- name: DeleteStockTransaction :exec
DELETE FROM stock_transaction
WHERE "transation_id" = $1;


-- name: UpdateStockTransationResult :one
Update stock_transaction
SET result = sqlc.arg(result)
WHERE transation_id = sqlc.arg(transation_id)
RETURNING *;

