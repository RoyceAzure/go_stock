package worker

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const TaskSendVerifyEmail = "task:send_verify_email"

type PayloadSendVerifyEmail struct {
	UserName string `json:"user_name"`
	UserId   int64  `json:"user_id"`
}

/*
製作task  這裡把所有資訊都包進task裡面  包括retry delay 甚至是要使用甚麼優先級的queue

使用client enqueue task
*/
func (distributor *RedisTaskDistributor) DisstributeTaskSendVerifyEmail(
	ctx context.Context,
	payload *PayloadSendVerifyEmail,
	opts ...asynq.Option,
) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload %w", err)
	}

	//沒有提到要如何執行task , 只有給payload跟設定
	task := asynq.NewTask(TaskSendVerifyEmail, jsonData, opts...)
	taskInfo, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task %w", err)
	}

	log.Info().
		Str("type", task.Type()).
		Bytes("body", task.Payload()).
		Str("queue", taskInfo.Queue).
		Int("max_retry", taskInfo.MaxRetry).
		Msg("enqueued task")
	return nil
}

/*
因為這個func 簽名(ctx context.Context, task *asynq.Task) error 是asynq定義
所以asynq內部會根據error內容作相應處理

如果回傳得error 有wrap asynq.SkipRetry，則不會執行retry
*/
func (processor *RedisTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendVerifyEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal task payload %w", asynq.SkipRetry)
	}
	user, err := processor.store.GetUser(ctx, payload.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user doesn't exists %w", asynq.SkipRetry)
		}
		return fmt.Errorf("failed to get usesr %w", err)
	}

	//TODO : send email to user
	log.Info().Str("type", task.Type()).
		Bytes("task payload", task.Payload()).
		Str("user email", user.Email).
		Msg("porcessed task")
	return nil
}
