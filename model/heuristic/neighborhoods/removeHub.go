package neighborhoods

import (
	"github.com/rodrigo-brito/facility-location/model/network"
	"github.com/rodrigo-brito/facility-location/model/solution"
	"github.com/rodrigo-brito/facility-location/util"
	"github.com/rodrigo-brito/facility-location/util/async"
)

func RemoveHubPerturbation(solution *solution.Solution) {
	if len(solution.Hubs) > 1 {
		indexHub := util.Random(0, len(solution.Hubs)-1)
		solution.RemoveHub(solution.Hubs[indexHub])
	}
}

func RemoveHubLocalSearch(data *network.Data, bestSolution *solution.Solution) (newSolution bool) {
	tasks := make([]async.Task, 0)
	updatedChannel := make(chan bool, len(bestSolution.Hubs))

	for _, i := range bestSolution.Hubs {
		hub := i

		tasks = append(tasks, func(data *network.Data, solution *solution.Solution) {
			tempSolution := solution.GetCopy()
			if len(tempSolution.Hubs) < 2 {
				return
			}

			tempSolution.RemoveHub(hub)
			ShiftLocalSearch(data, tempSolution)

			updatedChannel <- solution.UpdateIfBetter(tempSolution, data)
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
