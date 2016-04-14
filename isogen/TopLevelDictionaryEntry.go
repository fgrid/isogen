package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type TopLevelDictionaryEntry struct {
	XSIType        string            `xml:"xsitype,attr"`
	XMIId          string            `xml:"http://www.omg.org/XMI id,attr"`
	Name           string            `xml:"name,attr"`
	Definition     string            `xml:"definition,attr"`
	Element        []*Element        `xml:"element"`
	MessageElement []*MessageElement `xml:"messageElement"`
	PackageName    string
}

const (
	simpleTypeTemplate = `package {{.PackageName}}

type {{.Name}} string
`
	complexTypeTemplate = `package {{.PackageName}}

// {{replace .Definition "\n" "\n\t// " -1}}
type {{.Name}} struct {
{{range .MessageElement}}
	{{.Declaration}}
{{end}}
}

{{range .MessageElement}}{{.Access $.Name}}{{end}}
`
)

var (
	simpleType  *template.Template
	complexType *template.Template
)

func init() {
	funcMap := template.FuncMap{
		"replace": strings.Replace,
	}
	var err error
	if simpleType, err = template.New("simpleType").Funcs(funcMap).Parse(simpleTypeTemplate); err != nil {
		log.Fatalf("could not compile simpleTypeTemplate - %s", err.Error())
	}
	if complexType, err = template.New("complexType").Funcs(funcMap).Parse(complexTypeTemplate); err != nil {
		log.Fatalf("could not compile complexTypeTemplate - %s", err.Error())
	}
}

func (t *TopLevelDictionaryEntry) Analyse() {
	typeMap[t.XMIId] = t
}

func (t *TopLevelDictionaryEntry) Generate(packageName string) {
	t.PackageName = packageName
	switch t.XSIType {
	case "iso20022:Amount", "iso20022:Binary", "iso20022:CodeSet",
		"iso20022:Date", "iso20022:DateTime", "iso20022:ExternalSchema",
		"iso20022:IdentifierSet", "iso20022:Indicator", "iso20022:Quantity",
		"iso20022:Rate", "iso20022:Text", "iso20022:Time",
		"iso20022:UserDefined", "iso20022:Year", "iso20022:YearMonth":
		t.GenerateSimpleType()
	case "iso20022:MessageComponent", "iso20022:ChoiceComponent":
		t.GenerateComplexType(typeMap)
	}
}

func (t *TopLevelDictionaryEntry) GenerateSimpleType() {
	f, err := os.OpenFile(t.Name+".go", os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		log.Fatalf("can not create file %s - %s", t.Name+".go", err.Error())
	}
	defer f.Close()
	simpleType.Execute(f, t)
}

func (t *TopLevelDictionaryEntry) GenerateComplexType(typeMap TypeMap) {
	f, err := os.OpenFile(t.Name+".go", os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		log.Fatalf("can not create file %s - %s", t.Name+".go", err.Error())
	}
	defer f.Close()
	complexType.Execute(f, t)
}
