package extractseqprot

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/fmassicano/rosalind/utils"
)

type ExtractSeqProt struct {
	status string
}

func (esp ExtractSeqProt) String() string {
	return fmt.Sprintf("Status: %v", esp.status)
}

func (esp ExtractSeqProt) Process(f string) string {
	err := utils.Analyze(f, extractSeqProt)
	if err == nil {
		esp.status = "Failed"
	} else {
		esp.status = ""
	}
	return esp.status
}

func extractSeqProt(handle io.Reader) error {

	scanner := bufio.NewScanner(handle)
	firstMark := true

	for scanner.Scan() {

		matched, _ := regexp.MatchString(`>`, scanner.Text())

		if matched {
			protein := strings.Split(scanner.Text(), "|")
			fmt.Print(protein[2], " ")

			if firstMark {
				firstMark = !firstMark
			} else {
				// fmt.Println(scanner.Text())
			}
			continue
		}

		// fmt.Print(scanner.Text())

	}

	// fmt.Println(scanner.Text())

	return nil
}
