package main

import "encoding/xml"

type Repository struct {
	XMLName                  xml.Name                  `xml:urn:iso:std:iso:20022:2013:ecore Repository"`
	DataDictionary           *DataDictionary           `xml:"dataDictionary"`
	BusinessProcessCatalogue *BusinessProcessCatalogue `xml:"businessProcessCatalogue"`
}

func (r *Repository) Analyse() {
	r.DataDictionary.Analyse()
}

func (r *Repository) Generate(packageName string) {
	r.DataDictionary.Generate(packageName)
}
