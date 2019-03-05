package consensusandprofile

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/fmassicano/rosalind/utils"
)

type Profile struct {
	p          map[string][]int
	A, C, G, T []int
	consensus  string
}

func (prof *Profile) init(s string) {
	prof.p = map[string][]int{
		"A": make([]int, len(s)),
		"C": make([]int, len(s)),
		"G": make([]int, len(s)),
		"T": make([]int, len(s))}
	// fmt.Println("DNA string = ", len(s))
	// fmt.Println("A len = ", len(prof.p["A"]))
	// fmt.Println("C len = ", len(prof.p["C"]))
	// fmt.Println("G len = ", len(prof.p["G"]))
	// fmt.Println("T len = ", len(prof.p["T"]))
}

func (prof Profile) String() string {
	return fmt.Sprintf("%v\nA: %v\nC: %v\nG: %v\nT: %v\n", prof.consensus, utils.ConvertIntSlicesToString(prof.p["A"]), utils.ConvertIntSlicesToString(prof.p["C"]), utils.ConvertIntSlicesToString(prof.p["G"]), utils.ConvertIntSlicesToString(prof.p["T"]))
}

func (prof Profile) Process(f string) string {
	// fmt.Println("Process")
	utils.Analyze(f, prof.config)
	utils.Analyze(f, prof.buildProfile)
	prof.getConsesus()
	return prof.String()
}

func (prof *Profile) config(handle io.Reader) error {
	// fmt.Println("config")
	scanner := bufio.NewScanner(handle)
	var str strings.Builder
	i := 0
	for scanner.Scan() {
		matched, _ := regexp.MatchString(`>`, scanner.Text())

		if matched {
			i++
			if i > 1 {
				break
			}
			continue
		}

		str.WriteString(scanner.Text())

	}

	prof.init(str.String())

	return nil
}

func (prof *Profile) buildProfile(handle io.Reader) error {
	// fmt.Println("buildProfile")
	scanner := bufio.NewScanner(handle)
	var str strings.Builder
	firstMark := false
	for scanner.Scan() {
		matched, _ := regexp.MatchString(`>`, scanner.Text())

		if matched {
			if firstMark {
				prof.applyCount(&str)
			} else {
				firstMark = true
			}
			continue
		}

		str.WriteString(scanner.Text())

	}
	prof.applyCount(&str)
	// fmt.Println(prof.p)
	return nil
}
func (prof *Profile) applyCount(str *strings.Builder) {
	prof.getCount(str.String())
	str.Reset()
}

func (prof *Profile) getCount(s string) {
	// fmt.Println("getCount")
	for i, v := range s {
		switch v {
		case 'A':
			prof.p["A"][i]++
		case 'C':
			prof.p["C"][i]++
		case 'G':
			prof.p["G"][i]++
		case 'T':
			prof.p["T"][i]++
		}
	}
}

func (prof *Profile) getConsesus() {
	//	fmt.Println("getConsesus")
	var str strings.Builder
	for i := 0; i < len(prof.p["A"]); i++ {
		str.WriteString(max([4]int{prof.p["A"][i], prof.p["C"][i], prof.p["G"][i], prof.p["T"][i]}))
	}

	prof.consensus = str.String()
}

// vec is a seq of A C G T
func max(vec [4]int) string {
	var r string
	symbol := [4]string{"A", "C", "G", "T"}
	max := 0
	for i, v := range vec {
		if v > max {
			r = symbol[i]
			max = v
		}
	}
	return r
}
