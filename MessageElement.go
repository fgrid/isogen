package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/template"
)

const (
	complexAccessTemplate = `
func ({{.Receiver}} *{{.ReceiverType}}) Add{{.Element}}() *{{.ElementType}} {
	{{.Receiver}}.{{.Element}} = new({{.ElementType}})
	return {{.Receiver}}.{{.Element}}
}
`
	complexArrayAccessTemplate = `
func ({{.Receiver}} *{{.ReceiverType}}) Add{{.Element}}() *{{.ElementType}} {
	newValue := new ({{.ElementType}})
	{{.Receiver}}.{{.Element}} = append({{.Receiver}}.{{.Element}}, newValue)
	return newValue
}
`
	simpleAccessTemplate = `
func ({{.Receiver}} *{{.ReceiverType}}) Set{{.Element}}(value string) {
	{{.Receiver}}.{{.Element}} = (*{{.ElementType}})(&value)
}
`
	simpleArrayAccessTemplate = `
func ({{.Receiver}} *{{.ReceiverType}}) Add{{.Element}}(value string) {
	{{.Receiver}}.{{.Element}} = append({{.Receiver}}.{{.Element}}, (*{{.ElementType}})(&value))
}
`
	amountAccessTemplate = `
func ({{.Receiver}} *{{.ReceiverType}}) Set{{.Element}}(value, currency string) {
	{{.Receiver}}.{{.Element}} = New{{.ElementType}}(value, currency)
}
`
	amountArrayAccessTemplate = `
func ({{.Receiver}} *{{.ReceiverType}}) Add{{.Element}}(value, currency string) {
	{{.Receiver}}.{{.Element}} = append({{.Receiver}}.{{.Element}}, New{{.ElementType}}(value, currency))
}
`
)

var (
	complexAccess      *template.Template
	complexArrayAccess *template.Template
	simpleAccess       *template.Template
	simpleArrayAccess  *template.Template
	amountAccess       *template.Template
	amountArrayAccess  *template.Template
)

func init() {
	complexAccess = prepareTemplate("complexAccess", complexAccessTemplate)
	complexArrayAccess = prepareTemplate("complexArrayAccess", complexArrayAccessTemplate)
	simpleAccess = prepareTemplate("simpleAccess", simpleAccessTemplate)
	simpleArrayAccess = prepareTemplate("simpleArrayAccess", simpleArrayAccessTemplate)
	amountAccess = prepareTemplate("amountAccess", amountAccessTemplate)
	amountArrayAccess = prepareTemplate("amountArrayAccess", amountArrayAccessTemplate)
}

func prepareTemplate(name, body string) *template.Template {
	tmpl, err := template.New(name).Parse(body)
	if err != nil {
		log.Fatalf("could not compile template %s - %s", name, err.Error())
	}
	return tmpl
}

type MessageElement struct {
	XSIType     string `xml:"xsitype,attr"`
	XMIId       string `xml:"http://www.omg.org/XMI id,attr"`
	Definition  string `xml:"definition,attr"`
	Name        string `xml:"name,attr"`
	MaxOccurs   string `xml:"maxOccurs,attr"`
	MinOccurs   string `xml:"minOccurs,attr"`
	XMLTag      string `xml:"xmlTag,attr"`
	SimpleType  string `xml:"simpleType,attr"`
	ComplexType string `xml:"complexType,attr"`
	Type        string `xml:"type,attr"`
}

func (m *MessageElement) typeID() (typeID string, complex bool) {
	if len(m.SimpleType) > 0 {
		return m.SimpleType, false
	}
	if len(m.ComplexType) > 0 {
		return m.ComplexType, true
	}
	if len(m.Type) > 0 {
		return m.Type, true
	}
	panic(fmt.Sprintf("message element with undefined type: %+v", m))
}

func (m *MessageElement) IsArray() bool {
	return m.MaxOccurs == "" || (m.MaxOccurs != "0" && m.MaxOccurs != "1")
}

func (m *MessageElement) ArrayDeclaration() string {
	if m.IsArray() {
		return "[]"
	}
	return ""
}

func (m *MessageElement) Declaration() string {
	return fmt.Sprintf(
		"// %s\n\t%s %s*%s `xml:\"%s%s\"`",
		strings.Replace(m.Definition, "\n", "\n\t// ", -1),
		m.Name, m.ArrayDeclaration(), m.MemberType(), m.XMLTag, m.optional())
}

func (m *MessageElement) DeclarationOut(basePackageName string) string {
	return fmt.Sprintf(
		"// %s\n\t%s %s*%s.%s `xml:\"%s%s\"`",
		strings.Replace(m.Definition, "\n", "\n\t// ", -1),
		m.Name, m.ArrayDeclaration(), basePackageName, m.MemberType(), m.XMLTag, m.optional())
}

type context struct {
	Receiver       string
	ReceiverType   string
	Element        string
	ElementType    string
	ElementXSIType string
}

func (m *MessageElement) Access(basePackageName, receiverType string) string {
	c := &context{
		Receiver:     strings.ToLower(receiverType[:1]),
		ReceiverType: receiverType,
		Element:      m.Name,
	}
	typeID, complex := m.typeID()
	t := typeMap[typeID]
	c.ElementXSIType = t.XSIType
	c.ElementType = t.Name
	if basePackageName != "" {
		c.ElementType = basePackageName + "." + c.ElementType
	}
	tmpl := chooseTemplate(complex, m.IsArray(), c.ElementXSIType == "iso20022:Amount")
	return m.buildAccess(c, tmpl)
}

func chooseTemplate(complex, array, amount bool) *template.Template {
	if complex {
		return choose(array, complexArrayAccess, complexAccess)
	}
	if amount {
		return choose(array, amountArrayAccess, amountAccess)
	}
	return choose(array, simpleArrayAccess, simpleAccess)
}

func choose(array bool, arrayTemplate, singleTemplate *template.Template) *template.Template {
	if array {
		return arrayTemplate
	}
	return singleTemplate
}

func (m *MessageElement) buildAccess(c *context, tmpl *template.Template) string {
	var buf bytes.Buffer
	tmpl.Execute(&buf, c)
	return buf.String()
}

func (m *MessageElement) optional() string {
	if m.MinOccurs == "0" || m.MinOccurs == "" {
		return ",omitempty"
	}
	return ""
}

func (m *MessageElement) MemberType() string {
	typeID, _ := m.typeID()
	return typeMap[typeID].Name
}

func (m *MessageElement) Analyse() {
	typeID, _ := m.typeID()
	typeMap[typeID].Used()
}
