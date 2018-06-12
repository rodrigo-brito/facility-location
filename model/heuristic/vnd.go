package heuristic

import (
	"github.com/rodrigo-brito/facility-location/model/heuristic/neighborhoods"
	"github.com/rodrigo-brito/facility-location/model/network"
	"github.com/rodrigo-brito/facility-location/model/solution"
	"github.com/rodrigo-brito/facility-location/util/log"
)

func VND(data *network.Data, solution *solution.Solution, neighborhoods ...neighborhoods.Neighborhood) {
	log.Info("VND started")
	for position := 0; position < len(neighborhoods); position++ {
		if newSolution := neighborhoods[position](data, solution); newSolution {
			log.Info("Neighborhood restarted.")
			position = -1
		}
	}
}
