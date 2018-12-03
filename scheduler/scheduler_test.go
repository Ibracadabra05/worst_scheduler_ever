package scheduler_test

import (
	"testing"
	"time"

	"github.com/Ibracadabra05/worst_scheduler_ever/jobs"
	"github.com/Ibracadabra05/worst_scheduler_ever/scheduler"
)

func TestRun(t *testing.T) {
	utc, _ := time.LoadLocation("utc")
	validLastRun := time.Date(
		2018,
		12,
		1,
		0,
		0,
		0,
		0,
		utc,
	)

	duration, err := time.ParseDuration("1m")

	if err != nil {
		t.Error(err)
	}

	job, err := jobs.NewJob(
		"lakdljahskdfjhaflskdfj",
		duration,
		validLastRun,
	)
	if err != nil {
		t.Error(err)
	}

	scheduler, err := scheduler.NewScheduler(
		[]jobs.Job{job},
	)

	scheduler.Run()
}
