package neighborhoods

import (
	"github.com/rodrigo-brito/facility-location/model/network"
	"github.com/rodrigo-brito/facility-location/model/solution"
)

type Neighborhood func(data *network.Data, solution *solution.Solution) bool

type Perturbation func(solution *solution.Solution)
