// Things not yet implemented:
// - enforcing use="restricted" on attributes
// - namespaces

package goxsd

import (
	"strings"
)

// xmlTree is the representation of an XML element node in a tree. It
// contains information about whether
// - it is of a basic data type or a composite type (in which case its
//   type equals its name)
// - if it represents a list of children to its parent
// - if it has children of its own
// - any attributes
// - if the element contains any character data
type xmlTree struct {
	Name       string
	Type       string
	StructName string
	Annotation string
	List       bool
	Cdata      bool
	OmitEmpty  bool
	Attribs    []xmlAttrib
	Children   []*xmlTree
}

type xmlAttrib struct {
	Name      string
	Type      string
	OmitEmpty bool
}

type builder struct {
	schemas    []xsdSchema
	complTypes map[string]xsdComplexType
	simplTypes map[string]xsdSimpleType
	elements   map[string]struct{}
}

// NewBuilder creates a new initialized builder populated with the given
// xsdSchema slice.
func NewBuilder(schemas []xsdSchema) *builder {
	return &builder{
		schemas:    schemas,
		complTypes: make(map[string]xsdComplexType),
		simplTypes: make(map[string]xsdSimpleType),
		elements:   make(map[string]struct{}),
	}
}

// BuildXML generates and returns a tree of xmlTree objects based on a set of
// parsed XSD schemas.
func (b *builder) BuildXML() []*xmlTree {
	var roots []xsdElement
	for _, s := range b.schemas {
		roots = append(roots, s.Elements...)

		for _, t := range s.ComplexTypes {
			b.complTypes[t.Name] = t
		}

		for _, t := range s.SimpleTypes {
			b.simplTypes[t.Name] = t
		}
	}

	var xelems []*xmlTree
	for _, e := range roots {
		xelems = appendXMLElement(xelems, b.buildFromElement, e)
	}

	return xelems
}

func appendXMLElement(xmlElems []*xmlTree, build func(xsdElement) (*xmlTree, bool), xsdElem xsdElement) []*xmlTree {
	xmlElem, ok := build(xsdElem)
	if !ok {
		return xmlElems
	}
	return append(xmlElems, xmlElem)
}

// buildFromElement builds an xmlTree from an xsdElement, recursively
// traversing the XSD type information to build up an XML element hierarchy.
func (b *builder) buildFromElement(e xsdElement) (*xmlTree, bool) {
	if _, ok := b.elements[e.Name+e.Type]; ok {
		return nil, false
	}
	b.elements[e.Name+e.Type] = struct{}{}

	if e.Ref != "" {
		e.Name, e.Type = e.Ref, e.Ref
	}

	typ := e.Name
	if e.Type != "" {
		typ = e.Type
	}

	xelem := &xmlTree{Name: e.Name, Type: typ, Annotation: e.Annotation}

	if e.isList() {
		xelem.List = true
	}

	if e.omittable() {
		xelem.OmitEmpty = true
	}

	if !e.inlineType() {
		switch t := b.findType(e.Type).(type) {
		case xsdComplexType:
			b.buildFromComplexType(xelem, t)
		case xsdSimpleType:
			b.buildFromSimpleType(xelem, t)
		case string:
			b.buildFromSimpleType(xelem, xsdSimpleType{
				Restriction: xsdRestriction{Base: t},
			})
		}
		return xelem, true
	}

	if e.ComplexType != nil { // inline complex type
		b.buildFromComplexType(xelem, *e.ComplexType)
		return xelem, true
	}

	if e.SimpleType != nil { // inline simple type
		b.buildFromSimpleType(xelem, *e.SimpleType)
		return xelem, true
	}

	return xelem, true
}

// buildFromComplexType takes an xmlTree and an xsdComplexType, containing
// XSD type information for xmlTree enrichment.
func (b *builder) buildFromComplexType(xelem *xmlTree, t xsdComplexType) {
	if t.Choice != nil {
		for _, e := range t.Choice {
			e.Choice = true
			xelem.Children = appendXMLElement(xelem.Children, b.buildFromElement, e)
		}
	}

	if t.Sequence != nil { // Does the element have children?
		for _, e := range t.Sequence {
			xelem.Children = appendXMLElement(xelem.Children, b.buildFromElement, e)
		}
	}

	if t.SeqChoice != nil {
		for _, e := range t.SeqChoice {
			e.Choice = true
			xelem.Children = appendXMLElement(xelem.Children, b.buildFromElement, e)
		}
	}

	if t.Attributes != nil {
		b.buildFromAttributes(xelem, t.Attributes)
	}

	if t.ComplexContent != nil {
		b.buildFromComplexContent(xelem, *t.ComplexContent)
	}

	if t.SimpleContent != nil {
		b.buildFromSimpleContent(xelem, *t.SimpleContent)
	}
}

