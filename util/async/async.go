package async

import (
	"sync"

	"github.com/rodrigo-brito/facility-location/model/network"
	"github.com/rodrigo-brito/facility-location/model/solution"
	"github.com/rodrigo-brito/facility-location/util/log"
)

type Task func(data *network.Data, solution *solution.Solution)

func taskConsumer(tasks chan Task, wg *sync.WaitGroup, data *network.Data, solution *solution.Solution) {
	for {
		task, more := <-tasks
		if !more {
			return
		}

		task(data, solution)
		log.Info("done")
		wg.Done()
	}
}

func Run(data *network.Data, solution *solution.Solution, maxAsyncTasks int, tasks ...Task) {
	wg := new(sync.WaitGroup)
	tasksChannel := make(chan Task)

	for i := 0; i < maxAsyncTasks; i++ {
		go taskConsumer(tasksChannel, wg, data, solution)
	}

	for i := 0; i < len(tasks); i++ {
		wg.Add(1)
		tasksChannel <- tasks[i]
	}

	close(tasksChannel)
	wg.Wait()
}
