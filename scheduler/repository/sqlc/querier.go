// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package repository

import (
	"context"
)

type Querier interface {
	BatchDeleteSDAVGALL(ctx context.Context, dollar_1 []int64) error
	BulkInsertDAVGALL(ctx context.Context, arg []BulkInsertDAVGALLParams) (int64, error)
	BulkInsertSPR(ctx context.Context, arg []BulkInsertSPRParams) (int64, error)
	CreateSDAVGALL(ctx context.Context, arg CreateSDAVGALLParams) (StockDayAvgAll, error)
	CreateSPR(ctx context.Context, arg CreateSPRParams) (StockPriceRealtime, error)
	DeleteSDAVGALLCodePrexForTest(ctx context.Context, arg DeleteSDAVGALLCodePrexForTestParams) error
	GetSDAVGALLs(ctx context.Context, arg GetSDAVGALLsParams) ([]StockDayAvgAll, error)
	GetSPRs(ctx context.Context, arg GetSPRsParams) ([]StockPriceRealtime, error)
}

var _ Querier = (*Queries)(nil)