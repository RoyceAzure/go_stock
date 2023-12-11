// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateFund(ctx context.Context, arg CreateFundParams) (Fund, error)
	CreateRealizedProfitLoss(ctx context.Context, arg CreateRealizedProfitLossParams) (RealizedProfitLoss, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateStock(ctx context.Context, arg CreateStockParams) (Stock, error)
	CreateStockTransaction(ctx context.Context, arg CreateStockTransactionParams) (StockTransaction, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateUserStock(ctx context.Context, arg CreateUserStockParams) (UserStock, error)
	CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmailParams) (VerifyEmail, error)
	DeleteFund(ctx context.Context, fundID int64) error
	DeleteStock(ctx context.Context, stockID int64) error
	DeleteStockTransaction(ctx context.Context, transationID int64) error
	DeleteUser(ctx context.Context, userID int64) error
	DeleteUserStock(ctx context.Context, userStockID int64) error
	GetFund(ctx context.Context, fundID int64) (Fund, error)
	GetFundByUidandCurForUpdateNoK(ctx context.Context, arg GetFundByUidandCurForUpdateNoKParams) (Fund, error)
	GetFundByUidandFid(ctx context.Context, arg GetFundByUidandFidParams) (Fund, error)
	GetFundByUidandFidForUpdateNoK(ctx context.Context, arg GetFundByUidandFidForUpdateNoKParams) (Fund, error)
	GetFundByUserId(ctx context.Context, arg GetFundByUserIdParams) ([]Fund, error)
	GetFunds(ctx context.Context, arg GetFundsParams) ([]Fund, error)
	GetRealizedProfitLoss(ctx context.Context, id int64) (RealizedProfitLoss, error)
	GetRealizedProfitLosssByUserId(ctx context.Context, arg GetRealizedProfitLosssByUserIdParams) ([]RealizedProfitLoss, error)
	GetRealizedProfitLosssByUserIdDetial(ctx context.Context, arg GetRealizedProfitLosssByUserIdDetialParams) ([]GetRealizedProfitLosssByUserIdDetialRow, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetStock(ctx context.Context, stockID int64) (Stock, error)
	GetStockByCN(ctx context.Context, stockName string) (Stock, error)
	GetStockByCode(ctx context.Context, stockCode string) (Stock, error)
	GetStockForUpdate(ctx context.Context, stockID int64) (Stock, error)
	GetStockTransaction(ctx context.Context, transationID int64) (StockTransaction, error)
	GetStockTransactionsByDate(ctx context.Context, arg GetStockTransactionsByDateParams) ([]StockTransaction, error)
	GetStockTransactionsByStockId(ctx context.Context, arg GetStockTransactionsByStockIdParams) ([]StockTransaction, error)
	GetStockTransactionsByUserId(ctx context.Context, arg GetStockTransactionsByUserIdParams) ([]StockTransaction, error)
	GetStockTransactionsFilter(ctx context.Context, arg GetStockTransactionsFilterParams) ([]StockTransaction, error)
	GetStocks(ctx context.Context, arg GetStocksParams) ([]Stock, error)
	GetUser(ctx context.Context, userID int64) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserForUpdateNoKey(ctx context.Context, userID int64) (User, error)
	GetUserStock(ctx context.Context, userStockID int64) (UserStock, error)
	GetUserStockByUidandSidForUpdateNoK(ctx context.Context, arg GetUserStockByUidandSidForUpdateNoKParams) (UserStock, error)
	GetUserStocks(ctx context.Context, arg GetUserStocksParams) ([]UserStock, error)
	GetUserStocksByPDate(ctx context.Context, arg GetUserStocksByPDateParams) ([]UserStock, error)
	GetUserStocksByStockId(ctx context.Context, arg GetUserStocksByStockIdParams) ([]GetUserStocksByStockIdRow, error)
	GetUserStocksByUserAStock(ctx context.Context, arg GetUserStocksByUserAStockParams) ([]UserStock, error)
	GetUserStocksByUserId(ctx context.Context, arg GetUserStocksByUserIdParams) ([]GetUserStocksByUserIdRow, error)
	GetserStockByUidandSid(ctx context.Context, arg GetserStockByUidandSidParams) (UserStock, error)
	GetstockByTS(ctx context.Context, arg GetstockByTSParams) ([]Stock, error)
	Getusers(ctx context.Context, arg GetusersParams) ([]User, error)
	UpdateFund(ctx context.Context, arg UpdateFundParams) (Fund, error)
	UpdateStock(ctx context.Context, arg UpdateStockParams) (Stock, error)
	UpdateStockCPByCode(ctx context.Context, arg UpdateStockCPByCodeParams) (Stock, error)
	UpdateStockTransationResult(ctx context.Context, arg UpdateStockTransationResultParams) (StockTransaction, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateUserStock(ctx context.Context, arg UpdateUserStockParams) (UserStock, error)
	UpdateVerifyEmail(ctx context.Context, arg UpdateVerifyEmailParams) (VerifyEmail, error)
}

var _ Querier = (*Queries)(nil)
