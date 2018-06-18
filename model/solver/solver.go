package solver

import (
	"fmt"
	"time"

	"github.com/rodrigo-brito/facility-location/model/heuristic"
	"github.com/rodrigo-brito/facility-location/model/heuristic/neighborhoods"
	"github.com/rodrigo-brito/facility-location/model/network"
	"github.com/rodrigo-brito/facility-location/model/solution"
	"github.com/rodrigo-brito/facility-location/util/log"
)

const defaultAsyncTasks = 1

type Solver struct {
	MaxAsyncTasks int
	Data          *network.Data
	BestSolution  *solution.Solution

	TargetCost *float64
	StartTime  time.Time
	EndTime    time.Time
}

func (s *Solver) Print() error {
	GAP := float64(0)
	fmt.Println("-------------- ")
	fmt.Printf("%d-%.1f - GAP|TIME|FO Hubs: %v\n", s.Data.Size, s.Data.ScaleFactor, s.BestSolution.Hubs)
	if s.TargetCost != nil {
		GAP = (s.BestSolution.GetCost(s.Data) - *s.TargetCost) / *s.TargetCost * 100
	}
	fmt.Printf("%.4f,%.4f,%.4f\n", GAP, s.EndTime.Sub(s.StartTime).Seconds(), s.BestSolution.GetCost(s.Data))
	return nil
}

func (s *Solver) initializeSolution() {
	s.BestSolution = heuristic.NewSolution(s.Data)
}

func (s *Solver) Solve() error {
	log.Info("Starting solver...")

	// Start timer
	s.StartTime = time.Now()

	// Initialize Solution - GRASP
	s.initializeSolution()

	// VNS
	heuristic.VNS(
		s.Data, s.BestSolution,
		neighborhoods.ShiftPerturbation,
		neighborhoods.RemoveHubPerturbation,
		neighborhoods.AddHubPerturbation,
		neighborhoods.SwapFunctionPerturbation,
	)

	//Finalize timer
	s.EndTime = time.Now()

	// Display the best result and time
	s.Print()

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

func WithTarget(targetValue float64) OptFunc {
	return func(solver *Solver) {
		if targetValue > 0 {
			solver.TargetCost = &targetValue
		}
	}
}

func WithInitialSolution(solution *solution.Solution) OptFunc {
	return func(solver *Solver) {
		solver.BestSolution = solution
	}
}
