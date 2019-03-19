package calculatingexpectedoffspring

import (
	"fmt"
	"strconv"
	"strings"
)

type ExpOffSpring struct {
	probGenotypes map[int]float32
	expDominant   float32
}

func (eo *ExpOffSpring) init() {
	// eo.probGenotypes = make(map[int]float32)
	eo.probGenotypes = map[int]float32{1: 2.0, 2: 2.0, 3: 2.0, 4: (2 * .75), 5: (2 * 0.5), 6: 0.0}
	eo.expDominant = 0.0
}

func (eo ExpOffSpring) String() string {
	return fmt.Sprintf("%.2f", eo.expDominant)
}

func (eo ExpOffSpring) Process(s string) string {
	eo.init()
	vals := strings.Split(s, " ")
	eo.calc(vals)
	return eo.String()
}

func (eo *ExpOffSpring) calc(vals []string) {
	for i, v := range eo.probGenotypes {
		val, _ := strconv.ParseFloat(strings.Trim(vals[i-1], " "), 32)
		eo.expDominant = eo.expDominant + (float32(val) * v)
	}
}
