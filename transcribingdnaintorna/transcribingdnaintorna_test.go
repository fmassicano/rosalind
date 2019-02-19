package transcribingdnaintorna

import "testing"

func TestTranscribingDnaIntoRna(t *testing.T) {

	var rna = RNA{}

	var tests = []struct {
		s, want string
	}{
		{"GATGGAACTTGACTACGTAAATT", "GAUGGAACUUGACUACGUAAAUU"},
		{"GAATTTAAA", "GAAUUUAAA"},
	}

	for _, c := range tests {
		got := rna.Process(c.s)
		if got != c.want {
			t.Errorf("transcribingDnaIntoRna(%q) == %q, want %q", c.s, got, c.want)
		}
		// Reset() // If don't the Process continue accumulate.
	}

}
