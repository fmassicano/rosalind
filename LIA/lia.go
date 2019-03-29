package LIA

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type LIA struct {
	prob float64
}

func (lia LIA) String() string {
	return fmt.Sprintf("%.20f", lia.prob)
}

func (lia LIA) Process(s string) string {

	input := strings.Split(s, " ")

	K, _ := strconv.Atoi(input[0])
	N, _ := strconv.Atoi(input[1])

	lia.independents_alleles(K, N)

	return lia.String()
}

// ∑N_i=n (N i) 0.25^(i) * 0.75^(N−i).
func (lia *LIA) independents_alleles(k, n int) {

	// fmt.Printf("k = %v and n = %v\n", k, n)

	N := math.Pow(float64(2), float64(k))

	for i := float64(n); i <= N; i++ {

		a := ncr(N, i)
		b := math.Pow(0.25, float64(i))
		c := math.Pow(.75, float64(N-i))

		lia.prob = lia.prob + (a * b * c)

	}

}

// n! / r! * (n - r)!
func ncr(N, i float64) float64 {
	return fact(N) / (fact(i) * fact(N-i))
}

func fact(x float64) float64 {
	f := float64(1)
	for i := float64(1); i <= x; i++ {
		f = i * f
	}
	return f
}
