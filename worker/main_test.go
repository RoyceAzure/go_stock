package worker

import (
	"database/sql"
	"os"
	"testing"

	db "github.com/RoyceAzure/go-stockinfo-project/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo-schduler/util/config"
	"github.com/RoyceAzure/go-stockinfo-shared/utility/mail"
	"github.com/hibiken/asynq"
)

var testWorker TaskProcessor

func TestMain(m *testing.M) {
	setUp()
	os.Exit(m.Run())
}

func setUp() {
	config, _ := config.LoadConfig("../")
	conn, _ := sql.Open(config.DBDriver, config.DBSource)
	store := db.NewStore(conn)
	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisQueueAddress,
	}
	testWorker = NewTestRedisTaskProcessor(redisOpt, store, nil)
}

func NewTestRedisTaskProcessor(redisOpt asynq.RedisClientOpt, store db.Store, mailer mail.EmailSender) TaskProcessor {
	server := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Queues: map[string]int{
				QueueCritical: 10,
				QueueDefault:  5,
			},
			Logger: NewLoggerAdapter(),
		},
	)

	return &RedisTaskProcessor{
		server: server,
		store:  store,
		mailer: mailer,
	}
}
