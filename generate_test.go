package goxsd

import (
	"testing"

	"github.com/alexsergivan/transliterator"
)

func TestLint(t *testing.T) {
	for i, tt := range []struct {
		input, want string
	}{
		{"string", "string"},
		{"int64", "int64"},
		{"time.Time", "time.Time"},
		{"time.Time", "time.Time"},
		{"foo cpu baz", "FooCPUBaz"},
		{"foo Cpu baz", "FooCPUBaz"},
		{"test Id", "TestID"},
		{"json and html", "JSONAndHTML"},
		{"Json and Html", "JSONAndHTML"},
	} {
		translate := func(s string) string { return s }
		if got := lint(tt.input, translate); got != tt.want {
			t.Errorf("[%d] title(%q) = %q, want %q", i, tt.input, got, tt.want)
		}
	}
}

func TestLintTransliterator(t *testing.T) {
	for i, tt := range []struct {
		input, want string
	}{
		{"bäume", "Baeume"},
	} {
		trans := transliterator.NewTransliterator(nil)
		translate := func(s string) string { return trans.Transliterate(s, "de") }
		if got := lint(tt.input, translate); got != tt.want {
			t.Errorf("[%d] title(%q) = %q, want %q", i, tt.input, got, tt.want)
		}
	}
}

// func TestLintTitle(t *testing.T) {
// 	for i, tt := range []struct {
// 		input, want string
// 	}{
// 		{"foo cpu baz", "FooCPUBaz"},
// 		{"test Id", "TestID"},
// 		{"json and html", "JSONAndHTML"},
// 	} {
// 		if got := lintTitle(tt.input); got != tt.want {
// 			t.Errorf("[%d] title(%q) = %q, want %q", i, tt.input, got, tt.want)
// 		}
// 	}
// }
