package f311sfc0v512_test

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/alexsergivan/transliterator"
	"github.com/danil/goxsd"
	"github.com/google/go-cmp/cmp"
)

//go:embed f311_sfc0_512.go
var code []byte

func TestGenerate(t *testing.T) {
	schm, err := goxsd.ParseXSDFile("SFC0_512.xsd")
	if err != nil {
		t.Fatal(err)
	}

	bldr := goxsd.NewBuilder(schm)
	trans := transliterator.NewTransliterator(nil)
	translate := func(s string) string { return trans.Transliterate(s, "ru") }
	gen := goxsd.Generator{Package: "f311sfc0v512", Translator: translate}

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
