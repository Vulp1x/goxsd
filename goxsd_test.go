package goxsd

import (
	"bytes"
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kr/pretty"
)

var (
	tests = []struct {
		line     int
		exported bool
		prefix   string
		xsd      string
		xml      xmlTree
		gosrc    string
	}{
		{
			line:     line(),
			exported: false,
			prefix:   "",
			xsd: `<schema>
	<element name="titleList" type="titleListType">
	</element>
	<complexType name="titleListType">
		<sequence>
			<element ref="titleType" minOccurs="0" />
			<element name="title" type="originalTitleType" maxOccurs="unbounded" />
		</sequence>
	</complexType>
	<complexType name="originalTitleType">
		<simpleContent>
			<extension base="titleType">
				<attribute name="original" type="boolean">
				</attribute>
			</extension>
		</simpleContent>
	</complexType>
	<complexType name="titleType">
		<simpleContent>
			<restriction base="textType">
				<maxLength value="300" />
			</restriction>
		</simpleContent>
	</complexType>
	<complexType name="textType">
		<simpleContent>
			<extension base="string">
				<attribute name="language" type="language">
				</attribute>
			</extension>
		</simpleContent>
	</complexType>
</schema>`,
			xml: xmlTree{
				Name: "titleList",
				Type: "titleListType",
				Children: []*xmlTree{
					&xmlTree{
						Name:      "titleType",
						Type:      "StringXML",
						Cdata:     true,
						OmitEmpty: true,
						Attribs: []xmlAttrib{
							{Name: "language", Type: "LanguageXML"},
						},
					},
					&xmlTree{
						Name:  "title",
						Type:  "StringXML",
						Cdata: true,
						List:  true,
						Attribs: []xmlAttrib{
							{Name: "language", Type: "LanguageXML"},
							{Name: "original", Type: "BooleanXML"},
						},
					},
				},
			},
			gosrc: `
	// TitleListType is generated from an XSD element.
	type TitleListType struct {
		TitleType *TitleType ` + "`xml:\"titleType,omitempty\"`" + `
		Title     []Title    ` + "`xml:\"title\"`" + `
	}

	// TitleType is generated from an XSD element.
	type TitleType struct {
		Language  LanguageXML ` + "`xml:\"language,attr\"`" + `
		TitleType StringXML   ` + "`xml:\",cdata\"`" + `
	}

	// Title is generated from an XSD element.
	type Title struct {
		Language LanguageXML ` + "`xml:\"language,attr\"`" + `
		Original BooleanXML  ` + "`xml:\"original,attr\"`" + `
		Title    StringXML   ` + "`xml:\",cdata\"`" + `
	}`,
		},

		{
			line:     line(),
			exported: false,
			prefix:   "",
			xsd: `<schema>
	<element name="tagList">
		<complexType>
			<sequence>
				<element name="tag" type="tagReferenceType" minOccurs="0" maxOccurs="unbounded" />
			</sequence>
		</complexType>
	</element>
	<complexType name="tagReferenceType">
		<simpleContent>
			<extension base="nidType">
				<attribute name="type" type="tagTypeType" use="required" />
			</extension>
		</simpleContent>
	</complexType>
	<simpleType name="nidType">
		<restriction base="string">
			<pattern value="[0-9a-zA-Z\-]+" />
		</restriction>
	</simpleType>
	<simpleType name="tagTypeType">
		<restriction base="string">
		</restriction>
	</simpleType>
</schema>`,
			xml: xmlTree{
				Name: "tagList",
				Type: "tagList",
				Children: []*xmlTree{
					&xmlTree{
						Name:      "tag",
						Type:      "StringXML",
						List:      true,
						Cdata:     true,
						OmitEmpty: true,
						Attribs: []xmlAttrib{
							{Name: "type", Type: "StringXML"},
						},
					},
				},
			},
			gosrc: `
	// TagList is generated from an XSD element.
	type TagList struct {
		Tag []Tag ` + "`xml:\"tag,omitempty\"`" + `
	}

	// Tag is generated from an XSD element.
	type Tag struct {
		Type StringXML ` + "`xml:\"type,attr\"`" + `
		Tag  StringXML ` + "`xml:\",cdata\"`" + `
	}`,
		},

		{
			line:     line(),
			exported: false,
			prefix:   "",
			xsd: `<schema>
				<element name="tagId" type="tagReferenceType" />
	<complexType name="tagReferenceType">
		<simpleContent>
			<extension base="string">
				<attribute name="type" type="string" use="required" />
			</extension>
		</simpleContent>
	</complexType>
</schema>`,
			xml: xmlTree{
				Name:  "tagId",
				Type:  "StringXML",
				List:  false,
				Cdata: true,
				Attribs: []xmlAttrib{
					{Name: "type", Type: "StringXML"},
				},
			},
			gosrc: `
	// TagID is generated from an XSD element.
	type TagID struct {
		Type  StringXML ` + "`xml:\"type,attr\"`" + `
		TagID StringXML ` + "`xml:\",cdata\"`" + `
	}`,
		},

		{
			line:     line(),
			exported: true,
			prefix:   "xxx",
			xsd: `<schema>
	<element name="url" type="tagReferenceType" />
	<complexType name="tagReferenceType">
		<simpleContent>
			<extension base="string">
				<attribute name="type" type="string" use="required" />
			</extension>
		</simpleContent>
	</complexType>
</schema>`,
			xml: xmlTree{
				Name:  "url",
				Type:  "StringXML",
				List:  false,
				Cdata: true,
				Attribs: []xmlAttrib{
					{Name: "type", Type: "StringXML"},
				},
			},
			gosrc: `
	// XxxURL is generated from an XSD element.
	type XxxURL struct {
		Type StringXML ` + "`xml:\"type,attr\"`" + `
		URL  StringXML ` + "`xml:\",cdata\"`" + `
	}`,
		},

		// Windows-1252 encoding
		{
			line: line(),
			xsd: `<?xml version="1.0" encoding="Windows-1252"?>
	<schema>
		<element name="empty" type="tagReferenceType" />
		<complexType name="tagReferenceType"/>
	</schema>`,
			xml: xmlTree{
				Name: "empty",
				Type: "tagReferenceType",
			},
			gosrc: `
	// TagReferenceType is generated from an XSD element.
	type TagReferenceType struct{}`,
		},
	}
)

