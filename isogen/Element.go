package main

type Element struct {
	XSIType string `xml:"xsitype,attr"`
	XMIId   string `xml:"http://www.omg.org/XMI id,attr"`
	Name    string `xml:"name"`
}
