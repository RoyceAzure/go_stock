// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: stock.sql

package db

import (
	"context"
	"database/sql"
)

const createStock = `-- name: CreateStock :one
INSERT INTO stock(
    stock_code,
    stock_name,
    current_price,
    market_cap,
    cr_user
) VALUES(
    $1, $2, $3, $4, $5
)   RETURNING stock_id, stock_code, stock_name, current_price, market_cap, cr_date, up_date, cr_user, up_user
`

type CreateStockParams struct {
	StockCode    string `json:"stock_code"`
	StockName    string `json:"stock_name"`
	CurrentPrice string `json:"current_price"`
	MarketCap    int64  `json:"market_cap"`
	CrUser       string `json:"cr_user"`
}

func (q *Queries) CreateStock(ctx context.Context, arg CreateStockParams) (Stock, error) {
	row := q.db.QueryRowContext(ctx, createStock,
		arg.StockCode,
		arg.StockName,
		arg.CurrentPrice,
		arg.MarketCap,
		arg.CrUser,
	)
	var i Stock
	err := row.Scan(
		&i.StockID,
		&i.StockCode,
		&i.StockName,
		&i.CurrentPrice,
		&i.MarketCap,
		&i.CrDate,
		&i.UpDate,
		&i.CrUser,
		&i.UpUser,
	)
	return i, err
}

const deleteStock = `-- name: DeleteStock :exec
DELETE FROM stock
WHERE stock_id = $1
`

func (q *Queries) DeleteStock(ctx context.Context, stockID int64) error {
	_, err := q.db.ExecContext(ctx, deleteStock, stockID)
	return err
}

const getStock = `-- name: GetStock :one
SELECT stock_id, stock_code, stock_name, current_price, market_cap, cr_date, up_date, cr_user, up_user FROM stock
WHERE stock_id = $1 LIMIT 1
`

func (q *Queries) GetStock(ctx context.Context, stockID int64) (Stock, error) {
	row := q.db.QueryRowContext(ctx, getStock, stockID)
	var i Stock
	err := row.Scan(
		&i.StockID,
		&i.StockCode,
		&i.StockName,
		&i.CurrentPrice,
		&i.MarketCap,
		&i.CrDate,
		&i.UpDate,
		&i.CrUser,
		&i.UpUser,
	)
	return i, err
}

const getstockByCN = `-- name: GetstockByCN :many
SELECT stock_id, stock_code, stock_name, current_price, market_cap, cr_date, up_date, cr_user, up_user FROM stock
WHERE stock_name = $1
LIMIT $2
OFFSET $3
`

type GetstockByCNParams struct {
	StockName string `json:"stock_name"`
	Limit     int32  `json:"limit"`
	Offset    int32  `json:"offset"`
}

func (q *Queries) GetstockByCN(ctx context.Context, arg GetstockByCNParams) ([]Stock, error) {
	rows, err := q.db.QueryContext(ctx, getstockByCN, arg.StockName, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Stock{}
	for rows.Next() {
		var i Stock
		if err := rows.Scan(
			&i.StockID,
			&i.StockCode,
			&i.StockName,
			&i.CurrentPrice,
			&i.MarketCap,
			&i.CrDate,
			&i.UpDate,
			&i.CrUser,
			&i.UpUser,
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

const getstockByTS = `-- name: GetstockByTS :many
SELECT stock_id, stock_code, stock_name, current_price, market_cap, cr_date, up_date, cr_user, up_user FROM stock
WHERE stock_code = $1
LIMIT $2
OFFSET $3
`

type GetstockByTSParams struct {
	StockCode string `json:"stock_code"`
	Limit     int32  `json:"limit"`
	Offset    int32  `json:"offset"`
}

func (q *Queries) GetstockByTS(ctx context.Context, arg GetstockByTSParams) ([]Stock, error) {
	rows, err := q.db.QueryContext(ctx, getstockByTS, arg.StockCode, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Stock{}
	for rows.Next() {
		var i Stock
		if err := rows.Scan(
			&i.StockID,
			&i.StockCode,
			&i.StockName,
			&i.CurrentPrice,
			&i.MarketCap,
			&i.CrDate,
			&i.UpDate,
			&i.CrUser,
			&i.UpUser,
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

const getstocks = `-- name: Getstocks :many
SELECT stock_id, stock_code, stock_name, current_price, market_cap, cr_date, up_date, cr_user, up_user FROM  stock
ORDER BY stock_id
LIMIT $1
OFFSET $2
`

type GetstocksParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) Getstocks(ctx context.Context, arg GetstocksParams) ([]Stock, error) {
	rows, err := q.db.QueryContext(ctx, getstocks, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Stock{}
	for rows.Next() {
		var i Stock
		if err := rows.Scan(
			&i.StockID,
			&i.StockCode,
			&i.StockName,
			&i.CurrentPrice,
			&i.MarketCap,
			&i.CrDate,
			&i.UpDate,
			&i.CrUser,
			&i.UpUser,
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

const updateStock = `-- name: UpdateStock :one
UPDATE stock
SET current_price = $2,
up_date = $3,
up_user = $4
WHERE stock_id = $1
RETURNING stock_id, stock_code, stock_name, current_price, market_cap, cr_date, up_date, cr_user, up_user
`

type UpdateStockParams struct {
	StockID      int64          `json:"stock_id"`
	CurrentPrice string         `json:"current_price"`
	UpDate       sql.NullTime   `json:"up_date"`
	UpUser       sql.NullString `json:"up_user"`
}

func (q *Queries) UpdateStock(ctx context.Context, arg UpdateStockParams) (Stock, error) {
	row := q.db.QueryRowContext(ctx, updateStock,
		arg.StockID,
		arg.CurrentPrice,
		arg.UpDate,
		arg.UpUser,
	)
	var i Stock
	err := row.Scan(
		&i.StockID,
		&i.StockCode,
		&i.StockName,
		&i.CurrentPrice,
		&i.MarketCap,
		&i.CrDate,
		&i.UpDate,
		&i.CrUser,
		&i.UpUser,
	)
	return i, err
}

const updateStockCPByCode = `-- name: UpdateStockCPByCode :one
UPDATE stock
SET 
    stock_name = COALESCE($1, stock_name),
    current_price = COALESCE($2, current_price)
WHERE stock.stock_code = $3
RETURNING stock_id, stock_code, stock_name, current_price, market_cap, cr_date, up_date, cr_user, up_user
`

type UpdateStockCPByCodeParams struct {
	StockName    sql.NullString `json:"stock_name"`
	CurrentPrice sql.NullString `json:"current_price"`
	StockCode    string         `json:"stock_code"`
}

func (q *Queries) UpdateStockCPByCode(ctx context.Context, arg UpdateStockCPByCodeParams) (Stock, error) {
	row := q.db.QueryRowContext(ctx, updateStockCPByCode, arg.StockName, arg.CurrentPrice, arg.StockCode)
	var i Stock
	err := row.Scan(
		&i.StockID,
		&i.StockCode,
		&i.StockName,
		&i.CurrentPrice,
		&i.MarketCap,
		&i.CrDate,
		&i.UpDate,
		&i.CrUser,
		&i.UpUser,
	)
	return i, err
}
