package bouncer

import (
	"log"
	"sync"
	"time"
)

type Bouncer struct {
	sealed bool
	wg     sync.WaitGroup
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

		for retries < task.Config.RetriesAmount || task.Config.ScheduleEvery > 0 || firstRun {
			defer b.wg.Done()
			err := task.Func()

			if err == nil {
				if task.Config.ScheduleEvery > 0 {
					time.Sleep(task.Config.ScheduleEvery)
					retries = 0
					continue
				}

				// This part is called when task is not scheduled every X but we still got success
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
	if b.sealed {
		log.Fatalln("Cannot add new tasks after Wait method")
	}

	b.internalSchedule(task)
}

func (b *Bouncer) ScheduleMultiple(tasks []Task) {
	if b.sealed {
		log.Fatalln("Cannot add new tasks after Wait method")
	}

	for _, task := range tasks {
		b.internalSchedule(task)
	}
}
func (b *Bouncer) Wait() {
	b.sealed = true
	b.wg.Wait()
}
