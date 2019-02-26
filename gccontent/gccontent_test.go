package gccontent

import (
	"math"
	"strings"
	"testing"
)

func TestGCcontent(t *testing.T) {
	var gc = GC{}
	var tests = []struct {
		s, idWant string
		pWant     float32
	}{
		{">Rosalind_6404\nCCTGCGGAAGATCGGCACTAGAATAGCCAGAACCGTTTCTCTGAGGCTTCCGGCCTTCCC\nTCCCACTAATAATTCTGAGG\n>Rosalind_5959\nCCATCGGTAGCGCATCCTTAGTCCAATTAAGTCCCTATCCAGGCGCTCCGCCGAAGGTCT\nATATCCATTTGTCAGCAGACACGC\n>Rosalind_0808\nCCACCCTCGTGGTATGGCTAGGCATTCAGGAACCGGAGAACGCTTCAGACCAGCCCGGAC\nTGGGAACCTGCGGGCAGTAGGTGGAAT",
			"Rosalind_0808",
			60.919540},
	}

	for _, c := range tests {
		gc.getAllGCcontent(strings.NewReader(c.s))
		gc.getTheHighestGCcount()
		gotID := gc.ID
		gotCGf := gc.gcf
		if gotID != c.idWant {
			t.Errorf("ID is %q want %q\n", gotID, c.idWant)
		}
		if math.Abs(float64(gotCGf-c.pWant)) >= 0.001 {
			t.Errorf("Percentage of CG is %v want %v\n", gotCGf, c.pWant)
		}
	}

}
