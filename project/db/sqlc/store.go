package db

import (
	"context"
	"database/sql"
	"fmt"
)

/*
New(db *sql.DB) 會把db放到Queries.db

	type Queries struct {
		db                                DBTX
		tx                                *sql.Tx
		...
*/

// 製作一個Repo介面  可以用來做Mock
type Store interface {
	Querier
	TransferStockTx(ctx context.Context, arg TransferStockTxParams) (TransferStockTxResults, error)
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResults, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

// Queries struct 本身就有包含 *sql.DB
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// 注意最後是return tx.Commit() 就表示有可能在commit時也會有error
// 一個克制化的通用trans func
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	//又創建了一個Queries 差別在於裡面的*sql.DB 是個Tx
	//因為封裝Repo操作都在Query裡面，所以建立玩trans還必須重新丟回Repo做操作
	q := New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err : %v", err, rbErr)
		}
	}

	return tx.Commit()
}
