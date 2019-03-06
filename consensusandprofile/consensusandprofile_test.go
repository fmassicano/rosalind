package consensusandprofile

import (
	"strings"
	"testing"
)

func TestGetProfileConsensus(t *testing.T) {

	var prof = Profile{}

	var tests = []struct {
		s, want string
	}{
		{">Rosalind_1\nATCCAGCT\n>Rosalind_2\nGGGCAACT\n>Rosalind_3\nATGGATCT\n>Rosalind_4\nAAGCAACC\n>Rosalind_5\nTTGGAACT\n>Rosalind_6\nATGCCATT\n>Rosalind_7\nATGGCACT", "ATGCAACT\nA: 5 1 0 0 5 5 0 0\nC: 0 0 1 4 2 0 6 1\nG: 1 1 6 3 0 1 0 0\nT: 1 5 0 0 0 1 1 6"},
	}

	for _, c := range tests {
		prof.config(strings.NewReader(c.s))
		prof.buildProfile(strings.NewReader(c.s))
		prof.getConsesus()
		got := prof.String()
		if strings.Trim(got, " \n") != c.want {
			//if got != c.want {
			t.Errorf("getProfileConsensus(%q) == %q, want %q", c.s, got, c.want)
		}
	}

}
