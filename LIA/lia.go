package LIA

// ∑N_i=n (N i) 0.25^(i) * 0.75^(N−i).

func ncr(N, i int) int {

	return 0
}

func fact(x int) int {
	f := 1
	for i := 1; i <= x; i++ {
		f = i * f
	}
	return f
}
