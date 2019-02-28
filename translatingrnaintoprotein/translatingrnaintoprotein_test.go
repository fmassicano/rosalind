package translatingrnaintoprotein

import "testing"

func TestTranslatingRNAintoProtein(t *testing.T) {

	var rna = RNAintoProtein{}

	var tests = []struct {
		s, want string
	}{
		{"AUGGCCAUGGCGCCCAGAACUGAGAUCAAUAGUACCCGUAUUAACGGGUGA", "MAMAPRTEINSTRING"},
	}

	for _, c := range tests {
		got := rna.Process(c.s)
		if got != c.want {
			t.Errorf("transcribingDnaIntoRna(%q) == %q, want %q", c.s, got, c.want)
		}
	}

}
