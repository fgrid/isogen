package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

type MessageDefinition struct {
	XMIId                       string                      `xml:"http://www.omg.org/XMI id,attr"`
	Definition                  string                      `xml:"definition,attr"`
	Name                        string                      `xml:"name,attr"`
	XMLTag                      string                      `xml:"xmlTag,attr"`
	MessageBuildingBlock        []MessageElement            `xml:"messageBuildingBlock"`
	MessageDefinitionIdentifier MessageDefinitionIdentifier `xml:"messageDefinitionIdentifier"`
	BasePackagePath             string
	BasePackageName             string
}

const _Template = `package {{.MessageDefinitionIdentifier.BusinessArea}}

import (
	"encoding/xml"

	"{{.BasePackagePath}}"
)

type {{.DocumentType}} struct {
	XMLName xml.Name {{.DocumentName}}
	Message *{{.Name}} {{.XMLName}}
}

func (d *{{.DocumentType}}) AddMessage() *{{.Name}} {
	d.Message = new({{.Name}})
	return d.Message
}

// {{replace .Definition "\n" "\n// " -1}}
type {{.Name}} struct {
{{range .MessageBuildingBlock}}
	{{.DeclarationOut $.BasePackageName}}
{{end}}
}

{{range .MessageBuildingBlock}}{{.AccessOut $.BasePackageName $.Name}}{{end}}
`

var _tmpl *template.Template

func init() {
	funcMap := template.FuncMap{
		"replace": strings.Replace,
	}
	var err error
	if _tmpl, err = template.New("complexType").Funcs(funcMap).Parse(_Template); err != nil {
		log.Fatalf("could not compile template - %s", err.Error())
	}
}

func (md *MessageDefinition) Generate(packageName string) {
	if md.Name == "RequestToModifyPaymentV03" {
		log.Printf("skipping amigous message definition %s", md.Name)
		return
	}
	md.BasePackagePath = packageName
	parts := strings.Split(packageName, "/")
	md.BasePackageName = parts[len(parts)-1]
	if err := os.MkdirAll(md.MessageDefinitionIdentifier.BusinessArea, 0770); err != nil {
		log.Printf("could not create subdirectory %s - %s",
			md.MessageDefinitionIdentifier.BusinessArea, err.Error())
		return
	}
	fileName := fmt.Sprintf("%s/%s.go",
		md.MessageDefinitionIdentifier.BusinessArea,
		md.Name,
	)
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		log.Fatalf("can not create file %s - %s", fileName, err.Error())
	}
	defer f.Close()
	_tmpl.Execute(f, md)
}

func (md *MessageDefinition) XMLName() string {
	return "`xml:\"" + md.XMLTag + "\"`"
}

func (md *MessageDefinition) DocumentType() string {
	return "Document" + md.MessageDefinitionIdentifier.MessageFunctionality + md.MessageDefinitionIdentifier.Flavour + md.MessageDefinitionIdentifier.Version
}

func (md *MessageDefinition) DocumentName() string {
	return "`xml:\"urn:iso:std:iso:20022:tech:xsd:" + md.MessageDefinitionIdentifier.String() + " Document\"`"
}