func removeComments(buf bytes.Buffer) bytes.Buffer {
	lines := strings.Split(buf.String(), "\n")
	for i, l := range lines {
		if strings.HasPrefix(strings.TrimSpace(l), "//") {
			lines = append(lines[:i], lines[i+1:]...)
		}
	}
	return *bytes.NewBufferString(strings.Join(lines, "\n"))
}

func TestBuildXmlElem(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)

	for _, tt := range tests {
		tt := tt
		t.Run(strconv.Itoa(tt.line), func(t *testing.T) {
			////////////////////////////////////////////////////////////////////////////////
			// t.Parallel()
			link := fmt.Sprintf("%s:%d", testFile, tt.line)

			schemas, err := parse(strings.NewReader(tt.xsd), "test")
			if err != nil {
				t.Fatal(err)
			}

			if testing.Verbose() {
				t.Logf("\n%s", pretty.Sprint(schemas))
			}

			if len(schemas) != 1 {
				t.Fatalf("len(parse(schema)) = %d, want %d, test: %s", len(schemas), 1, link)
			}

			bldr := NewBuilder(schemas)
			elems := bldr.BuildXML()
			if testing.Verbose() {
				t.Log(pretty.Sprint(elems))
			}
			if len(elems) != 1 {
				t.Fatalf("len(bldr.buildXML) = %d, want %d, test: %s", len(elems), 1, link)
			}
			e := elems[0]
			if !reflect.DeepEqual(tt.xml, *e) {
				if testing.Verbose() {
					t.Error(pretty.Sprintf("bldr.buildXML() => %# v, want %# v, test: %s", *e, tt.xml, link))
				} else {
					t.Errorf("Unexpected XML element: %s, test: %s", e.Name, link)
				}
			}
		})
	}
}

func TestGenerateGo(t *testing.T) {
	_, testFile, _, _ := runtime.Caller(0)

	for _, tt := range tests {
		tt := tt
		t.Run(strconv.Itoa(tt.line), func(t *testing.T) {
			////////////////////////////////////////////////////////////////////////////////
			// t.Parallel()
			link := fmt.Sprintf("%s:%d", testFile, tt.line)

			var out bytes.Buffer
			g := Generator{Prefix: tt.prefix, Exported: tt.exported}

			err := g.Do(&out, []*xmlTree{&tt.xml})
			if err != nil {
				t.Fatalf("generator do: %s, test: %s", err, link)
			}

			want := strings.TrimSpace(tt.gosrc)
			get := strings.TrimSpace(out.String())

			if want != get {
				t.Errorf("\ncode diff:\n%s\nwant:\n%s\nget:\n%s\ntest: %s", cmp.Diff(want, get), want, get, link)
			}
		})
	}
}

func TestParseSubmissionPortal(t *testing.T) {
	schema, err := ParseXSDFile("testdata/SP.common.xsd")
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.Sprint(schema))
}

func line() int { _, _, l, _ := runtime.Caller(1); return l }
