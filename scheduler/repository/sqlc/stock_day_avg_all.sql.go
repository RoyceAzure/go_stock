// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: stock_day_avg_all.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const batchDeleteSDAVGALL = `-- name: BatchDeleteSDAVGALL :exec
DELETE FROM "stock_day_avg_all"
WHERE  id = ANY($1::bigint[])
`

func (q *Queries) BatchDeleteSDAVGALL(ctx context.Context, dollar_1 []int64) error {
	_, err := q.db.Exec(ctx, batchDeleteSDAVGALL, dollar_1)
	return err
}

type BulkInsertDAVGALLParams struct {
	Code            string         `json:"code"`
	StockName       string         `json:"stock_name"`
	ClosePrice      pgtype.Numeric `json:"close_price"`
	MonthlyAvgPrice pgtype.Numeric `json:"monthly_avg_price"`
}

const createSDAVGALL = `-- name: CreateSDAVGALL :one
INSERT INTO "stock_day_avg_all" (
    code,
    stock_name,
    close_price,
    monthly_avg_price
) VALUES(
    $1, $2, $3, $4
) RETURNING id, code, stock_name, close_price, monthly_avg_price, cr_date, up_date, cr_user, up_user
`

type CreateSDAVGALLParams struct {
	Code            string         `json:"code"`
	StockName       string         `json:"stock_name"`
	ClosePrice      pgtype.Numeric `json:"close_price"`
	MonthlyAvgPrice pgtype.Numeric `json:"monthly_avg_price"`
}

func (q *Queries) CreateSDAVGALL(ctx context.Context, arg CreateSDAVGALLParams) (StockDayAvgAll, error) {
	row := q.db.QueryRow(ctx, createSDAVGALL,
		arg.Code,
		arg.StockName,
		arg.ClosePrice,
		arg.MonthlyAvgPrice,
	)
	var i StockDayAvgAll
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.StockName,
		&i.ClosePrice,
		&i.MonthlyAvgPrice,
		&i.CrDate,
		&i.UpDate,
		&i.CrUser,
		&i.UpUser,
	)
	return i, err
}

const deleteSDAVGALLCodePrexForTest = `-- name: DeleteSDAVGALLCodePrexForTest :exec
DELETE FROM "stock_day_avg_all"
WHERE id in (
    SELECT id 
    FROM "stock_day_avg_all" as s
    WHERE substring(s.code, 0, $1) = $2
)
`

type DeleteSDAVGALLCodePrexForTestParams struct {
	Len        int32  `json:"len"`
	CodePrefix string `json:"code_prefix"`
}

func (q *Queries) DeleteSDAVGALLCodePrexForTest(ctx context.Context, arg DeleteSDAVGALLCodePrexForTestParams) error {
	_, err := q.db.Exec(ctx, deleteSDAVGALLCodePrexForTest, arg.Len, arg.CodePrefix)
	return err
}

const getSDAVGALLs = `-- name: GetSDAVGALLs :many
SELECT id, code, stock_name, close_price, monthly_avg_price, cr_date, up_date, cr_user, up_user FROM "stock_day_avg_all"
WHERE ($1::bigint IS NULL OR id = $1)
    AND ($2::varchar IS NULL OR code = $2)
    AND ($3::varchar IS NULL OR stock_name = $3)
    AND ($4::decimal IS NULL OR close_price <= $4)
    AND ($5::decimal IS NULL OR close_price >= $5)
    AND ($6::decimal IS NULL OR monthly_avg_price <= $6)
    AND ($7::decimal IS NULL OR monthly_avg_price >= $7)
    AND ($8::timestamptz IS NULL OR cr_date >= $8)
    AND ($9::timestamptz IS NULL OR cr_date <= $9)
ORDER BY code
LIMIT $11
OFFSET $10
`

type GetSDAVGALLsParams struct {
	ID          pgtype.Int8        `json:"id"`
	Code        pgtype.Text        `json:"code"`
	StockName   pgtype.Text        `json:"stock_name"`
	CpUpper     pgtype.Numeric     `json:"cp_upper"`
	CpLower     pgtype.Numeric     `json:"cp_lower"`
	MapUpper    pgtype.Numeric     `json:"map_upper"`
	MapLower    pgtype.Numeric     `json:"map_lower"`
	CrDateStart pgtype.Timestamptz `json:"cr_date_start"`
	CrDateEnd   pgtype.Timestamptz `json:"cr_date_end"`
	Offsets     int32              `json:"offsets"`
	Limits      int32              `json:"limits"`
}

func (q *Queries) GetSDAVGALLs(ctx context.Context, arg GetSDAVGALLsParams) ([]StockDayAvgAll, error) {
	rows, err := q.db.Query(ctx, getSDAVGALLs,
		arg.ID,
		arg.Code,
		arg.StockName,
		arg.CpUpper,
		arg.CpLower,
		arg.MapUpper,
		arg.MapLower,
		arg.CrDateStart,
		arg.CrDateEnd,
		arg.Offsets,
		arg.Limits,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []StockDayAvgAll{}
	for rows.Next() {
		var i StockDayAvgAll
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.StockName,
			&i.ClosePrice,
			&i.MonthlyAvgPrice,
			&i.CrDate,
			&i.UpDate,
			&i.CrUser,
			&i.UpUser,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
