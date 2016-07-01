package main

type BusinessProcessCatalogue struct {
	TopLevelCatalogueEntry []TopLevelCatalogueEntry `xml:"topLevelCatalogueEntry"`
}

func (b *BusinessProcessCatalogue) Generate(packageName string) {
	for _, entry := range b.TopLevelCatalogueEntry {
		entry.Generate(packageName)
	}
}

func (b *BusinessProcessCatalogue) GenerateMessage(packageName, messageType string) {
	for _, entry := range b.TopLevelCatalogueEntry {
		entry.GenerateMessage(packageName, messageType)
	}

}
