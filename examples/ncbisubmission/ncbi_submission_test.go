package ncbisubmission_test

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/danil/goxsd"
)

//go:embed ncbi_submission.go
var code []byte

func TestGenerate(t *testing.T) {
	schm, err := goxsd.ParseXSDFile("submission-comb.xsd")
	if err != nil {
		t.Fatal(err)
	}

	bldr := goxsd.NewBuilder(schm)
	gen := goxsd.Generator{Package: "ncbisubmission"}

	var buf bytes.Buffer

	err = gen.Do(&buf, bldr.BuildXML())
	if err != nil {
		t.Fatal(err)
	}

	want := string(code)

	if buf.String() != want {
		t.Errorf("want:\n%s\nget:\n%s", want, buf.String())
	}
}
