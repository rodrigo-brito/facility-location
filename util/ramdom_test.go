package util

import (
	"testing"

	"github.com/smartystreets/assertions"
)

func TestRandom(t *testing.T) {
	timesRepeat := 100
	tt := []struct {
		Min int
		Max int
	}{
		{1, 10},
		{1, 2},
		{0, 1},
		{1, 1},
		{0, 0},
		{1000, 2000},
	}

	for _, tc := range tt {
		for i := 0; i < timesRepeat; i++ {
			res := Random(tc.Min, tc.Max)
			assertions.ShouldBeLessThanOrEqualTo(t, res, tc.Max)
			assertions.ShouldBeGreaterThanOrEqualTo(t, res, tc.Min)
		}
	}
}
