package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/danil/goxsd"
)

//go:generate sh -c "go run . -p ncbisubmission ../../examples/ncbisubmission/submission-comb.xsd > ../../examples/ncbisubmission/ncbi_submission.go"
//go:generate sh -c "go run . -p f311sfc0v512 ../../examples/f311sfc0v512/SFC0_512.xsd > ../../examples/f311sfc0v512/f311_sfc0_512.go"
//go:generate sh -c "go run . -p mvk14201912 ../../examples/mvk14201912/mvk_1.4_201912.xsd > ../../examples/mvk14201912/mvk_1_4_201912.go"
//go:generate sh -c "go run . -p omu10 ../../examples/omu10/OMU_1.0.xsd > ../../examples/omu10/omu_1_0.go"

var (
	output, pckg, prefix string
	exported             bool

	usage = `Usage: goxsd [options] <xsd_file>

Options:
  -o <file>     Destination file [default: stdout]
  -p <package>  Package name [default: goxsd]
  -e            Generate exported structs [default: false]
  -x <prefix>   Struct name prefix [default: ""]

goxsd is a tool for generating XML decoding/encoding Go structs, according
to an XSD schema.
`
)

func main() {
	flag.StringVar(&output, "o", "", "Name of output file")
	flag.StringVar(&pckg, "p", "goxsd", "Name of the Go package")
	flag.StringVar(&prefix, "x", "", "Struct name prefix")
	flag.BoolVar(&exported, "e", false, "Generate exported structs")
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Println(usage)
		os.Exit(1)
	}
	xsdFile := flag.Arg(0)

	s, err := goxsd.ParseXSDFile(xsdFile)
	if err != nil {
		log.Fatal(err)
	}

	out := os.Stdout
	if output != "" {
		if out, err = os.Create(output); err != nil {
			fmt.Println("Could not create or truncate output file:", output)
			os.Exit(1)
		}
	}

	bldr := goxsd.NewBuilder(s)

	gen := goxsd.Generator{
		Package:  pckg,
		Prefix:   prefix,
		Exported: exported,
	}

	if err := gen.Do(out, bldr.BuildXML()); err != nil {
		fmt.Println("Code generation failed unexpectedly:", err.Error())
		os.Exit(1)
	}
}
