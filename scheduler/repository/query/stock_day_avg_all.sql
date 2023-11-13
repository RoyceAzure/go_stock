-- name: CreateSDAVGALL :one
INSERT INTO "stock_day_avg_all" (
    code,
    stock_name,
    close_price,
    monthly_avg_price
) VALUES(
    $1, $2, $3, $4
) RETURNING *;

-- name: BulkInsertDAVGALL :copyfrom
INSERT INTO "stock_day_avg_all" (
    code,
    stock_name,
    close_price,
    monthly_avg_price
) VALUES(
    $1, $2, $3, $4
);

-- name: GetSDAVGALLs :many
SELECT * FROM "stock_day_avg_all"
WHERE (sqlc.narg(id)::bigint IS NULL OR id = sqlc.narg(id))
    AND (sqlc.narg(code)::varchar IS NULL OR code = sqlc.narg(code))
    AND (sqlc.narg(stock_name)::varchar IS NULL OR stock_name = sqlc.narg(stock_name))
    AND (sqlc.narg(cp_upper)::decimal IS NULL OR close_price <= sqlc.narg(cp_upper))
    AND (sqlc.narg(cp_lower)::decimal IS NULL OR close_price >= sqlc.narg(cp_lower))
    AND (sqlc.narg(map_upper)::decimal IS NULL OR monthly_avg_price <= sqlc.narg(map_upper))
    AND (sqlc.narg(map_lower)::decimal IS NULL OR monthly_avg_price >= sqlc.narg(map_lower))
    AND (sqlc.narg(cr_date_start)::timestamptz IS NULL OR cr_date >= sqlc.narg(cr_date_start))
    AND (sqlc.narg(cr_date_end)::timestamptz IS NULL OR cr_date <= sqlc.narg(cr_date_end))
ORDER BY code
LIMIT sqlc.arg(limits)
OFFSET sqlc.arg(offsets);


-- name: BatchDeleteSDAVGALL :exec
DELETE FROM "stock_day_avg_all"
WHERE  id = ANY($1::bigint[]);

-- name: DeleteSDAVGALLCodePrexForTest :exec
DELETE FROM "stock_day_avg_all"
WHERE id in (
    SELECT id 
    FROM "stock_day_avg_all" as s
    WHERE substring(s.code, 0, sqlc.arg(len)) = sqlc.arg(code_prefix)
);
