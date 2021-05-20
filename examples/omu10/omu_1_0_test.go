package omu10

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/danil/goxsd"
)

//go:embed omu_1_0.go
var sourceCode []byte

func TestGenerate(t *testing.T) {
	schm, err := goxsd.ParseXSDFile("OMU_1.0.xsd")
	if err != nil {
		t.Fatal(err)
	}

	bldr := goxsd.NewBuilder(schm)
	gen := goxsd.Generator{Package: "omu10"}

	var buf bytes.Buffer

	err = gen.Do(&buf, bldr.BuildXML())
	if err != nil {
		t.Fatal(err)
	}

	expected := string(sourceCode)

	if buf.String() != expected {
		t.Errorf("expected:\n%s\nreceived:\n%s", expected, buf.String())
	}
}
