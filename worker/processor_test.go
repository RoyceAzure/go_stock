package worker

import (
	"sync"
	"testing"

	"github.com/hibiken/asynq"
)

/*
這可能需要一個goroutine執行，那要怎知道err???
*/
func TestIntergrateBatchUpdateStock(t *testing.T) {
	mux := asynq.NewServeMux()
	mux.HandleFunc(TaskBatchUpdateStock, testWorker.ProcessTaskBatchUpdateStock)
	var wg sync.WaitGroup
	wg.Add(1)
	go testWorker.StartWithHandler(mux)
	wg.Wait()
}
