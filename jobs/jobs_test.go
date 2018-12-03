package jobs_test

import (
	"testing"
	"time"

	"github.com/Ibracadabra05/worst_scheduler_ever/jobs"
)

func TestNextRunWorks(t *testing.T) {
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
	validNextRun := time.Date(
		2018,
		12,
		1,
		0,
		1,
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

	if job.NextRun() != validNextRun {
		t.Errorf("Next Run Failed! got: %s, expected: %s", job.NextRun(), validNextRun)
	}
}
