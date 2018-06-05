package heuristic

import (
	"math"

	"github.com/rodrigo-brito/hub-spoke-go/model/network"
	"github.com/rodrigo-brito/hub-spoke-go/model/solution"
	"github.com/rodrigo-brito/hub-spoke-go/util"
	"github.com/rodrigo-brito/hub-spoke-go/util/log"
)

const variabilityFactor = 0.05

func NewSolution(data *network.Data) *solution.Solution {
	var bestSolution *solution.Solution

	nodesBlocked := make([]bool, data.Size, data.Size)

	// Include initial node (minor cost)
	for initialNode := 0; initialNode < data.Size; initialNode ++ {
		initialSolution := solution.New(data.Size)
		initialSolution.AddHub(initialNode)
		initialSolution.NormalizeAllocation(data)

		for {
			minMarginCost := math.Inf(1)
			maxMarginCost := float64(0)

			solutions := make([]*solution.Solution, 0)

			for node := 0; node < data.Size; node++ {
				// If node is blocked ou already is hub, skip
				if nodesBlocked[node] || initialSolution.Allocation[node][node] {
					continue
				}

				// copy the current solution
				tempSolution := *initialSolution
				tempSolution.AddHub(node)
				tempSolution.NormalizeAllocation(data)

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

				solutions = append(solutions, &tempSolution)
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

			index := indexesPool[util.Random(0, len(indexesPool)-1)]
			*initialSolution = *solutions[index]
		}

		if bestSolution == nil || initialSolution.GetCost(data) < bestSolution.GetCost(data) {
			log.Infof("New solution found hubs=%v FO=%.4f", initialSolution.Hubs, initialSolution.GetCost(data))
			bestSolution = initialSolution
		}
	}

	return bestSolution
}
