package bouncer

import (
	"time"
)

type Config struct {
	ScheduleAt    time.Time
	ScheduleIn    time.Duration
	RetryDelayMS  uint
	RetriesAmount uint
}

func (c *Config) SetScheduleAt(at time.Time) *Config {
	c.ScheduleAt = at
	return c
}

func (c *Config) SetScheduleIn(in time.Duration) *Config {
	c.ScheduleIn = in
	return c
}

func (c *Config) SetRetryDelayMS(delayMS uint) *Config {
	c.RetryDelayMS = delayMS
	return c
}

func (c *Config) SetRetriesAmount(amount uint) *Config {
	c.RetriesAmount = amount
	return c
}

type Task struct {
	Func   func() error
	Config Config
}
