package main

import (
	"testing"

	"github.com/kr/pretty"
)

func TestParseNcbiSubmissionProtocol(t *testing.T) {
	//schema, err := parseXSDFile("testdata/SP.common.xsd")
	schema, err := parseXSDFile("testdata/submission-comb.xsd")
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.Sprint(schema))
}
