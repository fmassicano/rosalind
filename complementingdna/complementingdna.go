package complementingdna

import (
	"sync"
)

// DNAcompl is a type used to keep the dna and its complementing
type DNAcompl struct {
	dna, compldna []rune
}

func (d *DNAcompl) init(dna string) {
	d.dna = []rune(dna)
	d.compldna = make([]rune, len(d.dna))
}

// Process is a function to return the complementing of dna
func (d DNAcompl) Process(s string) string {
	d.init(s)
	d.reverseConcurrency(10)
	d.complementConcurrency(5)
	return string(d.compldna)
}

func (d *DNAcompl) reverseConcurrency(n int) {
	var wg sync.WaitGroup
	chunkSize := (len(d.dna) + n - 1) / n

	for i := 0; i < len(d.dna); i += chunkSize {
		end := i + chunkSize

		if end > len(d.dna) {
			end = len(d.dna)
		}
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			d.reverse(start, end)
		}(i, end)

	}
	wg.Wait()
}

// reverse is a function that reverse a strand of dna
func (d *DNAcompl) reverse(start, end int) {
	// fmt.Printf("reverse: start = %v, end = %v\n", start, end)
	for i := start; i < end; i++ {
		d.compldna[i] = d.dna[len(d.dna)-i-1]
	}
}

func (d *DNAcompl) complementConcurrency(n int) {
	var wg sync.WaitGroup
	chunkSize := (len(d.dna) + n - 1) / n

	for i := 0; i < len(d.dna); i += chunkSize {
		end := i + chunkSize

		if end > len(d.dna) {
			end = len(d.dna)
		}
		wg.Add(1)
		go func(start, end int) {
			d.complement(start, end)
			wg.Done()
		}(i, end)
	}
	wg.Wait()
}

// complement is a function that complement a strand of dna
func (d *DNAcompl) complement(start, end int) {
	// fmt.Printf("complement: start = %v, end = %v\n", start, end)
	for i := start; i < end; i++ {
		switch d.compldna[i] {
		case 'A':
			d.compldna[i] = 'T'
		case 'T':
			d.compldna[i] = 'A'
		case 'C':
			d.compldna[i] = 'G'
		case 'G':
			d.compldna[i] = 'C'
		}
	}
}
