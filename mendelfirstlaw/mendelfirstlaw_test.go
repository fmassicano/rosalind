package mendelfirstlaw

import (
	"math"
	"testing"
)

func TestProbDominantAllele(t *testing.T) {
	var mfl = MendelFirstLaw{}
	var tests = []struct {
		k, m, n, want float64
	}{
		{2, 2, 2, 0.78333333},
		{23, 29, 29, 0.71358024},
	}

	for _, c := range tests {
		mfl.probDominantAllele(c.k, c.m, c.n)
		got := mfl.result
		if math.Abs(got-c.want) > 0.0001 {
			t.Errorf("ProbDominantAllele(%v,%v,%v) == %v, want %v", c.k, c.m, c.n, got, c.want)
		}
	}

}
