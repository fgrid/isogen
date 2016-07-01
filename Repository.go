package main

import "encoding/xml"

type Repository struct {
	XMLName                  xml.Name                  `xml:"urn:iso:std:iso:20022:2013:ecore Repository"`
	DataDictionary           *DataDictionary           `xml:"dataDictionary"`
	BusinessProcessCatalogue *BusinessProcessCatalogue `xml:"businessProcessCatalogue"`
}

func (r *Repository) Analyse() {
	r.DataDictionary.Analyse()
}

func (r *Repository) Generate(packageName string) {
	r.BusinessProcessCatalogue.Generate(packageName)
	r.DataDictionary.Generate(packageName)
}

func (r *Repository) GenerateMessage(packageName, messageType string) {
	r.BusinessProcessCatalogue.GenerateMessage(packageName, messageType)
	r.DataDictionary.Generate(packageName)
}
