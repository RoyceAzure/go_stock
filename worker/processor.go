package worker

import (
	"context"

	db "github.com/RoyceAzure/go-stockinfo-project/db/sqlc"
	"github.com/hibiken/asynq"
)

const (
	QueueCritical = "critical"
	QueueDefault  = "default"
)

/*
asynq有定義 processor 的interface
*/
type TaskProcessor interface {
	Start() error
	ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessor struct {
	server *asynq.Server
	store  db.Store
}

/*
server是processot自己建立的? 應該不是
因該是一個遠端host 已經有redis server相關組件並且已經啟動且監聽某個port
這裡只是建立連線與設定才對
*/
func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, store db.Store) TaskProcessor {
	server := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Queues: map[string]int{
				QueueCritical: 10,
				QueueDefault:  5,
			},
		},
	)

	return &RedisTaskProcessor{
		server: server,
		store:  store,
	}
}

/*
案照這個設計思維   所有processor應可以連到同一個redis server
且每個processor start 就是指start一個路由與handler的對應  且會永久執行?  對  就跟http server很像
一旦執行後就會被block  所以要用go routine執行
asynq 本身processor處理每個task就是使用go routine處理
error是指 sersver startup 過程的錯誤
*/
func (processor *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()
	mux.HandleFunc(TaskSendVerifyEmail, processor.ProcessTaskSendVerifyEmail)

	return processor.server.Start(mux)
}
