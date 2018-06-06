package solution

import (
	"fmt"
	"math"

	"github.com/rodrigo-brito/hub-spoke-go/model/network"
	"github.com/rodrigo-brito/hub-spoke-go/util/log"
)

// Solution store the network solution
type Solution struct {
	Size           int
	Hubs           []int
	Allocation     [][]bool
	AllocationNode []int
	Cost           *float64
}

// Generate the hubs list from hubs bin vector
func (s *Solution) generateHubList() {
	s.Hubs = make([]int, 0)
	for hub := range s.Allocation {
		if s.Allocation[hub][hub] {
			s.Hubs = append(s.Hubs, hub)
		}
	}
}

// AddHub includes a new hub to the solution
func (s *Solution) AddHub(hub int) {
	if s.Allocation[hub][hub] {
		panic("hub addiction not permitted") //TODO: remove check
	}

	s.Allocation[hub][s.AllocationNode[hub]] = false
	s.Allocation[hub][hub] = true
	s.AllocationNode[hub] = hub
	s.Hubs = append(s.Hubs, hub)
	s.Cost = nil

	if len(s.Hubs) == 1 {
		s.initializeAllocation()
	}

}

// Allocate all nodes to initial hub
func (s *Solution) initializeAllocation() {
	for i := range s.AllocationNode {
		s.Allocation[i][s.Hubs[0]] = true
		s.AllocationNode[i] = s.Hubs[0]
	}
}

// RemoveHub removes hub from solution
func (s *Solution) RemoveHub(hub int) {
	if len(s.Hubs) < 2 {
		panic("hub remove not permitted") //TODO: remove check
	}

	s.Allocation[hub][hub] = false
	s.generateHubList()
	s.Allocation[hub][s.Hubs[0]] = true
	s.Cost = nil
}

//  Alloc a node to a hub and remove last allocation
func (s *Solution) AllocNode(node, hub int) {
	if !s.Allocation[hub][hub] || s.Allocation[node][node] {
		panic("allocation not permitted")
	}

	s.Allocation[node][s.AllocationNode[node]] = false
	s.Allocation[node][hub] = true
	s.AllocationNode[node] = hub
	s.Cost = nil
}

// Print the solution in stdout
func (s *Solution) Print() {
	if s.Cost != nil {
		log.Infof("COST = %.4f", *s.Cost)
	}
	log.Infof("HUBS = %v", s.Hubs)
	fmt.Println("ALLOCATION")
	fmt.Println("----------")
	for _, line := range s.Allocation {
		for _, column := range line {
			fmt.Printf("%v\t", column)
		}
		fmt.Println()
	}
}

// New creates and allocate memory for a new solution
func New(size int, options ...SolutionOption) *Solution {
	s := &Solution{
		Size:           size,
		Hubs:           make([]int, 0),
		Allocation:     make([][]bool, size, size),
		AllocationNode: make([]int, size, size),
	}

	for i := range s.Allocation {
		s.Allocation[i] = make([]bool, size, size)
	}

	for _, option := range options {
		option(s)
	}

	return s
}

type SolutionOption func(s *Solution)

func WithInfinityCost() SolutionOption {
	return func(s *Solution) {
		inf := math.MaxFloat64
		s.Cost = &inf
	}
}

func (s *Solution) AllocateNearestHub(data *network.Data) {
	for i := 0; i < data.Size; i++ {
		for j := i; j < data.Size; j++ {
			if !s.Allocation[j][j] || s.Allocation[i][i] {
				continue
			}

			if !s.Allocation[i][i] && s.Allocation[j][j] && data.Distance[i][j] < data.Distance[i][s.AllocationNode[i]] {
				s.AllocNode(i, j)
			}
		}
	}
}

func (s *Solution) GetCost(data *network.Data) float64 {
	// If solution has calculated Cost
	if s.Cost != nil {
		return *s.Cost
	}

	s.Cost = new(float64)

	// Installation Cost
	for _, hub := range s.Hubs {
		*s.Cost += data.InstallationCost[hub]
	}

	// Transport Cost between node and hub
	for _, k := range s.Hubs {
		for i := 0; i < data.Size; i++ {
			if s.Allocation[i][k] && i != k {
				*s.Cost += (data.FlowDestiny[i] + data.FlowOrigin[i]) * data.Distance[i][k]
			}
		}
	}

	// Transport Cost between hubs
	for i := 0; i < data.Size; i++ {
		for j := i; j < data.Size; j++ {
			for _, k := range s.Hubs {
				for _, m := range s.Hubs {
					if k != m && s.Allocation[i][k] && s.Allocation[j][m] {
						*s.Cost += data.Flow[i][j]*data.Distance[k][m]*data.ScaleFactor +
							data.Flow[j][i]*data.Distance[m][k]*data.ScaleFactor
					}
				}
			}
		}
	}

	return *s.Cost
}

func (s *Solution) CopyTo(copy *Solution) {
	s.Hubs = make([]int, 0)
	copy.Cost = new(float64)
	*copy.Cost = *s.Cost
	for i := 0; i < s.Size; i++ {
		copy.AllocationNode[i] = s.AllocationNode[i]
		for j := 0; j < s.Size; j++ {
			copy.Allocation[i][j] = s.Allocation[i][j]
		}
	}
	copy.generateHubList()
}

func (s *Solution) GetCopy() *Solution {
	copy := New(s.Size)
	s.CopyTo(copy)
	return copy
}

func (s *Solution) Verify() { // TODO: remover
	for i := 0; i < s.Size; i++ {
		sum := 0
		for j := 0; j < s.Size; j++ {
			if s.Allocation[i][j] {
				sum++
			}
		}
		if sum != 1 {
			s.Print()
			panic("invalid solution [aloc]")
		}
		for _, hub := range s.Hubs {
			if !s.Allocation[hub][hub] {
				panic("invalid solution [hub aloc]")
			}
		}
	}
}
