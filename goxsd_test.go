package goxsd

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/kr/pretty"
)

var (
	tests = []struct {
		exported bool
		prefix   string
		xsd      string
		xml      xmlTree
		gosrc    string
	}{
		{
			exported: false,
			prefix:   "",
			xsd: `<schema>
	<element name="titleList" type="titleListType">
	</element>
	<complexType name="titleListType">
		<sequence>
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
				Type: "titleList",
				Children: []*xmlTree{
					&xmlTree{
						Name:  "title",
						Type:  "string",
						Cdata: true,
						List:  true,
						Attribs: []xmlAttrib{
							{Name: "language", Type: "string"},
							{Name: "original", Type: "bool"},
						},
					},
				},
			},
			gosrc: `
	type TitleList struct {
		Title []title ` + "`xml:\"title\"`" + `
	}

	type Title struct {
		Language string ` + "`xml:\"language,attr\"`" + `
		Original bool   ` + "`xml:\"original,attr\"`" + `
		Value    string ` + "`xml:\",chardata\"`" + `
	}`,
		},

		{
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
						Name:  "tag",
						Type:  "string",
						List:  true,
						Cdata: true,
						Attribs: []xmlAttrib{
							{Name: "type", Type: "string"},
						},
					},
				},
			},
			gosrc: `
	type TagList struct {
		Tag []tag ` + "`xml:\"tag\"`" + `
	}

	type Tag struct {
		Type  string ` + "`xml:\"type,attr\"`" + `
		Value string ` + "`xml:\",chardata\"`" + `
	}`,
		},

		{
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
				Type:  "string",
				List:  false,
				Cdata: true,
				Attribs: []xmlAttrib{
					{Name: "type", Type: "string"},
				},
			},
			gosrc: `
	type TagID struct {
		Type  string ` + "`xml:\"type,attr\"`" + `
		Value string ` + "`xml:\",chardata\"`" + `
	}`,
		},

		{
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
				Type:  "string",
				List:  false,
				Cdata: true,
				Attribs: []xmlAttrib{
					{Name: "type", Type: "string"},
				},
			},
			gosrc: `
	type XxxURL struct {
		Type  string ` + "`xml:\"type,attr\"`" + `
		Value string ` + "`xml:\",chardata\"`" + `
	}`,
		},

		// Windows-1252 encoding
		{
			xsd: `<?xml version="1.0" encoding="Windows-1252"?>
	<schema>
		<element name="empty" type="tagReferenceType" />
		<complexType name="tagReferenceType"/>
	</schema>`,
			xml: xmlTree{
				Name: "empty",
				Type: "empty",
			},
			gosrc: `
	type Empty struct {}`,
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
	for i, tst := range tests {
		schemas, err := parse(strings.NewReader(tst.xsd), "test")
		if err != nil {
			t.Fatal(err)
		}

		if testing.Verbose() {
			t.Logf("\n%s", pretty.Sprint(schemas))
		}

		if len(schemas) != 1 {
			t.Errorf("%d. len(parse(schema)) = %d, want %d", i, len(schemas), 1)
			continue
		}

		bldr := NewBuilder(schemas)
		elems := bldr.BuildXML()
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
				t.Error(pretty.Sprintf("%d.bldr.buildXML() => %# v, want %# v", i, *e, tst.xml))
			} else {
				t.Errorf("%d. Unexpected XML element: %s", i, e.Name)
			}
		}
	}
}

func TestGenerateGo(t *testing.T) {
	for i, tst := range tests {
		var out bytes.Buffer
		g := Generator{Prefix: tst.prefix, Exported: tst.exported}
		err := g.Do(&out, []*xmlTree{&tst.xml})
		if err != nil {
			t.Errorf("%d. generator do: %s", i, err)
			continue
		}
		out = removeComments(out)
		if strings.Join(strings.Fields(out.String()), "") != strings.Join(strings.Fields(tst.gosrc), "") {
			t.Errorf("%d.\nhave:\n%s\nwant:\n%s", i, out.String(), tst.gosrc)
			continue
		}
	}
}

/*func TestSquish(t *testing.T) {
	for i, tt := range []struct {
		input, want string
	}{
		{"Foo CPU Baz", "FooCPUBaz"},
		{"Test ID", "TestID"},
		{"JSON And HTML", "JSONAndHTML"},
	} {
		if got := squish(tt.input); got != tt.want {
			t.Errorf("[%d] squish(%q) = %q, want %q", i, tt.input, got, tt.want)
		}
	}
}*/

/*func TestReplace(t *testing.T) {
	for i, tt := range []struct {
		input, want string
	}{
		{"foo Cpu baz", "foo CPU baz"},
		{"test Id", "test ID"},
		{"Json and Html", "JSON and HTML"},
	} {
		if got := initialisms.Replace(tt.input); got != tt.want {
			t.Errorf("[%d] replace(%q) = %q, want %q", i, tt.input, got, tt.want)
		}
	}

	c := len(initialismPairs)

	for i := 0; i < c; i++ {
		input, want := initialismPairs[i], initialismPairs[i+1]

		if got := initialisms.Replace(input); got != want {
			t.Errorf("[%d] replace(%q) = %q, want %q", i, input, got, want)
		}

		i++
	}
}*/

func TestParseSubmissionPortal(t *testing.T) {
	schema, err := ParseXSDFile("testdata/SP.common.xsd")
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.Sprint(schema))
}
