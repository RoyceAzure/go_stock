package worker

import (
	"context"

	db "github.com/RoyceAzure/go-stockinfo/repository/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo/service"
	"github.com/RoyceAzure/go-stockinfo/shared/util/mail"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const (
	MailQueue = "mailQueue"
)

/*
asynq有定義 processor 的interface
*/
type TaskProcessor interface {
	Start() error
	StartWithHandler(handler *asynq.ServeMux) error
	ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error
	ProcessTaskBatchUpdateStock(ctx context.Context, task *asynq.Task) error
	ProcessTaskStockTransfer(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessor struct {
	server           *asynq.Server
	store            db.Store
	mailer           mail.EmailSender
	stockTranService service.ITransferService
}

/*
server是processot自己建立的? 是host建立asynq server去遠端redis 拿取任務資料
*/
func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt,
	store db.Store,
	mailer mail.EmailSender,
	stockTranService service.ITransferService) TaskProcessor {
	server := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Queues: map[string]int{
				MailQueue: 10,
			},
			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
				log.Error().
					Err(err).
					Str("type", task.Type()).
					Bytes("body", task.Payload()).
					Msg("process task failed")
			}),
			Logger: NewLoggerAdapter(),
		},
	)

	return &RedisTaskProcessor{
		server:           server,
		store:            store,
		mailer:           mailer,
		stockTranService: stockTranService,
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
	mux.HandleFunc(TaskBatchUpdateStock, processor.ProcessTaskBatchUpdateStock)
	return processor.server.Start(mux)
}
func (processor *RedisTaskProcessor) StartWithHandler(handler *asynq.ServeMux) error {
	return processor.server.Start(handler)
}
