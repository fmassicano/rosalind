package main

import (
	"os"

	"github.com/fmassicano/rosalind/complementingdna"
	"github.com/fmassicano/rosalind/runner"
)

func main() {
	// runner.Run(transcribingdnaintorna.RNA{}, os.Args[1])
	// runner.Run(countdnanucleotides.DNArecord{}, os.Args[1])
	runner.Run(complementingdna.DNAcompl{}, os.Args[1])
}
