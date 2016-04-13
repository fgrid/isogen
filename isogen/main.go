package main

import (
	"encoding/xml"
	"log"
	"os"
)

var typeMap TypeMap

func main() {
	typeMap = make(TypeMap)
	var repo Repository
	if err := xml.NewDecoder(os.Stdin).Decode(&repo); err != nil {
		log.Fatalf("could not decode repository from stdin: %s", err.Error())
	}
	repo.Analyse()
	repo.Generate("iso20022")
}
