package main

import (
	"os"

	// "github.com/fmassicano/rosalind/extractseqprot"

	"github.com/fmassicano/rosalind/runner"
)

func main() {
	// runner.Run(countdnanucleotides.DNArecord{}, os.Args[1])
	// runner.Run(transcribingdnaintorna.RNA{}, os.Args[1])
	// runner.Run(complementingdna.DNAcompl{}, os.Args[1])
	// runner.RunString(rabbitsrecurrencerelations.RabbitRecurrence{}, os.Args[1])
	// runner.RunString(gccontent.GC{}, os.Args[1])
	// runner.Run(countingpointmutations.CPM{}, os.Args[1])
	// runner.RunString(mendelfirstlaw.MendelFirstLaw{}, os.Args[1])
	// runner.Run(translatingrnaintoprotein.RNAintoProtein{}, os.Args[1])
	// runner.Run(findingamotifindna.Motif{}, os.Args[1])
	// runner.RunString(consensusandprofile.Profile{}, os.Args[1])
	// runner.RunString(mortalfibonaccirabbits.MortalFibonacci{}, os.Args[1])
	// runner.RunString(overlapgraphs.OverlapGraph{}, os.Args[1])
	// runner.RunString(calculatingexpectedoffspring.ExpOffSpring{}, os.Args[1])
	// runner.RunString(LIA.LIA{}, os.Args[1])
	runner.RunStringOnlyProcess(extractseqprot.ExtractSeqProt{}, os.Args[1])

}
