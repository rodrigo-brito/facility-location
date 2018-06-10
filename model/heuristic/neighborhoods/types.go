package neighborhoods

import (
	"github.com/rodrigo-brito/hub-spoke-go/model/network"
	"github.com/rodrigo-brito/hub-spoke-go/model/solution"
)

type Neighborhood func(data *network.Data, solution *solution.Solution) bool

type Perturbation func(solution *solution.Solution)
