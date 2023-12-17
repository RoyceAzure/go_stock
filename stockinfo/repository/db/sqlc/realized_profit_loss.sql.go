// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: realized_profit_loss.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createRealizedProfitLoss = `-- name: CreateRealizedProfitLoss :one
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
)   RETURNING id, transation_id, user_id, product_name, cost_per_price, cost_total_price, realized, realized_precent
`

type CreateRealizedProfitLossParams struct {
	TransationID    uuid.UUID `json:"transation_id"`
	UserID          int64     `json:"user_id"`
	ProductName     string    `json:"product_name"`
	CostPerPrice    string    `json:"cost_per_price"`
	CostTotalPrice  string    `json:"cost_total_price"`
	Realized        string    `json:"realized"`
	RealizedPrecent string    `json:"realized_precent"`
}

func (q *Queries) CreateRealizedProfitLoss(ctx context.Context, arg CreateRealizedProfitLossParams) (RealizedProfitLoss, error) {
	row := q.db.QueryRowContext(ctx, createRealizedProfitLoss,
		arg.TransationID,
		arg.UserID,
		arg.ProductName,
		arg.CostPerPrice,
		arg.CostTotalPrice,
		arg.Realized,
		arg.RealizedPrecent,
	)
	var i RealizedProfitLoss
	err := row.Scan(
		&i.ID,
		&i.TransationID,
		&i.UserID,
		&i.ProductName,
		&i.CostPerPrice,
		&i.CostTotalPrice,
		&i.Realized,
		&i.RealizedPrecent,
	)
	return i, err
}

const getRealizedProfitLoss = `-- name: GetRealizedProfitLoss :one
SELECT id, transation_id, user_id, product_name, cost_per_price, cost_total_price, realized, realized_precent FROM realized_profit_loss
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRealizedProfitLoss(ctx context.Context, id int64) (RealizedProfitLoss, error) {
	row := q.db.QueryRowContext(ctx, getRealizedProfitLoss, id)
	var i RealizedProfitLoss
	err := row.Scan(
		&i.ID,
		&i.TransationID,
		&i.UserID,
		&i.ProductName,
		&i.CostPerPrice,
		&i.CostTotalPrice,
		&i.Realized,
		&i.RealizedPrecent,
	)
	return i, err
}

const getRealizedProfitLosssByUserId = `-- name: GetRealizedProfitLosssByUserId :many
SELECT id, transation_id, user_id, product_name, cost_per_price, cost_total_price, realized, realized_precent FROM realized_profit_loss
WHERE user_id = $1
ORDER BY "transation_id"
LIMIT $2
OFFSET $3
`

type GetRealizedProfitLosssByUserIdParams struct {
	UserID int64 `json:"user_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetRealizedProfitLosssByUserId(ctx context.Context, arg GetRealizedProfitLosssByUserIdParams) ([]RealizedProfitLoss, error) {
	rows, err := q.db.QueryContext(ctx, getRealizedProfitLosssByUserId, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []RealizedProfitLoss{}
	for rows.Next() {
		var i RealizedProfitLoss
		if err := rows.Scan(
			&i.ID,
			&i.TransationID,
			&i.UserID,
			&i.ProductName,
			&i.CostPerPrice,
			&i.CostTotalPrice,
			&i.Realized,
			&i.RealizedPrecent,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRealizedProfitLosssByUserIdDetial = `-- name: GetRealizedProfitLosssByUserIdDetial :many
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
OFFSET $3
`

type GetRealizedProfitLosssByUserIdDetialParams struct {
	UserID int64 `json:"user_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetRealizedProfitLosssByUserIdDetialRow struct {
	UserID                  int64          `json:"user_id"`
	ProductName             string         `json:"product_name"`
	CostPerPrice            string         `json:"cost_per_price"`
	CostTotalPrice          string         `json:"cost_total_price"`
	TransactionType         sql.NullString `json:"transaction_type"`
	TransationAmt           sql.NullInt32  `json:"transation_amt"`
	TransationPricePerShare sql.NullString `json:"transation_price_per_share"`
	Realized                string         `json:"realized"`
	RealizedPrecent         string         `json:"realized_precent"`
	TransAt                 sql.NullTime   `json:"trans_at"`
}

func (q *Queries) GetRealizedProfitLosssByUserIdDetial(ctx context.Context, arg GetRealizedProfitLosssByUserIdDetialParams) ([]GetRealizedProfitLosssByUserIdDetialRow, error) {
	rows, err := q.db.QueryContext(ctx, getRealizedProfitLosssByUserIdDetial, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetRealizedProfitLosssByUserIdDetialRow{}
	for rows.Next() {
		var i GetRealizedProfitLosssByUserIdDetialRow
		if err := rows.Scan(
			&i.UserID,
			&i.ProductName,
			&i.CostPerPrice,
			&i.CostTotalPrice,
			&i.TransactionType,
			&i.TransationAmt,
			&i.TransationPricePerShare,
			&i.Realized,
			&i.RealizedPrecent,
			&i.TransAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
