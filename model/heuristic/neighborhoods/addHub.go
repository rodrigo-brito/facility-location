package neighborhoods

import (
	"github.com/rodrigo-brito/hub-spoke-go/model/network"
	"github.com/rodrigo-brito/hub-spoke-go/model/solution"
	"github.com/rodrigo-brito/hub-spoke-go/util"
	"github.com/rodrigo-brito/hub-spoke-go/util/log"
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

func AddHubLocalSearch(data *network.Data, solution *solution.Solution) (newSolution bool) {
	for node := 0; node < data.Size; node++ {
		if solution.Allocation[node][node] {
			continue
		}

		tempSolution := solution.GetCopy()
		tempSolution.AddHub(node)
		tempSolution.AllocateNearestHub(data)

		if tempSolution.GetCost(data) < solution.GetCost(data) {
			newSolution = true
			tempSolution.CopyTo(solution)
			log.Infof("ADD_HUB: New solution found FO=%.4f hubs=%v", solution.GetCost(data), solution.Hubs)
		}

	}

	log.Infof("Neighborhood addHub [%v]", newSolution)

	return
}
