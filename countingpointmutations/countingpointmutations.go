package countingpointmutations

import (
	"fmt"
	"strings"
)

type CPM struct {
	hammingDist               int
	firstStrand, secondStrand string
}

func (cpm *CPM) init(s string) {
	sn := strings.Split(s, "\n")
	cpm.firstStrand = sn[0]
	cpm.secondStrand = sn[1]
	cpm.hammingDist = 0
	// fmt.Println(cpm.firstStrand)
	// fmt.Println(cpm.secondStrand)
	// fmt.Println(cpm.hammingDist)
}

func (cpm CPM) String() string {
	return fmt.Sprintf("%v", cpm.hammingDist)
}

func (cpm CPM) Process(s string) string {
	cpm.init(s)
	cpm.hDist()
	return cpm.String()
}

func (cpm *CPM) hDist() {
	for i := 0; i < len(cpm.firstStrand); i++ {
		if cpm.firstStrand[i] != cpm.secondStrand[i] {
			cpm.hammingDist++
		}
	}
}
