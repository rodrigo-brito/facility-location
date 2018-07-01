package async

import (
	"sync"

	"github.com/rodrigo-brito/facility-location/model/network"
	"github.com/rodrigo-brito/facility-location/model/solution"
)

type Task func(data *network.Data, solution *solution.Solution)

func taskConsumer(tasks chan Task, wg *sync.WaitGroup, data *network.Data, solution *solution.Solution) {
	for {
		task, more := <-tasks
		if !more {
			return
		}

		task(data, solution)
		wg.Done()
	}
}

func Run(data *network.Data, solution *solution.Solution, maxAsyncTasks int, tasks ...Task) {
	wg := new(sync.WaitGroup)

	//for _, task := range tasks {
	//	wg.Add(1)
	//	go func() {
	//		task(data, solution)
	//		wg.Done()
	//	}()
	//}
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
