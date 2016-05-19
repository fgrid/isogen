package main

import (
	"encoding/xml"
	"flag"
	"log"
	"os"
)

var (
	typeMap                TypeMap
	destinationPackageName string
	messageType            string
)

func init() {
	flag.StringVar(&destinationPackageName, "package", "github.com/fgrid/iso20022", "base package name")
	flag.StringVar(&messageType, "message", "", "message type for generation (empty = all)")
}

func main() {
	flag.Parse()
	typeMap = make(TypeMap)
	var repo Repository
	if err := xml.NewDecoder(os.Stdin).Decode(&repo); err != nil {
		log.Fatalf("could not decode repository from stdin: %s", err.Error())
	}
	repo.Analyse()
	if messageType == "" {
		log.Printf("going to generate all")
		repo.Generate(destinationPackageName)
	} else {
		log.Printf("going to generate %q", messageType)
		repo.GenerateMessage(destinationPackageName, messageType)
	}
}
