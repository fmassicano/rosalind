package runner

import (
	"fmt"
	"io/ioutil"
)

type runner interface {
	Process(s string) string
}

func Run(f runner, p string) {
	dat, err := ioutil.ReadFile(p)

	if err != nil {
		panic(err)
	}

	fmt.Println(f.Process(string(dat)))
}
