// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: fund.sql

package db

import (
	"context"
	"database/sql"
)

const createFund = `-- name: CreateFund :one
INSERT INTO fund(
    user_id,
    balance,
    currency_type,
    cr_user
) VALUES(
    $1, $2, $3, $4
)   RETURNING fund_id, user_id, balance, currency_type, cr_date, up_date, cr_user, up_user
`

type CreateFundParams struct {
	UserID       int64  `json:"user_id"`
	Balance      string `json:"balance"`
	CurrencyType string `json:"currency_type"`
	CrUser       string `json:"cr_user"`
}

func (q *Queries) CreateFund(ctx context.Context, arg CreateFundParams) (Fund, error) {
	row := q.queryRow(ctx, q.createFundStmt, createFund,
		arg.UserID,
		arg.Balance,
		arg.CurrencyType,
		arg.CrUser,
	)
	var i Fund
	err := row.Scan(
		&i.FundID,
		&i.UserID,
		&i.Balance,
		&i.CurrencyType,
		&i.CrDate,
		&i.UpDate,
		&i.CrUser,
		&i.UpUser,
	)
	return i, err
}

const deleteFund = `-- name: DeleteFund :exec
DELETE FROM fund
WHERE fund_id = $1
`

func (q *Queries) DeleteFund(ctx context.Context, fundID int64) error {
	_, err := q.exec(ctx, q.deleteFundStmt, deleteFund, fundID)
	return err
}

const getFund = `-- name: GetFund :one
SELECT fund_id, user_id, balance, currency_type, cr_date, up_date, cr_user, up_user FROM fund
WHERE fund_id = $1 LIMIT 1
`

func (q *Queries) GetFund(ctx context.Context, fundID int64) (Fund, error) {
	row := q.queryRow(ctx, q.getFundStmt, getFund, fundID)
	var i Fund
	err := row.Scan(
		&i.FundID,
		&i.UserID,
		&i.Balance,
		&i.CurrencyType,
		&i.CrDate,
		&i.UpDate,
		&i.CrUser,
		&i.UpUser,
	)
	return i, err
}

const getfundByUidandFid = `-- name: GetfundByUidandFid :one
SELECT fund_id, user_id, balance, currency_type, cr_date, up_date, cr_user, up_user FROM fund
WHERE user_id = $1
AND fund_id = $2
`

type GetfundByUidandFidParams struct {
	UserID int64 `json:"user_id"`
	FundID int64 `json:"fund_id"`
}

func (q *Queries) GetfundByUidandFid(ctx context.Context, arg GetfundByUidandFidParams) (Fund, error) {
	row := q.queryRow(ctx, q.getfundByUidandFidStmt, getfundByUidandFid, arg.UserID, arg.FundID)
	var i Fund
	err := row.Scan(
		&i.FundID,
		&i.UserID,
		&i.Balance,
		&i.CurrencyType,
		&i.CrDate,
		&i.UpDate,
		&i.CrUser,
		&i.UpUser,
	)
	return i, err
}

const getfundByUidandFidForUpdateNoK = `-- name: GetfundByUidandFidForUpdateNoK :one
SELECT fund_id, user_id, balance, currency_type, cr_date, up_date, cr_user, up_user FROM fund
WHERE user_id = $1
AND fund_id = $2
FOR NO KEY UPDATE
`

type GetfundByUidandFidForUpdateNoKParams struct {
	UserID int64 `json:"user_id"`
	FundID int64 `json:"fund_id"`
}

func (q *Queries) GetfundByUidandFidForUpdateNoK(ctx context.Context, arg GetfundByUidandFidForUpdateNoKParams) (Fund, error) {
	row := q.queryRow(ctx, q.getfundByUidandFidForUpdateNoKStmt, getfundByUidandFidForUpdateNoK, arg.UserID, arg.FundID)
	var i Fund
	err := row.Scan(
		&i.FundID,
		&i.UserID,
		&i.Balance,
		&i.CurrencyType,
		&i.CrDate,
		&i.UpDate,
		&i.CrUser,
		&i.UpUser,
	)
	return i, err
}

const getfundByUserId = `-- name: GetfundByUserId :many
SELECT fund_id, user_id, balance, currency_type, cr_date, up_date, cr_user, up_user FROM fund
WHERE user_id = $1
LIMIT $2
OFFSET $3
`

type GetfundByUserIdParams struct {
	UserID int64 `json:"user_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetfundByUserId(ctx context.Context, arg GetfundByUserIdParams) ([]Fund, error) {
	rows, err := q.query(ctx, q.getfundByUserIdStmt, getfundByUserId, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Fund{}
	for rows.Next() {
		var i Fund
		if err := rows.Scan(
			&i.FundID,
			&i.UserID,
			&i.Balance,
			&i.CurrencyType,
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

const getfunds = `-- name: Getfunds :many
SELECT fund_id, user_id, balance, currency_type, cr_date, up_date, cr_user, up_user FROM  fund
ORDER BY fund_id
LIMIT $1
OFFSET $2
`

type GetfundsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) Getfunds(ctx context.Context, arg GetfundsParams) ([]Fund, error) {
	rows, err := q.query(ctx, q.getfundsStmt, getfunds, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Fund{}
	for rows.Next() {
		var i Fund
		if err := rows.Scan(
			&i.FundID,
			&i.UserID,
			&i.Balance,
			&i.CurrencyType,
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

const updateFund = `-- name: UpdateFund :one
UPDATE fund
SET balance = $2,
up_date = $3,
up_user = $4
WHERE fund_id = $1
RETURNING fund_id, user_id, balance, currency_type, cr_date, up_date, cr_user, up_user
`

type UpdateFundParams struct {
	FundID  int64          `json:"fund_id"`
	Balance string         `json:"balance"`
	UpDate  sql.NullTime   `json:"up_date"`
	UpUser  sql.NullString `json:"up_user"`
}

func (q *Queries) UpdateFund(ctx context.Context, arg UpdateFundParams) (Fund, error) {
	row := q.queryRow(ctx, q.updateFundStmt, updateFund,
		arg.FundID,
		arg.Balance,
		arg.UpDate,
		arg.UpUser,
	)
	var i Fund
	err := row.Scan(
		&i.FundID,
		&i.UserID,
		&i.Balance,
		&i.CurrencyType,
		&i.CrDate,
		&i.UpDate,
		&i.CrUser,
		&i.UpUser,
	)
	return i, err
}
