package db

import (
	"context"

	"github.com/RoyceAzure/go-stockinfo/shared/util/constants"
)

/*
AfterCreate is call back, use it err to check commit or rollback
*/
type CreateUserTxParams struct {
	CreateUserParams
	AfterCreate func(user User) error
}

type CreateUserTxResults struct {
	User User
}

/*
如果遇到db高流量情況，trsation會delay, 所以async操作需要設置delay
*/
func (store *SQLStore) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResults, error) {
	var result CreateUserTxResults

	//
	//後續 這個 func(q *Queries)  應該就使已經寫好且沒有Tx的版本
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.User, err = q.CreateUser(ctx, arg.CreateUserParams)
		if err != nil {
			return err
		}

		q.CreateFund(ctx, CreateFundParams{
			UserID:       result.User.UserID,
			Balance:      "0.00",
			CurrencyType: string(constants.TW),
			CrUser:       "SYSTEM",
		})

		return arg.AfterCreate(result.User)
	})
	return result, err
}
