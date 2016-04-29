package main

import "log"

type BusinessProcessCatalogue struct {
	TopLevelCatalogueEntry []TopLevelCatalogueEntry `xml:"topLevelCatalogueEntry"`
}

func (b *BusinessProcessCatalogue) Generate(packageName string) {
	log.Printf("generate business process catalogue in package %s", packageName)
	for _, entry := range b.TopLevelCatalogueEntry {
		entry.Generate(packageName)
	}
}
