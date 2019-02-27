package mendelfirstlaw

import (
	"fmt"
	"strconv"
	"strings"
)

type MendelFirstLaw struct {
	result float64
}

func (mfl MendelFirstLaw) String() string {
	return fmt.Sprintf("%v", mfl.result)
}

func (mfl MendelFirstLaw) Process(s string) string {

	d := strings.Split(s, " ")

	k, _ := strconv.ParseFloat(d[0], 64)
	m, _ := strconv.ParseFloat(d[1], 64)
	n, _ := strconv.ParseFloat(d[2], 64)

	mfl.probDominantAllele(k, m, n)

	return mfl.String()

}

func (mfl *MendelFirstLaw) probDominantAllele(k, m, n float64) {

	probDominantAlleleByCrossing := map[string]float64{"YYYY": 1, "YYYy": 1, "YYyy": 1, "YyYY": 1, "YyYy": 3. / 4, "Yyyy": 2. / 4, "yyYY": 1, "yyYy": 2. / 4, "yyyy": 0}

	mfl.result = calcProb(k, m, n, probDominantAlleleByCrossing)

}

func calcProb(k, m, n float64, pByCrossing map[string]float64) float64 {

	var r float64

	total := float64(k + m + n)

	prob := map[string]float64{
		"YYYY": (k / total) * ((k - 1) / (total - 1)),
		"YYYy": (k / total) * ((m) / (total - 1)),
		"YYyy": (k / total) * ((n) / (total - 1)),
		"YyYY": (m / total) * ((k) / (total - 1)),
		"YyYy": (m / total) * ((m - 1) / (total - 1)),
		"Yyyy": (m / total) * ((n) / (total - 1)),
		"yyYY": (n / total) * ((k) / (total - 1)),
		"yyYy": (n / total) * ((m) / (total - 1)),
		"yyyy": (n / total) * ((n - 1) / (total - 1)),
	}

	for key := range prob {
		r = r + (prob[key] * pByCrossing[key])
	}
	return r
}
