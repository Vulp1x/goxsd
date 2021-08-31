package mvk14201912_test

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/danil/goxsd"
	"github.com/google/go-cmp/cmp"
)

//go:embed mvk_1_4_201912.go
var code []byte

func TestGenerate(t *testing.T) {
	schm, err := goxsd.ParseXSDFile("mvk_1.4_201912.xsd")
	if err != nil {
		t.Fatal(err)
	}

	bldr := goxsd.NewBuilder(schm)
	gen := goxsd.Generator{Package: "mvk14201912"}

	var buf bytes.Buffer

	err = gen.Do(&buf, bldr.BuildXML())
	if err != nil {
		t.Fatal(err)
	}

	want := string(code)
	get := buf.String()

	if want != get {
		t.Errorf("\ncode diff:\n%s\nwant:\n%s\nget:\n%s", cmp.Diff(want, get), want, get)
	}
}
