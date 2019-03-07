package mortalfibonaccirabbits

import "testing"

func TestMortalRabbits(t *testing.T) {
	var mf = MortalFibonacci{}
	var tests = []struct {
		s, want string
	}{
		{"6 3", "4"},
		{"90 20", "2873957992832138936"},
	}

	for _, c := range tests {
		got := mf.Process(c.s)
		if got != c.want {
			t.Errorf("MortalRabbits(%q) == %q, want %q", c.s, got, c.want)
		}
	}

}
