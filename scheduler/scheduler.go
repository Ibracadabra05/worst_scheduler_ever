package scheduler

import (
	"fmt"
	"time"

	"github.com/Ibracadabra05/worst_scheduler_ever/jobs"
)

type Scheduler struct {
	jobs []jobs.Job
}

func (s *Scheduler) ScheduleJob(job jobs.Job) {
	fmt.Printf("Scheduled: %+v\n", job)
}

func (s *Scheduler) Run() {
	/*
		1. Check if any of the jobs should run
			a) If now >= NextRun
		2. "Schedule" running of that job
		3. Wait?
	*/

	for {
		for _, job := range s.jobs {
			if job.NextRun().Before(time.Now()) {
				s.ScheduleJob(job)
			}
		}
		time.Sleep(2 * time.Second)
	}
}

func NewScheduler(jobs []jobs.Job) (Scheduler, error) {
	return Scheduler{
		jobs,
	}, nil
}
