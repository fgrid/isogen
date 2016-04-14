package camt_test

import (
	"encoding/xml"
	"testing"

	"github.com/fgrid/iso20022/camt"
)

func TestDocument06000103(t *testing.T) {
	doc := new(camt.Document06000103)
	acctRptgReq := doc.AddAccountReportingRequest()
	grpHdr := acctRptgReq.AddGroupHeader()
	grpHdr.SetMessageIdentification("EXAMPLE camt.060")
	grpHdr.SetCreationDateTime("2011-01-06T17:15:00")
	rptgReq := acctRptgReq.AddReportingRequest()
	rptgReq.SetRequestedMessageNameIdentification("camt.052.001.03")
	rptgReq.AddAccount().AddIdentification().AddOther().SetIdentification("310141014141")
	rptgReq.AddAccountOwner().AddAgent().AddFinancialInstitutionIdentification().SetBICFI("AAAABEBB")
	rptgPrd := rptgReq.AddReportingPeriod()
	fromToDate := rptgPrd.AddFromToDate()
	fromToDate.SetFromDate("2011-01-10")
	fromToTime := rptgPrd.AddFromToTime()
	fromToTime.SetFromTime("07:30:30")

	buf, err := xml.MarshalIndent(doc, "", "  ")
	if err != nil {
		t.Errorf("could not marshal test document - %s", err.Error())
	}
	result := string(buf)

	t.Logf("serialized Document-006.001.03:\n%s\n", result)

}
