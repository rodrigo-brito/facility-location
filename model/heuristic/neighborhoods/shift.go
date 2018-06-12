package neighborhoods

import (
	"github.com/rodrigo-brito/facility-location/model/network"
	"github.com/rodrigo-brito/facility-location/model/solution"
	"github.com/rodrigo-brito/facility-location/util"
	"github.com/rodrigo-brito/facility-location/util/log"
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

func ShiftLocalSearch(data *network.Data, solution *solution.Solution) (newSolution bool) {
	for node := 0; node < data.Size; node++ {
		for _, hub := range solution.Hubs {
			if solution.Allocation[node][hub] || solution.Allocation[node][node] {
				continue
			}

			tempSolution := solution.GetCopy()
			tempSolution.AllocNode(node, hub)

			if tempSolution.GetCost(data) < solution.GetCost(data) {
				newSolution = true
				tempSolution.CopyTo(solution)
				log.Infof("SHIFT: New solution found FO=%.4f  hubs=%v", solution.GetCost(data), solution.Hubs)
			}
		}
	}

	log.Infof("Neighborhood shift [%v]", newSolution)

	return
}
