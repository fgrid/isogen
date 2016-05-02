package main

import (
	"encoding/xml"
	"flag"
	"log"
	"os"
)

var typeMap TypeMap
var destinationPackageName string

func init() {
	flag.StringVar(&destinationPackageName, "package", "github.com/fgrid/iso20022", "base package name")
}

func main() {
	flag.Parse()
	typeMap = make(TypeMap)
	var repo Repository
	if err := xml.NewDecoder(os.Stdin).Decode(&repo); err != nil {
		log.Fatalf("could not decode repository from stdin: %s", err.Error())
	}
	repo.Analyse()
	repo.Generate(destinationPackageName)
}
