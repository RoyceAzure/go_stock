-- name: CreateRealizedProfitLoss :one
INSERT INTO realized_profit_loss(
    transation_id,
    user_id,
    product_name,
    cost_per_price,
    cost_total_price,
    realized,
    realized_precent
) VALUES(
    $1, $2, $3, $4,$5,$6,$7
)   RETURNING *;

-- name: GetRealizedProfitLoss :one
SELECT * FROM realized_profit_loss
WHERE id = $1 LIMIT 1;

-- name: GetRealizedProfitLosssByUserId :many
SELECT * FROM realized_profit_loss
WHERE user_id = $1
ORDER BY "transation_id"
LIMIT $2
OFFSET $3;

-- name: GetRealizedProfitLosssByUserIdDetial :many
SELECT rpl.user_id,
rpl.product_name,
rpl.cost_per_price,
rpl.cost_total_price,
st.transaction_type,
st.transation_amt,
st.transation_price_per_share,
rpl.realized,
rpl.realized_precent,
st.cr_date AS trans_at
FROM realized_profit_loss AS rpl
LEFT JOIN stock_transaction AS st
ON rpl.transation_id = st.transation_id
WHERE rpl.user_id = $1
ORDER BY rpl.product_name
LIMIT $2
OFFSET $3;

