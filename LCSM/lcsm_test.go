package lcsm

import (
	"strings"
	"testing"
)

func TestFindingSharedMotif(t *testing.T) {

	var lc = LCSM{}

	var tests = []struct {
		s, want string
	}{
		{">Rosalind_1\nGATTACA\n>Rosalind_2\nTAGACCA\n>Rosalind_3\nATACA", "AC"},
	}

	for _, c := range tests {
		cll := CollectionLikelyLCSM{clcsm: make(map[string]int)}

		cll.getLikelyLCSM(dnaStrings)

		extractDNAstrings(strings.NewReader(c.s))

		cll.getLikelyLCSM(dnaStrings)

		lc.lcsm = cll.getLCSM(len(dnaStrings))

		got := lc.String()
		if strings.Trim(got, " \n") != c.want {
			//if got != c.want {
			t.Errorf("findingSharedMotif(%q) == %q, want %q", c.s, got, c.want)
		}
	}

}
