package main

import (
	"fmt"
	"io/ioutil"

	"github.com/fmassicano/rosalind/countdnanucleotides"
)

func main() {

	dat, err := ioutil.ReadFile("/Users/felipe/Downloads/rosalind_dna_1_dataset.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(countdnanucleotides.Process(string(dat)))

}
