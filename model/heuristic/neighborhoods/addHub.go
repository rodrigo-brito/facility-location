package neighborhoods

import (
	"github.com/rodrigo-brito/facility-location/model/network"
	"github.com/rodrigo-brito/facility-location/model/solution"
	"github.com/rodrigo-brito/facility-location/util"
	"github.com/rodrigo-brito/facility-location/util/async"
)

func AddHubPerturbation(solution *solution.Solution) {
	if len(solution.Hubs) < solution.Size {
		nodes := make([]int, 0)

		for node := 0; node < solution.Size; node++ {
			if !solution.Allocation[node][node] {
				nodes = append(nodes, node)
			}
		}

		indexNode := util.Random(0, len(nodes)-1)
		solution.AddHub(nodes[indexNode])
	}
}

func AddHubLocalSearch(data *network.Data, bestSolution *solution.Solution) bool {
	tasks := make([]async.Task, 0)
	updatedChannel := make(chan bool, data.Size)

	for i := 0; i < data.Size; i++ {
		node := i

		if bestSolution.Allocation[node][node] {
			continue
		}

		tasks = append(tasks, func(data *network.Data, solution *solution.Solution) {
			tempSolution := solution.GetCopy()
			tempSolution.AddHub(node)
			ShiftLocalSearch(data, tempSolution)

			updatedChannel <- solution.UpdateIfBetter(tempSolution, data)
		})
	}

	async.Run(data, bestSolution, data.MaxAsyncTask, tasks...)
	close(updatedChannel)

	for ok := range updatedChannel {
		if ok {
			return true
		}
	}
	return false
}
