package LIA

import (
	"math"
	"strconv"
	"testing"
)

func Test_IndependentsAlleles(t *testing.T) {
	var eo = LIA{}

	var tests = []struct {
		s, want string
	}{
		{"2 1", "0.684"},
	}

	for _, c := range tests {
		got, _ := strconv.ParseFloat(eo.Process(c.s), 64)
		want, _ := strconv.ParseFloat(c.want, 64)

		if math.Abs(want-got) > 0.001 {
			t.Errorf("IndependentsAlleles(%v) == %v, want %v", c.s, got, want)
		}
	}

}
