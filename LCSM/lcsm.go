package lcsm

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
	"sync"

	"github.com/fmassicano/rosalind-branchs/utils"
)

type LCSM struct {
	lcsm string
}

// CollectionLikelyLCSM is safe to use concurrently.
type CollectionLikelyLCSM struct {
	clcsm map[string]int
	mux   sync.Mutex
}

func (c *CollectionLikelyLCSM) inc(key string, init bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	if !init {
		if _, ok := c.clcsm[key]; ok {
			c.clcsm[key]++
		}
	} else {
		c.clcsm[key]++
	}

}

// Value returns the current value of the counter for the given key.
func (c *CollectionLikelyLCSM) value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.clcsm[key]
}

func (lm LCSM) String() string {

	if lm.lcsm == "" {
		return fmt.Sprintf("None value")
	}

	return fmt.Sprintf("%v", lm.lcsm)
}

var dnaStrings = []string{} // It is used at extractDNAstrings

func (lm LCSM) Process(f string) string {

	log.SetFlags(log.Ltime)

	cll := CollectionLikelyLCSM{clcsm: make(map[string]int)}

	utils.Analyze(f, extractDNAstrings)

	cll.getLikelyLCSM(dnaStrings)

	lm.lcsm = cll.getLCSM(len(dnaStrings))

	fmt.Println(len(cll.clcsm))

	return lm.String()
}

func extractDNAstrings(handle io.Reader) error {

	scanner := bufio.NewScanner(handle)
	var str strings.Builder
	firstMark := true

	for scanner.Scan() {

		matched, _ := regexp.MatchString(`>`, scanner.Text())

		if matched {
			fmt.Println(scanner.Text())
			if firstMark {
				firstMark = !firstMark
			} else {
				dnaStrings = append(dnaStrings, str.String())
				str.Reset()
			}
			continue
		}

		str.WriteString(scanner.Text())

	}

	dnaStrings = append(dnaStrings, str.String())

	return nil
}

func (c *CollectionLikelyLCSM) findLikelyLCSM(s string, size int, init bool) {

	var wg sync.WaitGroup
	tasks := make(chan int)

	for worker := 0; worker < 1000; worker++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for i := range tasks {
				c.inc(s[i:(i+size)], init)
			}
		}()

	}

	n := len(s) - (size - 1)
	for i := 0; i < n; i++ {
		tasks <- i
	}
	close(tasks)
	wg.Wait()
}

func (c *CollectionLikelyLCSM) xpto(i int, v string, init bool) {
	log.Println(i, len(v))
	for size := 2; size <= len(v); size++ {
		c.findLikelyLCSM(v, size, init)
	}
}

func (c *CollectionLikelyLCSM) getLikelyLCSM(ds []string) {
	init := true
	for i, v := range ds {
		c.xpto(i, v, init)
		if init {
			init = false
		}
	}
}

func (c *CollectionLikelyLCSM) getLCSM(n int) string {
	maxSize := 0
	var lcsm string

	for key, v := range c.clcsm {
		if len(key) > maxSize && v == n {
			lcsm, maxSize = key, len(key)
		}
	}
	return lcsm
}
