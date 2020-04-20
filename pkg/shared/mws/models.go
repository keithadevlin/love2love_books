package mws

import "encoding/xml"

//now the strcuct to receive the report request id

type ReportRequestRequestReportResponse struct {
	XMLName                          xml.Name                          `xml:"RequestReportResponse,omitempty" json:"RequestReportResponse,omitempty"`
	Attrxmlns                        string                            `xml:"xmlns,attr"  json:",omitempty"`
	ReportRequestRequestReportResult *ReportRequestRequestReportResult `xml:"http://mws.amazonaws.com/doc/2009-01-01/ RequestReportResult,omitempty" json:"RequestReportResult,omitempty"`
	ReportRequestResponseMetadata    *ReportRequestResponseMetadata    `xml:"http://mws.amazonaws.com/doc/2009-01-01/ ResponseMetadata,omitempty" json:"ResponseMetadata,omitempty"`
}

type ReportRequestRequestReportResult struct {
	XMLName                        xml.Name                        `xml:"RequestReportResult,omitempty" json:"RequestReportResult,omitempty"`
	ReportRequestReportRequestInfo *ReportRequestReportRequestInfo `xml:"http://mws.amazonaws.com/doc/2009-01-01/ ReportRequestInfo,omitempty" json:"ReportRequestInfo,omitempty"`
}

type ReportRequestReportRequestInfo struct {
	XMLName                             xml.Name                             `xml:"ReportRequestInfo,omitempty" json:"ReportRequestInfo,omitempty"`
	ReportRequestEndDate                *ReportRequestEndDate                `xml:"http://mws.amazonaws.com/doc/2009-01-01/ EndDate,omitempty" json:"EndDate,omitempty"`
	ReportRequestReportProcessingStatus *ReportRequestReportProcessingStatus `xml:"http://mws.amazonaws.com/doc/2009-01-01/ ReportProcessingStatus,omitempty" json:"ReportProcessingStatus,omitempty"`
	ReportRequestReportRequestId        *ReportRequestReportRequestId        `xml:"http://mws.amazonaws.com/doc/2009-01-01/ ReportRequestId,omitempty" json:"ReportRequestId,omitempty"`
	ReportRequestReportType             *ReportRequestReportType             `xml:"http://mws.amazonaws.com/doc/2009-01-01/ ReportType,omitempty" json:"ReportType,omitempty"`
	ReportRequestScheduled              *ReportRequestScheduled              `xml:"http://mws.amazonaws.com/doc/2009-01-01/ Scheduled,omitempty" json:"Scheduled,omitempty"`
	ReportRequestStartDate              *ReportRequestStartDate              `xml:"http://mws.amazonaws.com/doc/2009-01-01/ StartDate,omitempty" json:"StartDate,omitempty"`
	ReportRequestSubmittedDate          *ReportRequestSubmittedDate          `xml:"http://mws.amazonaws.com/doc/2009-01-01/ SubmittedDate,omitempty" json:"SubmittedDate,omitempty"`
}

type ReportRequestReportType struct {
	XMLName xml.Name `xml:"ReportType,omitempty" json:"ReportType,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type ReportRequestReportProcessingStatus struct {
	XMLName xml.Name `xml:"ReportProcessingStatus,omitempty" json:"ReportProcessingStatus,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type ReportRequestEndDate struct {
	XMLName xml.Name `xml:"EndDate,omitempty" json:"EndDate,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type ReportRequestScheduled struct {
	XMLName xml.Name `xml:"Scheduled,omitempty" json:"Scheduled,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type ReportRequestReportRequestId struct {
	XMLName xml.Name `xml:"ReportRequestId,omitempty" json:"ReportRequestId,omitempty"`
	Id      string   `xml:",chardata" json:",omitempty"`
}

type ReportRequestSubmittedDate struct {
	XMLName xml.Name `xml:"SubmittedDate,omitempty" json:"SubmittedDate,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type ReportRequestStartDate struct {
	XMLName xml.Name `xml:"StartDate,omitempty" json:"StartDate,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type ReportRequestResponseMetadata struct {
	XMLName                xml.Name                `xml:"ResponseMetadata,omitempty" json:"ResponseMetadata,omitempty"`
	ReportRequestRequestId *ReportRequestRequestId `xml:"http://mws.amazonaws.com/doc/2009-01-01/ RequestId,omitempty" json:"RequestId,omitempty"`
}

type ReportRequestRequestId struct {
	XMLName xml.Name `xml:"RequestId,omitempty" json:"RequestId,omitempty"`
	Id      string   `xml:",chardata" json:",omitempty"`
}

// end of report request structs

type OrderReportItem struct {
	OrderId              string
	OrderItemId          string
	PurchaseDate         string
	PaymentsDate         string
	BuyerEmail           string
	BuyerName            string
	BuyerPhoneNumber     string
	Sku                  string
	ProductName          string
	QuantityPurchased    string
	Currency             string
	ItemPrice            string
	ItemTax              string
	ShippingPrice        string
	ShippingTax          string
	ShipServiceLevel     string
	RecipientName        string
	ShipAddress1         string
	ShipAddress2         string
	ShipAddress3         string
	ShipCity             string
	ShipState            string
	ShipPostalCode       string
	ShipCountry          string
	ShipPhoneNumber      string
	DeliveryStartDate    string
	DeliveryEndDate      string
	DeliveryTimeZone     string
	DeliveryInstructions string
	SalesChannel         string
}
