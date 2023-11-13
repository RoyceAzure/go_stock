package main

import (
	"context"
	"os"
	"time"

	"github.com/RoyceAzure/go-stockinfo-schduler/api"
	repository "github.com/RoyceAzure/go-stockinfo-schduler/repository/sqlc"
	"github.com/RoyceAzure/go-stockinfo-schduler/service"
	"github.com/RoyceAzure/go-stockinfo-schduler/util/config"
	"github.com/go-co-op/gocron"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var task = func() {
	log.Print("test")
}

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot load config")
	}
	if config.Enviornmant == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx := context.Background()
	pgxPool, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("err create db connect")
	}
	dao := repository.NewSQLDao(pgxPool)

	service := service.NewService(dao)
	go runGinServer(config, dao, service)

	s := gocron.NewScheduler(time.Local)

	// Every starts the job immediately and then runs at the
	// specified interval
	// job, err := s.Every(5).Seconds().Do(task)
	// if err != nil {
	// 	// handle the error related to setting up the job
	// }

	// to wait for the interval to pass before running the first job
	// use WaitForSchedule or WaitForScheduleAll
	// s.Every(5).Second().WaitForSchedule().Do(task)

	s.WaitForScheduleAll()
	// s.Every(5).Second().Do(task) // waits for schedule
	// s.Every(5).Second().Do(task) // waits for schedule

	// // strings parse to duration
	// s.Every("5m").Do(task)

	// s.Every(5).Days().Do(task)

	// s.Every(1).Month(1, 2, 3).Do(task)

	// // set time
	// s.Every(1).Day().At("10:30").Do(task)

	// // set multiple times
	// s.Every(1).Day().At("10:30;08:00").Do(task)

	// s.Every(1).Day().At("10:30").At("08:00").Do(task)

	// // Schedule each last day of the month
	// s.Every(1).MonthLastDay().Do(task)

	// // Or each last day of every other month
	// s.Every(2).MonthLastDay().Do(task)

	// cron expressions supported
	s.Cron("* * * * *").Do(task) // every minute

	// cron second-level expressions supported
	s.CronWithSeconds("*/3 * 18 * * *").Do(task) // every second

	// you can start running the scheduler in two different ways:
	// starts the scheduler asynchronously
	// s.StartAsync()
	// starts the scheduler and blocks current execution path
	s.StartBlocking()

	// stop the running scheduler in two different ways:
	// stop the scheduler
	s.Stop()

	// stop the scheduler and notify the `StartBlocking()` to exit
	s.StopBlockingChan()
}

/*
build Server and run
*/
func runGinServer(configs config.Config, dao repository.Dao, service service.SyncDataService) {
	server, err := api.NewServer(configs, dao, service)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot start server")
	}
	err = server.Start(configs.HttpServerAddress)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("cannot start server")
	}
}
