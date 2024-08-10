package main

import (
	"log"
	"sync"
	"time"
)

type Bouncer struct {
	wg sync.WaitGroup
}

func NewBouncer() *Bouncer {
	return new(Bouncer)
}

func (b *Bouncer) internalSchedule(task Task) {
	if task.Func == nil {
		log.Fatalln("ERROR: task.Runner is nil")
	}
	if !task.Config.ScheduleAt.IsZero() && task.Config.ScheduleIn > 0 {
		log.Fatalln("ERROR: Cannot set both ScheduleAt and ScheduleIn")
		return
	}

	runFunc := func() {
		defer b.wg.Done()
		task.Func()
	}

	b.wg.Add(1)
	if !task.Config.ScheduleAt.IsZero() {
		waitFor := time.Until(task.Config.ScheduleAt)
		time.AfterFunc(waitFor, runFunc)

	} else if task.Config.ScheduleIn > 0 {
		waitFor := time.Until(time.Now().Add(task.Config.ScheduleIn))
		time.AfterFunc(waitFor, runFunc)

	} else {
		go runFunc()
	}
}

func (b *Bouncer) Schedule(task Task) {
	b.internalSchedule(task)
	b.wg.Wait()
}

func (b *Bouncer) ScheduleMultiple(tasks []Task) {
	for _, task := range tasks {
		b.internalSchedule(task)
	}

	b.wg.Wait()
}
