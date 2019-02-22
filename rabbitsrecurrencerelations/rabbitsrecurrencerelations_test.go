package rabbitsrecurrencerelations

import "testing"

func TestRabbitRecurrenceRelations(t *testing.T) {
	var rabrr = RabbitRecurrence{}
	var tests = []struct {
		s, want string
	}{
		{"5 3", "19"},
		{"5 4", "29"},
		{"35 4", "48127306357829"},
	}

	for _, c := range tests {
		got := rabrr.Process(c.s)
		if got != c.want {
			t.Errorf("RabbitRecurrenceRelations(%q) == %q, want %q", c.s, got, c.want)
		}
	}

}
