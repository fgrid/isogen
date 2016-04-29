package main

import "fmt"

type MessageDefinitionIdentifier struct {
	BusinessArea         string `xml:"businessArea,attr"`
	MessageFunctionality string `xml:"messageFunctionality,attr"`
	Flavour              string `xml:"flavour,attr"`
	Version              string `xml:"version,attr"`
}

func (mdi MessageDefinitionIdentifier) String() string {
	return fmt.Sprintf("%s.%s.%s.%s",
		mdi.BusinessArea,
		mdi.MessageFunctionality,
		mdi.Flavour,
		mdi.Version,
	)
}
