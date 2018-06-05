package heuristic

import (
	"math"

	"github.com/rodrigo-brito/hub-spoke-go/model/network"
	"github.com/rodrigo-brito/hub-spoke-go/model/solution"
	"github.com/rodrigo-brito/hub-spoke-go/util"
)

func selectBestNode(data *network.Data) int {
	var (
		indexMin = 0
		valueMin = data.InstallationCost[0]
	)

	for i := 1; i < data.Size; i++ {
		if data.InstallationCost[i] < valueMin {
			indexMin = i
			valueMin = data.InstallationCost[i]
		}
	}

	return indexMin
}

func NewSolution(data *network.Data) *solution.Solution {
	bestSolution := solution.New(data.Size)

	nodesBlocked := make([]bool, data.Size, data.Size)

	// Include initial node (minor cost)
	initialNode := selectBestNode(data)
	bestSolution.AddHub(initialNode)

	for {
		minMarginCost := math.Inf(1)
		maxMarginCost := float64(0)

		solutions := make([]*solution.Solution, 0)

		for node := 0; node < data.Size; node++ {
			// If node is blocked ou already is hub, skip
			if nodesBlocked[node] || bestSolution.Allocation[node][node] {
				continue
			}

			// copy the current solution
			tempSolution := *bestSolution
			tempSolution.AddHub(node)

			// If the solution is worst, block node insertion
			if tempSolution.GetCost(data) > bestSolution.GetCost(data) {
				nodesBlocked[node] = true
				continue
			}

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
		valueReference := minMarginCost + 0.05*(maxMarginCost-minMarginCost)
		for i, solution := range solutions {
			if solution.GetCost(data) <= valueReference {
				indexesPool = append(indexesPool, i)
			}
		}

		index := indexesPool[util.Random(0, len(indexesPool)-1)]
		*bestSolution = *solutions[index]
	}

	return bestSolution
}
