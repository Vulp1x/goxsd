package omu10

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/danil/goxsd"
	"github.com/google/go-cmp/cmp"
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
	get := buf.String()

	if want != get && want != "" && get != "" {
		t.Errorf("\ncode diff:\n%s", cmp.Diff(want, get))

	} else if want != get {
		t.Errorf(
			"\ncode diff:\n%s\nwant:\n%s\nget:\n%s",
			cmp.Diff(want, get), want, get,
		)
	}
}
