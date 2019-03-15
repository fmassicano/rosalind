package overlapgraphs

import (
	"bufio"
	"io"
	"regexp"
	"strings"

	"github.com/fmassicano/rosalind/utils"
)

// OverlapGraph is a struct to keep nodes and the links
type OverlapGraph struct {
	links []string
	nodes map[string][2]string // 0 contain prefix 1 contain sufix
}

func (og OverlapGraph) String() string {
	var str strings.Builder

	for _, v := range og.links {
		str.WriteString(v + "\n")
	}

	return str.String()
}

func (og *OverlapGraph) init() {
	// fmt.Println("init")
	og.links = make([]string, 1)
	og.nodes = make(map[string][2]string)
}

// Process is a function to run the process of this task
func (og OverlapGraph) Process(f string) string {
	// fmt.Println("Process")
	og.init()
	utils.Analyze(f, og.createNodes)
	og.createLinks()
	return og.String()
}

func (og *OverlapGraph) createNodes(handle io.Reader) error {
	// fmt.Println("createNodes")
	scanner := bufio.NewScanner(handle)
	var str strings.Builder
	var ID string
	firstMark := false

	for scanner.Scan() {
		matched, _ := regexp.MatchString(`>`, scanner.Text())

		if matched {

			if firstMark {
				og.fillNodesInfo(ID, &str)
			} else {
				firstMark = true
			}
			ID = strings.Trim(scanner.Text()[1:], " ")
			continue
		}

		str.WriteString(scanner.Text())

	}

	og.fillNodesInfo(ID, &str)

	return nil
}

func (og *OverlapGraph) fillNodesInfo(ID string, str *strings.Builder) {
	// fmt.Println("ID = ", ID)
	og.nodes[ID] = [2]string{str.String()[0:3], str.String()[len(str.String())-3 : len(str.String())]}
	str.Reset()
}

func (og *OverlapGraph) createLinks() {

	for keyFix, valueFix := range og.nodes {
		for keyVar, valueVar := range og.nodes {
			if keyVar != keyFix {
				if valueFix[1] == valueVar[0] {
					og.links = append(og.links, keyFix+" "+keyVar)
				}
			}
		}
	}

}
