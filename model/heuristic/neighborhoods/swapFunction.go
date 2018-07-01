package neighborhoods

import (
	"github.com/rodrigo-brito/facility-location/model/network"
	"github.com/rodrigo-brito/facility-location/model/solution"
	"github.com/rodrigo-brito/facility-location/util"
	"github.com/rodrigo-brito/facility-location/util/async"
)

func SwapFunctionPerturbation(solution *solution.Solution) {
	// generate list of nodes candidates
	nodes := make([]int, 0)
	for i := 0; i < solution.Size; i++ {
		if !solution.Allocation[i][i] {
			nodes = append(nodes, i)
		}
	}

	// select random node
	nodeIndex := util.Random(0, len(nodes)-1)
	node := nodes[nodeIndex]

	// make swap
	solution.SwapFunction(node, solution.AllocationNode[node])
}

func SwapFunctionLocalSearch(data *network.Data, bestSolution *solution.Solution) (newSolution bool) {
	tasks := make([]async.Task, 0)
	updatedChannel := make(chan bool, len(bestSolution.Hubs)*data.Size)

	for _, i := range bestSolution.Hubs {
		hub := i
		tasks = append(tasks, func(data *network.Data, solution *solution.Solution) {
			for node := 0; node < data.Size; node++ {
				tempSolution := solution.GetCopy()

				if !tempSolution.Allocation[node][hub] || tempSolution.Allocation[node][node] {
					continue
				}

				tempSolution.SwapFunction(node, hub)
				ShiftLocalSearch(data, tempSolution)

				updatedChannel <- solution.UpdateIfBetter(tempSolution, data)
			}
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
