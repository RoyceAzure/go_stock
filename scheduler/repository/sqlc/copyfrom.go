// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: copyfrom.go

package repository

import (
	"context"
)

// iteratorForBulkInsertDAVGALL implements pgx.CopyFromSource.
type iteratorForBulkInsertDAVGALL struct {
	rows                 []BulkInsertDAVGALLParams
	skippedFirstNextCall bool
}

func (r *iteratorForBulkInsertDAVGALL) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForBulkInsertDAVGALL) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].Code,
		r.rows[0].StockName,
		r.rows[0].ClosePrice,
		r.rows[0].MonthlyAvgPrice,
	}, nil
}

func (r iteratorForBulkInsertDAVGALL) Err() error {
	return nil
}

func (q *Queries) BulkInsertDAVGALL(ctx context.Context, arg []BulkInsertDAVGALLParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"stock_day_avg_all"}, []string{"code", "stock_name", "close_price", "monthly_avg_price"}, &iteratorForBulkInsertDAVGALL{rows: arg})
}
