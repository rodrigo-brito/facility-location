package solution

import (
	"github.com/rodrigo-brito/hub-spoke-go/model/network"
	"github.com/rodrigo-brito/hub-spoke-go/util/log"
)

// Solution store the network solution
type Solution struct {
	Hubs           []int
	Allocation     [][]bool
	AllocationNode []int
	cost           *float64
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
	s.Allocation[hub][hub] = true
	s.AllocationNode[hub] = hub
	s.cost = nil
	s.generateHubList()

	if len(s.Hubs) == 1 {
		s.initializeAllocation()
	}
}

// Allocate all nodes to initial hub
func (s *Solution) initializeAllocation() {
	for i := range s.AllocationNode {
		s.AllocNode(i, s.Hubs[0])
	}
}

// RemoveHub removes hub from solution
func (s *Solution) RemoveHub(hub int) {
	s.Allocation[hub][hub] = false
	s.cost = nil
	s.generateHubList()
}

//  Alloc a node to a hub and remove last allocation
func (s *Solution) AllocNode(node, hub int) {
	s.Allocation[node][s.AllocationNode[node]] = false
	s.Allocation[node][hub] = true
	s.AllocationNode[node] = hub
	s.cost = nil
}

// Print the solution in stdout
func (s *Solution) Print() {
	if s.cost != nil {
		log.Infof("COST = %.4f", *s.cost)
	}
	log.Infof("HUBS = %v", s.Hubs)
	//log.Info("ALLOCATION")
	//log.Info("----------")
	//for _, line := range s.Allocation {
	//	for _, column := range line {
	//		log.Infof("%v\t", column)
	//	}
	//	fmt.Println()
	//}
}

// New creates and allocate memory for a new solution
func New(size int) *Solution {
	s := &Solution{
		Hubs:           make([]int, 0),
		Allocation:     make([][]bool, size, size),
		AllocationNode: make([]int, size, size),
	}

	for i := range s.Allocation {
		s.Allocation[i] = make([]bool, size, size)
	}

	return s
}

func (s *Solution) NormalizeAllocation(data *network.Data) {
	for i := 0; i < data.Size; i++ {
		for _, hub := range s.Hubs {
			if data.Distance[i][hub] < data.Distance[i][s.AllocationNode[i]] {
				s.AllocNode(i, hub)
			}
		}
	}
}

func (s *Solution) GetCost(data *network.Data) float64 {
	// If solution has calculated cost
	if s.cost != nil {
		return *s.cost
	}

	s.cost = new(float64)

	// Installation cost
	for _, hub := range s.Hubs {
		*s.cost += data.InstallationCost[hub]
	}

	// Transport cost between node and hub
	for _, k := range s.Hubs {
		for i := 0; i < data.Size; i++ {
			if s.Allocation[i][k] && i != k {
				*s.cost += (data.FlowDestiny[i] + data.FlowOrigin[i]) * data.Distance[i][k]
			}
		}
	}

	// Transport cost between hubs
	for i := 0; i < data.Size; i++ {
		for j := i; j < data.Size; j++ {
			for _, k := range s.Hubs {
				for _, m := range s.Hubs {
					if k != m && s.Allocation[i][k] && s.Allocation[j][m] {
						*s.cost += data.Flow[i][j]*data.Distance[k][m]*data.ScaleFactor +
							data.Flow[j][i]*data.Distance[m][k]*data.ScaleFactor
					}
				}
			}
		}
	}

	return *s.cost
}
