package runner

import (
	"fmt"
	"io/ioutil"
)

type runner interface {
	Process(s string) string
}

// Run is a function that receive a function f which implement
// runner and process the data contained in the path file p.
func Run(f runner, p string) {
	dat, err := ioutil.ReadFile(p)

	if err != nil {
		panic(err)
	}

	fmt.Println(f.Process(string(dat)))
}

// RunString is a function that receive a function f which implement
// runner and process the s string.
func RunString(f runner, s string) {
	fmt.Println(f.Process(s))
}
