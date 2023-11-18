-- name: CreateSPR :one
INSERT INTO "stock_price_realtime" (
    code,
    stock_name,
    trade_volume,
    trade_value,
    opening_price,
    highest_price,
    lowest_price,
    closing_price,
    change,
    transaction,
    trans_time
) VALUES(
    $1, $2, $3, $4,$5,$6,$7,$8,$9,$10, $11
) RETURNING *;

-- name: BulkInsertSPR :copyfrom
INSERT INTO "stock_price_realtime" (
    code,
    stock_name,
    trade_volume,
    trade_value,
    opening_price,
    highest_price,
    lowest_price,
    closing_price,
    change,
    transaction,
    trans_time
) VALUES(
    $1, $2, $3, $4,$5,$6,$7,$8,$9,$10, $11
);

-- name: GetSPRs :many
SELECT * FROM "stock_price_realtime"
WHERE (sqlc.narg(code)::varchar IS NULL OR code = sqlc.narg(code))
    AND (sqlc.narg(stock_name)::varchar IS NULL OR stock_name = sqlc.narg(stock_name))
    AND (sqlc.narg(trans_time_start)::timestamptz IS NULL OR trans_time >= sqlc.narg(trans_time_start))
    AND (sqlc.narg(trans_time_end)::timestamptz IS NULL OR trans_time <= sqlc.narg(trans_time_end))
ORDER BY code
LIMIT sqlc.arg(limits)
OFFSET sqlc.arg(offsets);