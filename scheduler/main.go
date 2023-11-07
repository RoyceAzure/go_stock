package main

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

var task = func() {
	log.Print("test")
}

func main() {
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
