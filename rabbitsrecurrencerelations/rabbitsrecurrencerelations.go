package rabbitsrecurrencerelations

import (
	"fmt"
	"strconv"
	"strings"
)

// RabbitRecurrence is
type RabbitRecurrence struct {
	r []int
}

func (rabc *RabbitRecurrence) init(n int) {
	rabc.r = make([]int, n)
	rabc.r[0] = 1
	rabc.r[1] = 1
}

func (rabc RabbitRecurrence) String() string {
	return fmt.Sprintf("%v", rabc.r[len(rabc.r)-1])
}

// Process functions return the number of rabbits
// pairs that will be present after n months
func (rabc RabbitRecurrence) Process(s string) string {
	v := strings.Split(s, " ")
	n, _ := strconv.Atoi(v[0])
	k, _ := strconv.Atoi(v[1])
	rabc.init(n)
	rabc.xpot(n, k)
	return rabc.String()
}

func (rabc *RabbitRecurrence) xpot(n, k int) {
	// fmt.Printf("n = %v; k = %v\n", n, k)
	x, y := 1, 1 // First and Second generations
	for i := 2; i < (n); i++ {
		x, y = y, (k*x)+y
		rabc.r[i] = y
	}
}
