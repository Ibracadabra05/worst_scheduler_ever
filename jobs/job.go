package jobs

import "time"

type Job struct {
	command string
	rate    time.Duration
	lastRun time.Time
}

func (j *Job) Command() string {
	return j.command
}

func (j *Job) Rate() time.Duration {
	return j.rate
}

func (j *Job) LastRun() time.Time {
	return j.lastRun
}

func (j *Job) NextRun() time.Time {
	return j.lastRun.Add(j.rate)
}

func NewJob(command string, rate time.Duration, lastRun time.Time) (Job, error) {
	return Job{
		command,
		rate,
		lastRun,
	}, nil
}
