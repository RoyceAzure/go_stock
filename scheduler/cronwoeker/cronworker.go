package cronworker

import (
	"context"
	"time"

	"github.com/RoyceAzure/go-stockinfo-schduler/service"
	"github.com/go-co-op/gocron"
)

type CornWorker interface {
	Start()
	StopAsync()
	SetUpSchdulerWorker(ctx context.Context)
}

type schdulerCronWorker struct {
	service service.Service
	Cron    *gocron.Scheduler
}

func NewSchdulerWorker(service service.Service) *schdulerCronWorker {
	s := gocron.NewScheduler(time.UTC)

	return &schdulerCronWorker{
		service: service,
		Cron:    s,
	}
}

func (cornWorker *schdulerCronWorker) SetUpSchdulerWorker(ctx context.Context) {
	cornWorker.Cron.CronWithSeconds("*/5 * * * * *").Do(cornWorker.service.DownloadAndInsertDataSVAA, ctx)
}

/*
Blocking start
*/
func (cornWorker *schdulerCronWorker) Start() {
	cornWorker.Cron.StartBlocking()
}

func (cornWorker *schdulerCronWorker) StopAsync() {
	cornWorker.Cron.Stop()
}
