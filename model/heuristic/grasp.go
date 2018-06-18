package heuristic

import (
	"math"

	"github.com/rodrigo-brito/facility-location/model/network"
	"github.com/rodrigo-brito/facility-location/model/solution"
	"github.com/rodrigo-brito/facility-location/util"
	"github.com/rodrigo-brito/facility-location/util/async"
)

const variabilityFactor = 0.05

func NewSolution(data *network.Data) *solution.Solution {
	var bestSolution = solution.New(data.Size, solution.WithInfinityCost())

	tasks := make([]async.Task, 0)
	for i := 0; i < data.Size; i++ {
		initialNode := i
		tasks = append(tasks, func(data *network.Data, bestSolution *solution.Solution) {
			nodesBlocked := make([]bool, data.Size, data.Size)
			initialSolution := solution.New(data.Size)
			initialSolution.AddHub(initialNode)
			initialSolution.AllocateNearestHub(data)

			bestSolution.UpdateIfBetter(initialSolution, data)

			for {
				minMarginCost := math.MaxFloat64
				maxMarginCost := float64(0)

				solutions := make([]*solution.Solution, 0)

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

			bestSolution.UpdateIfBetter(initialSolution, data)
		})
	}

	// Run operation async
	async.Run(data, bestSolution, data.MaxAsyncTask, tasks...)

	return bestSolution
}
