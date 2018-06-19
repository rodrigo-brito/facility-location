package neighborhoods

import (
	"github.com/rodrigo-brito/facility-location/model/network"
	"github.com/rodrigo-brito/facility-location/model/solution"
	"github.com/rodrigo-brito/facility-location/util"
	"github.com/rodrigo-brito/facility-location/util/async"
)

func ShiftPerturbation(solution *solution.Solution) {
	// invalid for star network and mash
	if len(solution.Hubs) == solution.Size || len(solution.Hubs) == 1 {
		return
	}

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

	// generate list of hub candidates
	hubs := make([]int, 0)
	for _, hub := range solution.Hubs {
		if hub != solution.AllocationNode[node] {
			hubs = append(hubs, hub)
		}
	}

	// select random hub
	indexHub := util.Random(0, len(hubs)-1)
	newHub := hubs[indexHub]

	// make allocation
	solution.AllocNode(node, newHub)
}

func ShiftLocalSearch(data *network.Data, bestSolution *solution.Solution) (newSolution bool) {
	tasks := make([]async.Task, 0)
	updatedChannel := make(chan bool, data.Size*len(bestSolution.Hubs))

	for i := 0; i < data.Size; i++ {
		node := i
		tasks = append(tasks, func(data *network.Data, solution *solution.Solution) {
			for _, hub := range bestSolution.Hubs {
				tempSolution := solution.GetCopy()
				if tempSolution.Allocation[node][hub] || tempSolution.Allocation[node][node] {
					continue
				}

				tempSolution.AllocNode(node, hub)
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
