-- name: CreateFund :one
INSERT INTO fund(
    user_id,
    balance,
    currency_type,
    cr_user
) VALUES(
    $1, $2, $3, $4
)   RETURNING *;

-- name: GetFund :one
SELECT * FROM fund
WHERE fund_id = $1 LIMIT 1;

-- name: GetfundByUserId :many
SELECT * FROM fund
WHERE user_id = $1
LIMIT $2
OFFSET $3;

-- name: GetfundByUidandFid :one
SELECT * FROM fund
WHERE user_id = $1
AND fund_id = $2;

-- name: GetfundByUidandFidForUpdateNoK :one
SELECT * FROM fund
WHERE user_id = $1
AND fund_id = $2
FOR NO KEY UPDATE;


-- name: Getfunds :many
SELECT * FROM  fund
ORDER BY fund_id
LIMIT $1
OFFSET $2;

-- name: UpdateFund :one
UPDATE fund
SET balance = $2,
up_date = $3,
up_user = $4
WHERE fund_id = $1
RETURNING *;

-- name: DeleteFund :exec
DELETE FROM fund
WHERE fund_id = $1;
