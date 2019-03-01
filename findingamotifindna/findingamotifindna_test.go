package findingamotifindna

import (
	"strings"
	"testing"
)

func TestGetMotifPositions(t *testing.T) {

	var motif = Motif{}

	var tests = []struct {
		s, want string
	}{
		{"GATATATGCATATACTT\nATAT", "2 4 10"},
		{"ATAT\nGATATATGCATATACTT", "The length of 's' needs to be bigger or equal than t."},
	}

	for _, c := range tests {
		got := motif.Process(c.s)
		if strings.Trim(got, " \n") != c.want {
			t.Errorf("getMotifPositions(%q) == %q, want %q", c.s, got, c.want)
		}
	}

}
