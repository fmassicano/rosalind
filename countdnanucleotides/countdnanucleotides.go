package countdnanucleotides

import (
	"fmt"
	"sync"
)

/*
DNArecord is a type which contain the
count of each nucleotide
This struct use sync.Mutex now is safe to use concurrently.
*/
type DNArecord struct {
	A int
	C int
	G int
	T int
}

/*
Process generate the count string of A, C, G and T
*/
func (dna DNArecord) Process(s string) string {
	var wg sync.WaitGroup
	var m sync.Mutex
	defer dna.Reset(&m)
	dna.countDNAnucleotidesConcurrency(s, 10, &wg, &m)
	return dna.String()
}

func (dna *DNArecord) countDNAnucleotidesConcurrency(s string, n int, wg *sync.WaitGroup, m *sync.Mutex) {

	chunkSize := (len(s) + n - 1) / n

	for i := 0; i < len(s); i += chunkSize {
		end := i + chunkSize

		if end > len(s) {
			end = len(s)
		}
		wg.Add(1)

		go func(i, end int) {
			// fmt.Printf("%v - i = %v; end = %v\n", time.Now().Format("04:05.000000"), i, end)
			st := s[i:end]
			dna.CountDNAnucleotides(st, m)
			defer wg.Done()
		}(i, end)

	}
	wg.Wait()
}

/*
Reset the dna variable
*/
func (dna *DNArecord) Reset(m *sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	dna.A, dna.C, dna.G, dna.T = 0, 0, 0, 0
}

func (dna *DNArecord) String() string {
	return fmt.Sprintf("%v %v %v %v", dna.A, dna.C, dna.G, dna.T)
}

/*
CountDNAnucleotides count the number of events of each nucleotides (A,C,G,T)
  Given: A DNA string s of length at most 1000 nt.
  Return: Four integers (separated by spaces) counting the respective number of times that the symbols 'A', 'C', 'G', and 'T' occur in s.
*/
func (dna *DNArecord) CountDNAnucleotides(s string, m *sync.Mutex) {
	// fmt.Printf("%v -- CountDNAnucleotides\n", time.Now().Format("04:05.000000"))
	m.Lock()
	defer m.Unlock()
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
