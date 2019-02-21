package countdnanucleotides

import "testing"

func TestCountDNAnucleotides(t *testing.T) {
	var dna = DNArecord{}
	var tests = []struct {
		s, want string
	}{
		{"AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC", "20 12 17 21"},
		{"AAAAAA", "6 0 0 0"},
		{"AGCTTAA", "3 1 1 2"},
	}

	for _, c := range tests {
		got := dna.Process(c.s)
		if got != c.want {
			t.Errorf("CountDNAnucleotides(%q) == %q, want %q", c.s, got, c.want)
		}
		// dna.Reset() // If don't the Process continue accumulate.
	}

}
