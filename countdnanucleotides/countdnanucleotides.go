package countdnanucleotides

import (
	"fmt"
)

/*
DNArecord is a type which contain the
count of each nucleotide
*/
type DNArecord struct {
	A int
	C int
	G int
	T int
}

var dna = DNArecord{}

/*
Process generate the count string of A, C, G and T
*/
func Process(s string) string {

	chunkSize := (len(s) + 10 - 1) / 10

	done := make(chan bool)

	for i := 0; i < len(s); i += chunkSize {
		end := i + chunkSize

		if end > len(s) {
			end = len(s)
		}

		go func(s string) {
			dna.CountDNAnucleotides(s)
			done <- true
		}(s[i:end])

		<-done
	}

	return dna.String()

}

/*
Reset the dna variable
*/
func Reset() {
	dna = DNArecord{}
}

func (dna DNArecord) String() string {
	return fmt.Sprintf("%v %v %v %v", dna.A, dna.C, dna.G, dna.T)
}

/*
CountDNAnucleotides count the number of events of each nucleotides (A,C,G,T)
  Given: A DNA string s of length at most 1000 nt.
  Return: Four integers (separated by spaces) counting the respective number of times that the symbols 'A', 'C', 'G', and 'T' occur in s.
*/
func (dna *DNArecord) CountDNAnucleotides(s string) {

	for _, v := range s {

		switch v {
		case 'A':
			dna.A++
		case 'C':
			dna.C++
		case 'G':
			dna.G++
		case 'T':
			dna.T++
		case '\n':
		default:
			fmt.Printf("%v is not a nucleotide\n", v)
		}
	}
}
