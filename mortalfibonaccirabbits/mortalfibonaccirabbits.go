package mortalfibonaccirabbits

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

type MortalFibonacci struct {
	r []*big.Int
}

func (mf *MortalFibonacci) init(n int) {
	mf.r = make([]*big.Int, n)
	mf.r[0] = big.NewInt(int64(1))
	mf.r[1] = big.NewInt(int64(1))
}

func (mf MortalFibonacci) String() string {
	return fmt.Sprintf("%v", mf.r[len(mf.r)-1])
}

func (mf MortalFibonacci) Process(s string) string {
	v := strings.Split(s, " ")
	n, _ := strconv.Atoi(v[0])
	m, _ := strconv.Atoi(v[1])
	mf.init(n)
	mf.xpto(n, m, 1)
	return mf.String()
}

func (mf *MortalFibonacci) xpto(n, m, k int) {
	var F, M = make([]*big.Int, n), make([]*big.Int, n)

	F[0] = big.NewInt(int64(1))
	M[0] = big.NewInt(int64(0))

	for i := 1; i < (n); i++ {
		F[i] = new(big.Int).Mul(M[i-1], big.NewInt(int64(k)))
		M[i] = new(big.Int).Add(F[i-1], M[i-1])
		if i-m >= 0 {
			M[i] = new(big.Int).Sub(M[i], F[i-m])
		}
		mf.r[i] = new(big.Int).Add(F[i], M[i])
	}

}
