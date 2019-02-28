package main

import (
	"os"

	"github.com/fmassicano/rosalind/runner"
	"github.com/fmassicano/rosalind/translatingrnaintoprotein"
)

func main() {
	// runner.Run(countdnanucleotides.DNArecord{}, os.Args[1])
	// runner.Run(transcribingdnaintorna.RNA{}, os.Args[1])
	// runner.Run(complementingdna.DNAcompl{}, os.Args[1])
	// runner.RunString(rabbitsrecurrencerelations.RabbitRecurrence{}, os.Args[1])
	// runner.RunString(gccontent.GC{}, os.Args[1])
	// runner.Run(countingpointmutations.CPM{}, os.Args[1])
	// runner.RunString(mendelfirstlaw.MendelFirstLaw{}, os.Args[1])
	runner.Run(translatingrnaintoprotein.RNAintoProtein{}, os.Args[1])
}
