# BOUNCER
Simplistic golang scheduler library

## Usage
```go
func main() {
    b := bouncer.New()

    taskFunc := func() error {
        fmt.Println("You're in")
        return nil
    }

    taskIn := bouncer.Task{
        Func: taskFunc,
        Config: bouncer.Config{
            ScheduleIn: time.Duration(5 * time.Second),
        },
    }

    b.Schedule(taskIn)

    taskAt := bouncer.Task{
        Func: taskFunc,
        Config: bouncer.Config{
            // Will be scheduled in 10 hours from now but it can be pure time.Time struct
            ScheduleAt: time.Now().Add(10 * time.Hour), 
            RetriesAmount: 5,
            RetryDelay: time.Second * 5,
        },
    }

    // Also you can use builder
    taskEvery := bouncer.NewTask().SetScheduleEvery(time.Hour * 12)
    b.ScheduleMultiple([]bouncer.Task{taskAt, taskEvery}) 

    // After that all code is waiting for all tasks to be completed
    // no more tasks can be scheduled after this point
    b.Wait()
}
```

### Available Methods:
- ```Schedule(bouncer.Task)```
- ```ScheduleMultiple([]bouncer.Task)```
- ```Wait()``` - blocks the app until all tasks are finished

## Support
If you have encountered a bug or have any ideas how to improve this library - don't be afraid to open an issue with an explanation.
