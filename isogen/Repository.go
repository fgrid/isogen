package main

import (
	"encoding/xml"
	"log"
)

type Repository struct {
	XMLName                  xml.Name                  `xml:"urn:iso:std:iso:20022:2013:ecore Repository"`
	DataDictionary           *DataDictionary           `xml:"dataDictionary"`
	BusinessProcessCatalogue *BusinessProcessCatalogue `xml:"businessProcessCatalogue"`
}

func (r *Repository) Analyse() {
	r.DataDictionary.Analyse()
}

func (r *Repository) Generate(packageName string) {
	log.Printf("generate repository in package %s", packageName)
	r.DataDictionary.Generate(packageName)
	r.BusinessProcessCatalogue.Generate(packageName)
}
