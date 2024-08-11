package bouncer

import (
	"time"
)

type Config struct {
	ScheduleAt    time.Time
	ScheduleIn    time.Duration
	ScheduleEvery time.Duration
	RetryDelay    time.Duration
	RetriesAmount uint
}

func (t *Task) SetScheduleAt(at time.Time) *Task {
	t.Config.ScheduleAt = at
	return t
}

func (t *Task) SetScheduleIn(in time.Duration) *Task {
	t.Config.ScheduleIn = in
	return t
}

func (t *Task) SetRetryDelay(delay time.Duration) *Task {
	t.Config.RetryDelay = delay
	return t
}

func (t *Task) SetRetriesAmount(amount uint) *Task {
	t.Config.RetriesAmount = amount
	return t
}

func (t *Task) SetScheduleEvery(timeBetween time.Duration) *Task {
	t.Config.ScheduleEvery = timeBetween
	return t
}

type Task struct {
	Func   func() error
	Config Config
}

func NewTask() *Task {
	return new(Task)
}
