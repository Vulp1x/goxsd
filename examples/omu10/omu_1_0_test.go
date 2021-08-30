package omu10

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/danil/goxsd"
)

//go:embed omu_1_0.go
var code []byte

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

	want := string(code)

	if buf.String() != want {
		t.Errorf("want:\n%s\nget:\n%s", want, buf.String())
	}
}
