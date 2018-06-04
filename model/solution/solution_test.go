package solution

import (
	"testing"

	"github.com/rodrigo-brito/hub-spoke-go/model/network"
	"github.com/rodrigo-brito/hub-spoke-go/util/log"
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

func TestSolution_Value(t *testing.T) {
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
	assert.InDelta(t, expectedValue, s.Value(data), 0.0001)
}
