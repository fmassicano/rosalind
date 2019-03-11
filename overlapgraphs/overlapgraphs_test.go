package overlapgraphs

import (
	"strings"
	"testing"
)

func TestGraphCreator(t *testing.T) {
	var og = OverlapGraph{}
	var tests = []struct {
		s, want string
	}{
		{">Rosalind_0498\nAAATAAA\n>Rosalind_2391\nAAATTTT\n>Rosalind_2323\nTTTTCCC\n>Rosalind_0442\nAAATCCC\n>Rosalind_5013\nGGGTGGG", "Rosalind_0498 Rosalind_2391\nRosalind_0498 Rosalind_0442\nRosalind_2391 Rosalind_2323"},
	}

	for _, c := range tests {
		og.init()
		og.createNodes(strings.NewReader(c.s))
		og.createLinks()
		got := strings.Trim(og.String(), "\n")
		if c.want != got {
			t.Errorf("GraphCreator(%v) == %v, want %v", c.s, got, c.want)
		}
	}

}
