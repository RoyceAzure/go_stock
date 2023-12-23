// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: session.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createSession = `-- name: CreateSession :one
INSERT INTO session(
  id,
  user_id,
  refresh_token,
  user_agent,
  client_ip,
  is_blocked,
  expired_at
) VALUES(
    $1, $2, $3, $4, $5, $6, $7
)   RETURNING id, user_id, refresh_token, user_agent, client_ip, is_blocked, cr_date, expired_at
`

type CreateSessionParams struct {
	ID           uuid.UUID `json:"id"`
	UserID       int64     `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiredAt    time.Time `json:"expired_at"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, createSession,
		arg.ID,
		arg.UserID,
		arg.RefreshToken,
		arg.UserAgent,
		arg.ClientIp,
		arg.IsBlocked,
		arg.ExpiredAt,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.RefreshToken,
		&i.UserAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.CrDate,
		&i.ExpiredAt,
	)
	return i, err
}

const deleteSession = `-- name: DeleteSession :exec
DELETE FROM session
WHERE id = $1
`

func (q *Queries) DeleteSession(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteSession, id)
	return err
}

const getSession = `-- name: GetSession :one
SELECT id, user_id, refresh_token, user_agent, client_ip, is_blocked, cr_date, expired_at FROM session
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSession(ctx context.Context, id uuid.UUID) (Session, error) {
	row := q.db.QueryRowContext(ctx, getSession, id)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.RefreshToken,
		&i.UserAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.CrDate,
		&i.ExpiredAt,
	)
	return i, err
}

const getSessionByUserId = `-- name: GetSessionByUserId :one
SELECT id, user_id, refresh_token, user_agent, client_ip, is_blocked, cr_date, expired_at FROM session
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetSessionByUserId(ctx context.Context, userID int64) (Session, error) {
	row := q.db.QueryRowContext(ctx, getSessionByUserId, userID)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.RefreshToken,
		&i.UserAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.CrDate,
		&i.ExpiredAt,
	)
	return i, err
}
