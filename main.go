package main

import (
	"os"

	"github.com/fmassicano/rosalind/countdnanucleotides"

	"github.com/fmassicano/rosalind/runner"
	// "github.com/fmassicano/rosalind/transcribingdnaintorna"
)

func main() {
	// runner.Run(transcribingdnaintorna.RNA{}, os.Args[1])
	runner.Run(countdnanucleotides.DNArecord{}, os.Args[1])
}
