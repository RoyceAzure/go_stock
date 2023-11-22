// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: userRegister.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUserRegister = `-- name: CreateUserRegister :one
INSERT INTO "user_register"(
    user_id,
    stock_code,
    updated_at
) VALUES(
    $1, $2, $3
)   RETURNING user_id, stock_code, created_at, updated_at
`

type CreateUserRegisterParams struct {
	UserID    int64              `json:"user_id"`
	StockCode string             `json:"stock_code"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) CreateUserRegister(ctx context.Context, arg CreateUserRegisterParams) (UserRegister, error) {
	row := q.db.QueryRow(ctx, createUserRegister, arg.UserID, arg.StockCode, arg.UpdatedAt)
	var i UserRegister
	err := row.Scan(
		&i.UserID,
		&i.StockCode,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUserRegister = `-- name: DeleteUserRegister :exec
DELETE FROM "user_register"
WHERE user_id = $1
    AND stock_code = $2
`

type DeleteUserRegisterParams struct {
	UserID    int64  `json:"user_id"`
	StockCode string `json:"stock_code"`
}

func (q *Queries) DeleteUserRegister(ctx context.Context, arg DeleteUserRegisterParams) error {
	_, err := q.db.Exec(ctx, deleteUserRegister, arg.UserID, arg.StockCode)
	return err
}

const getUserRegisterByUserID = `-- name: GetUserRegisterByUserID :many
SELECT user_id, stock_code, created_at, updated_at FROM "user_register"
WHERE user_id = $1
`

func (q *Queries) GetUserRegisterByUserID(ctx context.Context, userID int64) ([]UserRegister, error) {
	rows, err := q.db.Query(ctx, getUserRegisterByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserRegister{}
	for rows.Next() {
		var i UserRegister
		if err := rows.Scan(
			&i.UserID,
			&i.StockCode,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getUserRegisters = `-- name: GetUserRegisters :many
SELECT user_id, stock_code, created_at, updated_at FROM  "user_register"
ORDER BY user_id
LIMIT $1
OFFSET $2
`

type GetUserRegistersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetUserRegisters(ctx context.Context, arg GetUserRegistersParams) ([]UserRegister, error) {
	rows, err := q.db.Query(ctx, getUserRegisters, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserRegister{}
	for rows.Next() {
		var i UserRegister
		if err := rows.Scan(
			&i.UserID,
			&i.StockCode,
			&i.CreatedAt,
			&i.UpdatedAt,
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
