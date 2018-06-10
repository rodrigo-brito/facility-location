package neighborhoods

import (
	"github.com/rodrigo-brito/hub-spoke-go/model/network"
	"github.com/rodrigo-brito/hub-spoke-go/model/solution"
	"github.com/rodrigo-brito/hub-spoke-go/util"
	"github.com/rodrigo-brito/hub-spoke-go/util/log"
)

func RemoveHubPerturbation(solution *solution.Solution) {
	if len(solution.Hubs) > 1 {
		indexHub := util.Random(0, len(solution.Hubs)-1)
		solution.RemoveHub(solution.Hubs[indexHub])
	}
}

func RemoveHubLocalSearch(data *network.Data, solution *solution.Solution) (newSolution bool) {
	for _, hub := range solution.Hubs {
		if len(solution.Hubs) < 2 {
			return
		}

		tempSolution := solution.GetCopy()
		tempSolution.RemoveHub(hub)
		tempSolution.AllocateNearestHub(data)

		if tempSolution.GetCost(data) < solution.GetCost(data) {
			newSolution = true
			tempSolution.CopyTo(solution)
			log.Infof("REMOVE_HUB: New solution found FO=%.4f hubs=%v", solution.GetCost(data), solution.Hubs)
		}

	}

	log.Infof("Neighborhood removeHub [%v]", newSolution)

	return
}
