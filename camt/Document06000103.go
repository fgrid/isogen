package camt

import (
	"encoding/xml"
)

type Document06000103 struct {
	XMLName                 xml.Name                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Document"`
	AccountReportingRequest *AccountReportingRequestV03 `xml:"AcctRptgReq"`
}

func (d *Document06000103) AddAccountReportingRequest() *AccountReportingRequestV03 {
	d.AccountReportingRequest = new(AccountReportingRequestV03)
	return d.AccountReportingRequest
}
