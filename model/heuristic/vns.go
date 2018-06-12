package heuristic

import (
	"github.com/rodrigo-brito/facility-location/model/heuristic/neighborhoods"
	"github.com/rodrigo-brito/facility-location/model/network"
	"github.com/rodrigo-brito/facility-location/model/solution"
	"github.com/rodrigo-brito/facility-location/util/log"
)

func VNS(data *network.Data, solution *solution.Solution, perturbations ...neighborhoods.Perturbation) {
	log.Info("VNS started")

	tempSolution := solution.GetCopy()

	for position := 0; position < len(perturbations); position++ {
		// apply perturbation
		perturbations[position](tempSolution)

		// apply local search - VND
		VND(
			data, tempSolution,
			neighborhoods.ShiftLocalSearch,
			neighborhoods.RemoveHubLocalSearch,
			neighborhoods.AddHubLocalSearch,
			neighborhoods.SwapFunctionLocalSearch,
		)

		if tempSolution.GetCost(data) < solution.GetCost(data) {
			tempSolution.CopyTo(solution)
			log.Infof("SHIFT: New solution found FO=%.4f  hubs=%v", solution.GetCost(data), solution.Hubs)
			position = -1
		}
	}
}
