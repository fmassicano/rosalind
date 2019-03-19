package calculatingexpectedoffspring

import "testing"

func Test_expOffSpringDominant(t *testing.T) {
	var eo = ExpOffSpring{}

	var tests = []struct {
		s, want string
	}{
		{"1 0 0 1 0 1", "3.50"},
		{"0 0 0 0 0 0", "0.00"},
		{"1 1 1 1 1 1", "8.50"},
	}

	for _, c := range tests {
		got := eo.Process(c.s)
		if got != c.want {
			t.Errorf("expOffSpringDominant(%q) == %q, want %q", c.s, got, c.want)
		}
	}

}
