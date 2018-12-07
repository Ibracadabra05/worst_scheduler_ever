package scheduler

import (
	"errors"
	"fmt"
	"time"

	"github.com/Ibracadabra05/worst_scheduler_ever/jobs"
)

type Scheduler interface {
	Run()
	ScheduleJob(job jobs.Job)
}

type SyncScheduler struct {
	jobs []jobs.Job
}

func (s SyncScheduler) ScheduleJob(job jobs.Job) {
	fmt.Printf("Scheduled: %+v\n", job)
}

func (s SyncScheduler) Run() {
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

func (s SyncScheduler) AddJobs(jobs []jobs.Job) {
	s.jobs = append(s.jobs, jobs...)
}

func SchedulerFactory(schedulerType string) (Scheduler, error) {
	if schedulerType == "sync" {
		return &SyncScheduler{
			[]jobs.Job{},
		}, nil
	}
	return nil, errors.New(fmt.Sprintf("Invalid schedulerType: %s", schedulerType))
}
