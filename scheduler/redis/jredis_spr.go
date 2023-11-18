package jredis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	repository "github.com/RoyceAzure/go-stockinfo-schduler/repository/sqlc"
	"github.com/go-redis/redis"
)

func (jredis *Jredis) BulkInsertSPR(ctx context.Context, key string, sprs []repository.StockPriceRealtime) error {
	data, err := json.Marshal(sprs)
	if err != nil {
		return fmt.Errorf("failed to encode spr : %w", err)
	}
	//如果key不存在，这个命令将设置key 的值为value。如果 key 已存在，命令不会做任何改动，并返回 0
	//改使用SET 因為資料不會逐條更新，要修改直接全部覆蓋
	res := jredis.client.Set(ctx, key, string(data), 0)
	if err := res.Err(); err != nil {
		return fmt.Errorf("failed to insert spr to redis : %w", err)
	}
	return nil
}

func (jredis *Jredis) FindSPRByID(ctx context.Context, key string) ([]repository.StockPriceRealtime, error) {
	strCmd := jredis.client.Get(ctx, key)
	if err := strCmd.Err(); errors.Is(err, redis.Nil) {
		return nil, ErrNotExists
	} else if err != nil {
		return nil, fmt.Errorf("find spr get err : %w", err)
	}
	var result []repository.StockPriceRealtime
	err := json.Unmarshal([]byte(strCmd.Val()), &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal spr, err : %w", err)
	}
	return result, nil
}

func (jredis *Jredis) DeleteSPRByID(ctx context.Context, key string) error {

	txn := jredis.client.TxPipeline()

	strCmd := txn.Del(ctx, key)
	if err := strCmd.Err(); errors.Is(err, redis.Nil) {
		txn.Discard()
		return ErrNotExists
	} else if err != nil {
		txn.Discard()
		return fmt.Errorf("delete spr get err : %w", err)
	}

	if _, err := txn.Exec(ctx); err != nil {
		return fmt.Errorf("delete spr get err : %w", err)
	}

	return nil
}
