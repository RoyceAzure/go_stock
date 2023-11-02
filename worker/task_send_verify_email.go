package worker

import (
	"context"
	"encoding/json"
	"fmt"

	db "github.com/RoyceAzure/go-stockinfo-project/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo-shared/utility"
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
func (distributor *RedisTaskDistributor) DistributeTaskSendVerifyEmail(
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

注意這段handler是在redis server內部執行  所以它回傳的err要額外設置  我們才能看得見

err == sql.ErrNORows 有可能是db 還沒有完成commit，所以就讓他retry
*/
func (processor *RedisTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendVerifyEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal task payload %w", asynq.SkipRetry)
	}
	user, err := processor.store.GetUser(ctx, payload.UserId)
	if err != nil {
		// if err == sql.ErrNoRows {
		// 	return fmt.Errorf("user doesn't exists %w", asynq.SkipRetry)
		// }
		return fmt.Errorf("failed to get usesr %w", err)
	}
	verifyEmail, err := processor.store.CreateVerifyEmail(ctx, db.CreateVerifyEmailParams{
		UserID:     user.UserID,
		Email:      user.Email,
		SecretCode: utility.RandomString(32),
	})
	if err != nil {
		return fmt.Errorf("failed to create verify email %w", err)
	}

	//TODO : send email to user
	subject := "Welcome to StockInfo"
	verifyURL := fmt.Sprintf("http://localhost:8080/v1/verify_email?email_id=%d&secret_code=%s",
		verifyEmail.ID, verifyEmail.SecretCode)

	content := fmt.Sprintf(`Hello %s, <br/>
	Thank you for registering <br/>
	Please <a href="%s">click here</a> to verify your email address.<br/>
	`, user.UserName, verifyURL)

	to := []string{user.Email}
	bcc := []string{"roycewnag@gmail.com"}

	err = processor.mailer.SendEmail(subject, content, to, nil, bcc, nil)
	if err != nil {
		return fmt.Errorf("failed to create verify email %w", err)
	}

	log.Info().Str("type", task.Type()).
		Bytes("task payload", task.Payload()).
		Str("user email", user.Email).
		Msg("porcessed task")
	return nil
}
