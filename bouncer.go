package bouncer

import (
	"log"
	"sync"
	"time"
)

type Bouncer struct {
	wg    sync.WaitGroup
	tasks []Task
}

func New() *Bouncer {
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
		var retries uint = 0
		var firstRun bool = true

		for retries < task.Config.RetriesAmount || firstRun {
			defer b.wg.Done()
			err := task.Func()

			if err == nil {
				break
			}

			if task.Config.RetryDelay > 0 {
				time.Sleep(task.Config.RetryDelay)
			}

			retries++
			firstRun = false
		}
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
	b.tasks = append(b.tasks, task)
}

func (b *Bouncer) ScheduleMultiple(tasks []Task) {
	b.tasks = append(b.tasks, tasks...)
}

func (b *Bouncer) Run() {
	for _, task := range b.tasks {
		b.internalSchedule(task)
	}

	b.wg.Wait()
}
