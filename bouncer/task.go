package bouncer

import (
	"time"
)

type Config struct {
	ScheduleAt     time.Time
	ScheduleIn     time.Duration
	RetryDelayMS   int
	RetriesAmount  int
	RetryOnFailure bool
}

func (c *Config) SetScheduleAt(at time.Time) *Config {
	c.ScheduleAt = at
	return c
}

func (c *Config) SetScheduleIn(in time.Duration) *Config {
	c.ScheduleIn = in
	return c
}

func (c *Config) SetRetryDelayMS(delayMS int) *Config {
	c.RetryDelayMS = delayMS
	return c
}

func (c *Config) SetRetriesAmount(amount int) *Config {
	c.RetriesAmount = amount
	return c
}

func (c *Config) SetRetryOnFailure(retry bool) *Config {
	c.RetryOnFailure = retry
	return c
}

type Task struct {
	Func   func() error
	Config Config
}
