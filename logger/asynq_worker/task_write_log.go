package asynq_worker

import (
	"context"
	"encoding/json"
	"fmt"

	repository "github.com/RoyceAzure/go-stockinfo-logger/repository/mongodb"
	"github.com/hibiken/asynq"
)

const TaskWriteLog = "task:write_log"

func (taskProcesser *WriteLogProcessor) WriteLog(ctx context.Context, task *asynq.Task) error {
	var payload repository.LogEntry

	if taskProcesser.mongoDao == nil {
		return fmt.Errorf("mongo dao has not init %w", asynq.SkipRetry)
	}

	err := json.Unmarshal(task.Payload(), &payload)
	if err != nil {
		return fmt.Errorf("failed to unmarshal task payload %w", asynq.SkipRetry)
	}

	err = taskProcesser.mongoDao.Insert(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}
