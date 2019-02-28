package translatingrnaintoprotein

import (
	"bytes"
	"fmt"
)

type RNAintoProtein struct {
	rna     string
	protein bytes.Buffer
}

func (rna2prot *RNAintoProtein) init(s string) {
	rna2prot.rna = s
}

func (rna2prot RNAintoProtein) String() string {
	return fmt.Sprintf("%v", rna2prot.protein.String())
}

func (rna2prot RNAintoProtein) Process(s string) string {
	rna2prot.init(s)
	rna2prot.translatingRNAintoProtein()
	return rna2prot.String()
}

var rnaCodonTable = map[string][]string{
	"F":    []string{"UUU", "UUC"},
	"L":    []string{"UUA", "UUG", "CUU", "CUC", "CUA", "CUG"},
	"S":    []string{"UCU", "UCC", "UCA", "UCG", "AGU", "AGC"},
	"Y":    []string{"UAU", "UAC"},
	"STOP": []string{"UAA", "UAG", "UAA", "UAG", "UGA"},
	"C":    []string{"UGU", "UGC"},
	"W":    []string{"UGG", "UAG"},
	"P":    []string{"CCU", "CCC", "CCA", "CCG"},
	"H":    []string{"CAU", "CAC"},
	"Q":    []string{"CAA", "CAG"},
	"R":    []string{"CGU", "CGC", "CGA", "CGG", "AGA", "AGG"},
	"I":    []string{"AUU", "AUC", "AUA"},
	"M":    []string{"AUG"},
	"T":    []string{"ACU", "ACC", "ACA", "ACG"},
	"N":    []string{"AAU", "AAC"},
	"K":    []string{"AAA", "AAG"},
	"V":    []string{"GUU", "GUC", "GUA", "", "GUG"},
	"A":    []string{"GCU", "GCC", "GCA", "GCG"},
	"D":    []string{"GAU", "GAC"},
	"E":    []string{"GAA", "GAG"},
	"G":    []string{"GGU", "GGC", "GGA", "GGG"},
}

func getProtein(codon string) string {
	for key, value := range rnaCodonTable {

		for _, v := range value {
			if v == codon {
				return key
			}
		}

	}
	return "-"
}

func (rna2prot *RNAintoProtein) translatingRNAintoProtein() {

	var codon, amino string

	for st, en := 0, 3; en < len(rna2prot.rna)+1; st, en = en, en+3 {
		codon = rna2prot.rna[st:en]
		amino = getProtein(codon)

		//fmt.Printf("%v -> %v\n", codon, amino)
		if amino != "STOP" {
			rna2prot.protein.WriteString(amino)
		} else {
			return
		}

	}

}