// buildFromSimpleType assumes restriction child and fetches the base value,
// assuming that value is of a XSD built-in data type.
func (b *builder) buildFromSimpleType(xelem *xmlTree, t xsdSimpleType) {
	xelem.Type = b.findType(t.Restriction.Base).(string)
}

func (b *builder) buildFromComplexContent(xelem *xmlTree, c xsdComplexContent) {
	if c.Extension != nil {
		b.buildFromExtension(xelem, c.Extension)
	}
}

// A simple content can refer to a text-only complex type
func (b *builder) buildFromSimpleContent(xelem *xmlTree, c xsdSimpleContent) {
	if c.Extension != nil {
		b.buildFromExtension(xelem, c.Extension)
	}

	if c.Restriction != nil {
		b.buildFromRestriction(xelem, c.Restriction)
	}
}

// buildFromExtension extends an existing type, simple or complex, with a
// sequence.
func (b *builder) buildFromExtension(xelem *xmlTree, e *xsdExtension) {
	switch t := b.findType(e.Base).(type) {
	case xsdComplexType:
		b.buildFromComplexType(xelem, t)
	case xsdSimpleType:
		b.buildFromSimpleType(xelem, t)
		// If element is of simpleType and has attributes, it must collect
		// its value as chardata.
		if e.Attributes != nil {
			xelem.Cdata = true
		}
	default:
		xelem.Type = t.(string)
		// If element is of built-in type but has attributes, it must collect
		// its value as chardata.
		if e.Attributes != nil {
			xelem.Cdata = true
		}
	}

	if e.Sequence != nil {
		for _, e := range e.Sequence {
			xelem.Children = appendXMLElement(xelem.Children, b.buildFromElement, e)
		}
	}

	if e.All != nil {
		for _, e := range e.All {
			xelem.Children = appendXMLElement(xelem.Children, b.buildFromElement, e)
		}
	}

	if e.Attributes != nil {
		b.buildFromAttributes(xelem, e.Attributes)
	}
}

func (b *builder) buildFromRestriction(xelem *xmlTree, r *xsdRestriction) {
	switch t := b.findType(r.Base).(type) {
	case xsdSimpleType:
		b.buildFromSimpleType(xelem, t)
	case xsdComplexType:
		b.buildFromComplexType(xelem, t)
	case xsdComplexContent:
		panic("Restriction on complex content is not implemented")
	default:
		panic("Unexpected base type to restriction")
	}
}

func (b *builder) buildFromAttributes(xelem *xmlTree, attrs []xsdAttribute) {
	for _, a := range attrs {
		attr := xmlAttrib{Name: a.Name}

		switch typ := b.findType(a.Type).(type) {
		case xsdSimpleType:
			// Get type name from simpleType
			// If Restriction.Base is a simpleType or complexType, we panic
			attr.Type = b.findType(typ.Restriction.Base).(string)

		case string:
			if a.SimpleType != nil {
				if t, ok := b.findType(a.SimpleType.Restriction.Base).(string); ok {
					typ = t
				}
			}

			// If empty, then simpleType is present as content, but we ignore that now.
			attr.Type = typ
		}

		if a.Use == "optional" {
			attr.OmitEmpty = true
		}

		xelem.Attribs = append(xelem.Attribs, attr)
	}
}

// findType takes a type name and checks if it is a registered XSD type
// (simple or complex), in which case that type is returned. If no such
// type can be found, the XSD specific primitive types are mapped to their
// Go correspondents. If no XSD type was found, the type name itself is
// returned.
func (b *builder) findType(name string) interface{} {
	name = stripNamespace(name)
	if t, ok := b.complTypes[name]; ok {
		return t
	}
	if t, ok := b.simplTypes[name]; ok {
		return t
	}

	switch name {
	case "boolean":
		return "bool"
	case "language", "Name", "token", "duration", "anyURI", "normalizedString":
		// FIXME: these types have additional constraints over string.
		// For example, a token cannot have leading/trailing whitespace.
		return "string"
	case "long", "short", "integer", "int":
		return "int"
	case "positiveInteger", "nonNegativeInteger":
		return "uint"
	case "unsignedShort":
		return "uint16"
	case "decimal":
		return "float64"
	case "dateTime", "date":
		return "time.Time"
	case "time":
		return "time.Duration"
	default:
		return name
	}
}

func stripNamespace(name string) string {
	if s := strings.Split(name, ":"); len(s) > 1 {
		return s[len(s)-1]
	}
	return name
}
