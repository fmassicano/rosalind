package gccontent

import (
	"bufio"
	"fmt"
	"io"
	"regexp"

	"github.com/fmassicano/rosalind/utils"
)

// GC is a struct to get the ID and frequency
type GC struct {
	ID  string
	gcf float32
	// 0 keeps the gc count and 1 keep the total
	track map[string][]float32
}

func (gc GC) String() string {
	return fmt.Sprintf("%v\n%v", gc.ID, gc.gcf)
}

// Process is a function used by runner to execute GCcontent
func (gc GC) Process(file string) string {
	utils.Analyze(file, gc.getAllGCcontent)
	gc.getTheHighestGCcount()
	return gc.String()
}

func (gc *GC) getAllGCcontent(handle io.Reader) error {
	gc.track = make(map[string][]float32)
	idRosalind := ""
	scanner := bufio.NewScanner(handle)

	for scanner.Scan() {
		matched, _ := regexp.MatchString(`>`, scanner.Text())

		if matched {
			idRosalind = scanner.Text()[1:]
			gc.track[idRosalind] = []float32{0, 0}
			continue
		}

		if idRosalind != "" {
			gc.track[idRosalind][0] = gc.track[idRosalind][0] + gc.gcCount(scanner.Text())
			gc.track[idRosalind][1] = gc.track[idRosalind][1] + float32(len(scanner.Text()))
		}

	}

	return nil
}

func (gc GC) gcCount(s string) float32 {
	var sum float32

	for _, v := range s {
		if v == 'G' || v == 'C' {
			sum++
		}
	}

	return sum
}

func (gc *GC) getTheHighestGCcount() {

	var p float32

	for k, v := range gc.track {
		pTemp := v[0] / v[1]
		if pTemp > p {
			gc.ID = k
			gc.gcf = 100 * pTemp
			p = pTemp
		}
	}

}
