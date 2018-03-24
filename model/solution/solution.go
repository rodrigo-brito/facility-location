package solution

import (
	"fmt"

	"github.com/rodrigo-brito/hub-spoke-go/util/log"
)

// Solution store the network solution
type Solution struct {
	Hubs       []int
	HubsBin    map[int]bool
	Allocation [][]bool
	Cost       float64
}

// Generate the hubs list from hubs bin vector
func (s *Solution) generateHubList() {
	s.Hubs = make([]int, 0)
	for hub, isHub := range s.HubsBin {
		if isHub {
			s.Hubs = append(s.Hubs, hub)
		}
	}
}

// AddHub includes a new hub to the solution
func (s *Solution) AddHub(hub int) {
	s.HubsBin[hub] = true
	s.generateHubList()
}

// RemoveHub removes hub from solution
func (s *Solution) RemoveHub(hub int) {
	s.HubsBin[hub] = false
	s.generateHubList()
}

// Print the solution in stdout
func (s *Solution) Print() {
	log.Infof("COST = %18.4f", s.Cost)
	log.Infof("HUBS = %v", s.Hubs)
	log.Infof("HUBS BIN = %v", s.HubsBin)
	log.Info("ALLOCATION")
	log.Info("----------")
	for _, line := range s.Allocation {
		for _, column := range line {
			log.Infof("%v\t", column)
		}
		fmt.Println()
	}
}

// New creates and allocate memory for a new solution
func New(size int) *Solution {
	return &Solution{
		Hubs:       make([]int, 0),
		HubsBin:    make(map[int]bool, size),
		Allocation: make([][]bool, size, size),
	}
}
