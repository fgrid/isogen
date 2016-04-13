package main

type DataDictionary struct {
	TopLevelDictionaryEntry []*TopLevelDictionaryEntry `xml:"topLevelDictionaryEntry"`
}

func (d *DataDictionary) Analyse() {
	for _, entry := range d.TopLevelDictionaryEntry {
		entry.Analyse()
	}
}

func (d *DataDictionary) Generate(packageName string) {
	for _, entry := range d.TopLevelDictionaryEntry {
		entry.Generate(packageName)
	}
}
