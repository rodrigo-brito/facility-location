package heuristic

import (
	"math"

	"github.com/rodrigo-brito/facility-location/model/network"
	"github.com/rodrigo-brito/facility-location/model/solution"
	"github.com/rodrigo-brito/facility-location/util"
	"github.com/rodrigo-brito/facility-location/util/log"
)

const variabilityFactor = 0.05

func NewSolution(data *network.Data) *solution.Solution {
	//var bestSolution = solution.New(data.Size, solution.WithInfinityCost())
	var bestSolution = solution.New(data.Size)
	bestSolution.AddHub(0)
	bestSolution.AllocateNearestHub(data)

	// Include initial node (minor cost)
	for initialNode := 0; initialNode < data.Size; initialNode++ {
		nodesBlocked := make([]bool, data.Size, data.Size)
		initialSolution := solution.New(data.Size)
		initialSolution.AddHub(initialNode)
		initialSolution.AllocateNearestHub(data)

		if initialSolution.GetCost(data) < bestSolution.GetCost(data) {
			log.Infof("New solution found FO=%.4f hubs=%v", initialSolution.GetCost(data), initialSolution.Hubs)
			initialSolution.CopyTo(bestSolution)
		}

		for {
			minMarginCost := math.MaxFloat64
			maxMarginCost := float64(0)

			solutions := make([]*solution.Solution, 0)

			initialSolution.Verify()

			for node := 0; node < data.Size; node++ {
				// If node is blocked ou already is hub, skip
				if nodesBlocked[node] || initialSolution.Allocation[node][node] {
					continue
				}

				tempSolution := initialSolution.GetCopy()
				tempSolution.AddHub(node)
				tempSolution.AllocateNearestHub(data)

				// If the solution is worst, block node insertion
				if tempSolution.GetCost(data) > initialSolution.GetCost(data) {
					nodesBlocked[node] = true
					continue
				}

				// Save the min and max values
				if tempSolution.GetCost(data) > maxMarginCost {
					maxMarginCost = tempSolution.GetCost(data)
				}

				if tempSolution.GetCost(data) < minMarginCost {
					minMarginCost = tempSolution.GetCost(data)
				}

				solutions = append(solutions, tempSolution)
			}

			if len(solutions) == 0 {
				break
			}

			var indexesPool []int
			valueReference := minMarginCost + variabilityFactor*(maxMarginCost-minMarginCost)
			for i, solution := range solutions {
				if solution.GetCost(data) <= valueReference {
					indexesPool = append(indexesPool, i)
				}
			}

			if len(indexesPool) > 0 {
				index := indexesPool[util.Random(0, len(indexesPool)-1)]
				solutions[index].CopyTo(initialSolution)
			}
		}

		if initialSolution.GetCost(data) < bestSolution.GetCost(data) {
			log.Infof("New solution found FO=%.4f hubs=%v", initialSolution.GetCost(data), initialSolution.Hubs)
			initialSolution.CopyTo(bestSolution)
		}
	}

	return bestSolution
}
