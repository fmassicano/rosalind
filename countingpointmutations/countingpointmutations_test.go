package countingpointmutations

import "testing"

func TestCountingPointMutations(t *testing.T) {
	var cpm = CPM{}
	var tests = []struct {
		s, want string
	}{
		{"GAGCCTACTAACGGGAT\nCATCGTAATGACGGCCT", "7"},
	}

	for _, c := range tests {
		got := cpm.Process(c.s)
		if got != c.want {
			t.Errorf("CountingPointMutations(%q) == %q, want %q", c.s, got, c.want)
		}
	}

}
