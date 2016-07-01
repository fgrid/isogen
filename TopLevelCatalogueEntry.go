package main

type TopLevelCatalogueEntry struct {
	XSIType           string               `xml:"xsitype,attr"`
	XMIId             string               `xml:"http://www.omg.org/XMI id,attr"`
	Name              string               `xml:"name,attr"`
	Definition        string               `xml:"definition,attr"`
	Code              string               `xml:"code,attr"`
	MessageDefinition []*MessageDefinition `xml:"messageDefinition"`
	PackageName       string
}

func (t *TopLevelCatalogueEntry) Generate(packageName string) {
	t.PackageName = packageName
	if t.XSIType != "iso20022:BusinessArea" {
		return
	}
	for _, md := range t.MessageDefinition {
		md.Analyse()
		md.Generate(packageName)
	}
}

func (t *TopLevelCatalogueEntry) GenerateMessage(packageName, messageType string) {
	t.PackageName = packageName
	if t.XSIType != "iso20022:BusinessArea" {
		return
	}
	for _, md := range t.MessageDefinition {
		if md.MessageDefinitionIdentifier.String() == messageType {
			md.Analyse()
			md.Generate(packageName)
		}
	}
}
