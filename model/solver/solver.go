package solver

import (
	"time"

	"github.com/rodrigo-brito/hub-spoke-go/model/heuristic"
	"github.com/rodrigo-brito/hub-spoke-go/model/network"
	"github.com/rodrigo-brito/hub-spoke-go/model/solution"
)

const defaultAsyncTasks = 1

type Solver struct {
	MaxAsyncTasks int
	Data          *network.Data

	BestSolution *solution.Solution
	BestCost     float64

	StartTime time.Time
	EndTime   time.Time
}

func (s *Solver) initializeSolution() {
	solution := heuristic.NewSolution(s.Data)
	s.BestSolution = solution
}

func (s *Solver) Solve() error {
	s.StartTime = time.Now()

	s.initializeSolution()
	s.BestSolution.Print()

	s.EndTime = time.Now()
	return nil
}

func New(options ...OptFunc) *Solver {
	solver := new(Solver)

	solver.MaxAsyncTasks = defaultAsyncTasks

	for _, opt := range options {
		opt(solver)
	}

	return solver
}

type OptFunc func(*Solver)

func WithNetworkData(data *network.Data) OptFunc {
	return func(solver *Solver) {
		solver.Data = data
	}
}

func WithMaxAsyncTasks(limit int) OptFunc {
	return func(solver *Solver) {
		solver.MaxAsyncTasks = limit
	}
}

func WithInitialSolution(solution *solution.Solution) OptFunc {
	return func(solver *Solver) {
		solver.BestSolution = solution
	}
}
