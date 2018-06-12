package solution

import (
	"testing"

	"github.com/rodrigo-brito/facility-location/model/network"
	"github.com/rodrigo-brito/facility-location/util/log"
	"github.com/stretchr/testify/assert"
)

var data *network.Data

func init() {
	var err error
	data, err = network.FromFile("../../data/ap10_2.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func TestSolution_GetCostCalculation(t *testing.T) {
	s := new(Solution)
	s.Hubs = []int{0, 3, 4}
	s.Allocation = [][]bool{
		{true, false, false, false, false, false, false, false, false, false},
		{false, false, false, true, false, false, false, false, false, false},
		{false, false, false, false, true, false, false, false, false, false},
		{false, false, false, true, false, false, false, false, false, false},
		{false, false, false, false, true, false, false, false, false, false},
		{false, false, false, true, false, false, false, false, false, false},
		{false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, true, false, false, false, false, false},
	}

	expectedValue := 90963539.4763
	assert.InDelta(t, expectedValue, s.GetCost(data), 0.0001)
}

func TestSolution_GetCostReturn(t *testing.T) {
	expectedValue := &[]float64{1}[0]
	s := &Solution{
		Cost: expectedValue,
	}
	s.Hubs = []int{0, 3, 4}

	assert.Equal(t, *expectedValue, s.GetCost(data))
}

func TestSolution_AddHub(t *testing.T) {
	s := New(data.Size)
	s.AddHub(1)
	s.AllocateNearestHub(data)
	assert.Equal(t, s.Allocation[1][1], true)
	assert.Equal(t, s.AllocationNode[1], 1)
	assert.Len(t, s.Hubs, 1)
	assert.Contains(t, s.Hubs, 1)

	for i := range s.Allocation {
		assert.True(t, s.Allocation[i][1])
	}

	s.AddHub(2)
	s.AllocateNearestHub(data)
	assert.Equal(t, s.Allocation[2][1], false)
	assert.Equal(t, s.Allocation[2][2], true)
	assert.Equal(t, s.AllocationNode[2], 2)
	assert.Len(t, s.Hubs, 2)
	assert.Contains(t, s.Hubs, 1)
	assert.Contains(t, s.Hubs, 2)
}

func TestSolution_RemoveHub(t *testing.T) {
	s := New(data.Size)

	s.AddHub(1)
	s.AddHub(3)
	s.AddHub(4)

	s.RemoveHub(3)
	assert.Equal(t, s.Allocation[3][3], false)
	assert.Len(t, s.Hubs, 2)
	assert.NotContains(t, s.Hubs, 3)

	s.RemoveHub(1)
	assert.Equal(t, s.Allocation[1][1], false)
	assert.Len(t, s.Hubs, 1)
	assert.NotContains(t, s.Hubs, 1)
}

func TestSolution_AllocNode(t *testing.T) {
	s := New(data.Size)
	s.AddHub(1)
	s.AddHub(2)

	s.AllocNode(3, 1)
	assert.Equal(t, s.Allocation[1][1], true)
	assert.Equal(t, s.Allocation[3][1], true)
	assert.Equal(t, s.AllocationNode[3], 1)

	s.AllocNode(3, 2)
	assert.Equal(t, s.Allocation[3][1], false)
	assert.Equal(t, s.Allocation[3][2], true)
	assert.Equal(t, s.AllocationNode[3], 2)
}
