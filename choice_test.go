package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/kr/pretty"
)

/*type testCase struct {
	xsd   string
	xml   xmlTree
	gosrc string
}*/

var (
	tests2 = []struct {
		exported bool
		prefix   string
		xsd      string
		xml      xmlTree
		gosrc    string
	}{
		{
			exported: false,
			prefix:   "",
			xsd: `<?xml version="1.0" encoding="UTF-8"?>
			<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:com="SP.common"
			    targetNamespace="SP.common" version="2.0">

			    <xs:element name="LocalId" type="typeLocalId"/>
				<!--  Local identifier in submission context. -->
				<xs:complexType name="typeLocalId">
				<xs:annotation>
				    <xs:documentation> Local identifier in submission context. </xs:documentation>
				</xs:annotation>
				<xs:simpleContent>
				    <xs:extension base="xs:string">
					<xs:attribute type="xs:string" name="submission_id" use="optional">
					    <xs:annotation>
						<xs:documentation> Optional submission id. If omitted, the current
						    submission is assumed. </xs:documentation>
					    </xs:annotation>
					</xs:attribute>
				    </xs:extension>
				</xs:simpleContent>
				</xs:complexType>

			</xs:schema>`,

			xml: xmlTree{
				Name:     "LocalId",
				Type:     "string",
				List:     false,
				Cdata:    true,
				Attribs:  []xmlAttrib{{Name: "submission_id", Type: "string"}},
				Children: nil,
			},
			gosrc: `
type LocalID struct {
	SubmissionID string ` + "`xml:\"submission_id,attr,omitempty\"`" + `
	Value    string ` + "`xml:\",chardata\"`" + `
}

				`,
		},
		{
			exported: false,
			prefix:   "",
			xsd: `<?xml version="1.0" encoding="UTF-8"?>
			<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:com="SP.common"
			    targetNamespace="SP.common" version="2.0">

			    <xs:element name="RefId" type="typeRefId"/>
				<!--  Local identifier in submission context. -->
				<xs:complexType name="typeLocalId">
				<xs:annotation>
				    <xs:documentation> Local identifier in submission context. </xs:documentation>
				</xs:annotation>
				<xs:simpleContent>
				    <xs:extension base="xs:string">
					<xs:attribute type="xs:string" name="submission_id" use="optional">
					    <xs:annotation>
						<xs:documentation> Optional submission id. If omitted, the current
						    submission is assumed. </xs:documentation>
					    </xs:annotation>
					</xs:attribute>
				    </xs:extension>
				</xs:simpleContent>
				</xs:complexType>

			    <!-- Reference to a record inside NCBI database or in Submission Portal. -->
			    <xs:complexType name="typeRefId">
				<xs:annotation>
				    <xs:documentation> Reference to a record inside NCBI database or in Submission Portal.
				    </xs:documentation>
				</xs:annotation>
				<xs:choice>
				    <xs:element name="LocalId" type="typeLocalId" minOccurs="0" maxOccurs="1">
					<xs:annotation>
					    <xs:documentation> Local id in submission context. </xs:documentation>
					</xs:annotation>
				    </xs:element>
				    <xs:element name="SPUID" type="com:typeSPUID" minOccurs="0" maxOccurs="1">
					<xs:annotation>
					    <xs:documentation> User-supplied unique id. </xs:documentation>
					</xs:annotation>
				    </xs:element>
				    <xs:element name="PrimaryId" type="com:typePrimaryId" minOccurs="0" maxOccurs="1">
					<xs:annotation>
					    <xs:documentation> NCBI accession. </xs:documentation>
					</xs:annotation>
				    </xs:element>
				</xs:choice>
			    </xs:complexType>
			    </xs:schema>`,
			xml: xmlTree{
				Name:     "RefId",
				Type:     "RefId",
				Children: nil,
			},
			gosrc: `
type RefId struct {
	choice
}

				`,
		},
	}
)

/*func removeComments(buf bytes.Buffer) bytes.Buffer {
	lines := strings.Split(buf.String(), "\n")
	for i, l := range lines {
		if strings.HasPrefix(l, "//") {
			lines = append(lines[:i], lines[i+1:]...)
		}
	}
	return *bytes.NewBufferString(strings.Join(lines, "\n"))
}*/

func TestBuildXmlElem2(t *testing.T) {
	for i, tst := range tests2 {
		schemas, err := parse(strings.NewReader(tst.xsd), "test")
		if err != nil {
			t.Fatal(err)
		}

		if testing.Verbose() {
			t.Logf("\n%s", pretty.Sprint(schemas))
		}

		if len(schemas) != 1 {
			t.Errorf("%d. len(parse) = %d, want %d", i, len(schemas), 1)
			continue
		}

		bldr := newBuilder(schemas)
		elems := bldr.buildXML()
		if testing.Verbose() {
			t.Log(pretty.Sprint(elems))
		}
		if len(elems) != 1 {
			t.Errorf("%d. len(bldr.buildXML) = %d, want %d", i, len(elems), 1)
			continue
		}
		e := elems[0]
		if !reflect.DeepEqual(tst.xml, *e) {
			if testing.Verbose() {
				t.Error(pretty.Sprintf("%d. bldr.buildXML() => %# v, want %# v", i, *e, tst.xml))
			} else {
				t.Errorf("%d. Unexpected XML element: %s", i, e.Name)
			}
			//pretty.Println(tst.xml)
			//pretty.Println(e)
		}
	}
}

func TestGenerateGo2(t *testing.T) {
	for i, tst := range tests2 {
		var out bytes.Buffer
		g := generator{prefix: tst.prefix, exported: tst.exported}
		g.do(&out, []*xmlTree{&tst.xml})
		out = removeComments(out)
		if strings.Join(strings.Fields(out.String()), "") != strings.Join(strings.Fields(tst.gosrc), "") {
			//t.Errorf("%d. have:\n%s\nwant:\n%sUnexpected generated Go source: %s", i, tst.xml.Name)
			t.Errorf("%d.\nhave:\n%s\nwant:\n%s", i, out.Bytes(), tst.gosrc)
			//t.Logf(out.String())
			//t.Logf(strings.Join(strings.Fields(out.String()), ""))
			//t.Logf(strings.Join(strings.Fields(tst.gosrc), ""))
		}
	}
}
