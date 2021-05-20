package goxsd

import (
	"testing"
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
		{"test Id", "TestID"},
		{"json and html", "JSONAndHTML"},
	} {
		if got := lint(tt.input); got != tt.want {
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
