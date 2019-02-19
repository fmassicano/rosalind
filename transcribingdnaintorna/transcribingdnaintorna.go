package transcribingdnaintorna

// An RNA string is a string formed from the alphabet containing 'A', 'C', 'G', and 'U'.
type RNA struct {
	nt, dna []rune
}

/*
Process generate the transcribing DNA into RNA
*/
func (rna RNA) Process(dna string) string {

	rna.init(dna)

	chunkSize := (len(rna.dna) + 10 - 1) / 10

	done := make(chan bool)

	for i := 0; i < len(rna.dna); i += chunkSize {
		end := i + chunkSize

		if end > len(rna.dna) {
			end = len(rna.dna)
		}

		go func(start, end int) {
			rna.transcribingDnaIntoRna(start, end)
			done <- true
		}(i, end)

		<-done
	}

	return string(rna.nt)
}

func (rna *RNA) init(dna string) {
	rna.dna = []rune(dna)
	rna.nt = make([]rune, len(rna.dna))
}

/*
Given a DNA string t corresponding to a coding strand, its transcribed RNA string u is formed by replacing all occurrences of 'T' in t with 'U' in u.
Given: A DNA string t having length at most 1000 nt.
Return: The transcribed RNA string of t.
*/
func (rna *RNA) transcribingDnaIntoRna(start, end int) {

	for i := start; i < end; i++ {
		if rna.dna[i] == 'T' {
			rna.nt[i] = 'U'
		} else {
			rna.nt[i] = rna.dna[i]
		}
	}

}
