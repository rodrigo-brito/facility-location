package neighborhoods

import (
	"github.com/rodrigo-brito/facility-location/model/network"
	"github.com/rodrigo-brito/facility-location/model/solution"
	"github.com/rodrigo-brito/facility-location/util"
	"github.com/rodrigo-brito/facility-location/util/log"
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

func SwapFunctionLocalSearch(data *network.Data, solution *solution.Solution) (newSolution bool) {
	for _, hub := range solution.Hubs {
		for node := 0; node < data.Size; node++ {
			if !solution.Allocation[node][hub] || solution.Allocation[node][node] {
				continue
			}

			tempSolution := solution.GetCopy()
			tempSolution.SwapFunction(node, hub)
			tempSolution.AllocateNearestHub(data) // TODO: test without it

			if tempSolution.GetCost(data) < solution.GetCost(data) {
				newSolution = true
				tempSolution.CopyTo(solution)
				log.Infof("SWAP_FUNCTION: New solution found hubs=%v FO=%.4f", solution.Hubs, solution.GetCost(data))
				break
			}
		}
	}

	log.Infof("Neighborhood swapFunction [%v]", newSolution)

	return
}
