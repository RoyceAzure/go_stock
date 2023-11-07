-- name: CreateSDAVGALL :one
INSERT INTO "stock_day_avg_all" (
    code,
    stock_name,
    close_price,
    monthly_avg_price
) VALUES(
    $1, $2, $3, $4
) RETURNING *;

-- name: GetSDAVGALLs :many
SELECT * FROM "stock_day_avg_all"
ORDER BY code
LIMIT $1
OFFSET $2;


-- name: BatchDeleteSDAVGALL :exec
DELETE FROM "stock_day_avg_all"
WHERE  id = ANY($1::bigint[]);