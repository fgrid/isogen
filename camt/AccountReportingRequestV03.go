package camt

import "github.com/fgrid/iso20022"

type AccountReportingRequestV03 struct {
	GroupHeader       *iso20022.GroupHeader59        `xml:"GrpHdr"`
	ReportingRequest  []*iso20022.ReportingRequest3  `xml:"RptgReq"`
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData"`
}

func (a *AccountReportingRequestV03) AddGroupHeader() *iso20022.GroupHeader59 {
	a.GroupHeader = new(iso20022.GroupHeader59)
	return a.GroupHeader
}

func (a *AccountReportingRequestV03) AddReportingRequest() *iso20022.ReportingRequest3 {
	newValue := new(iso20022.ReportingRequest3)
	a.ReportingRequest = append(a.ReportingRequest, newValue)
	return newValue
}

func (a *AccountReportingRequestV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
