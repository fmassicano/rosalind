package findingamotifindna

import (
	"fmt"
	"strings"
)

///// Error Handling //////
type MotifError struct {
	msg string
}

func (error MotifError) Error() string {
	return error.msg
}

type Motif struct {
	s, t    string
	locTinS []int
}

///////////////////////////

func (m Motif) String() string {

	var strr string
	for i := 0; i < len(m.locTinS); i++ {
		strr = strr + fmt.Sprintf("%v ", m.locTinS[i])
	}
	strr = strr + fmt.Sprintf("\n")
	return strr
}

func (m Motif) Process(s string) string {
	res, err := m.ProcessErr(s)
	if err != nil {
		return err.Error()
	}
	return res
}

func (m Motif) ProcessErr(s string) (string, error) {
	err := m.init(s)
	if err != nil {
		return "", err
	}
	m.getMotifPositions()
	// fmt.Println(m.locTinS)
	return m.String(), nil
}

func (m *Motif) init(s string) error {
	sr := strings.Split(s, "\n")
	m.s = sr[0]
	m.t = sr[1]
	if len(m.s) < len(m.t) {
		return &MotifError{"The length of 's' needs to be bigger or equal than t."}
	}
	// fmt.Println(m.s)
	// fmt.Println(m.t)
	return nil
}

func (m *Motif) getMotifPositions() {

	for start, end := 0, len(m.t); end < len(m.s); start, end = start+1, end+1 {
		// fmt.Println(start, end)
		if m.s[start:end] == m.t {
			// fmt.Println(start + 1)
			m.locTinS = append(m.locTinS, start+1)
		}
	}

}
