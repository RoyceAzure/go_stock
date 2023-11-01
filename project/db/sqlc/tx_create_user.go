package db

import (
	"context"
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
我要把CreateUser 跟　send email綁再一起嗎?

send email被放到call back裡面

要根據err是不是nil來判斷成功失敗
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

		return arg.AfterCreate(result.User)
	})
	return result, err
}
