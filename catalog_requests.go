package cv3go

import "encoding/xml"

type CatalogRequests struct {
	CatalogRequests []CatalogRequest `xml:"catalogRequests>CatalogRequest"`
}

type CatalogRequest struct {
	CatalogId     string `xml:"id,attr"`
	DateRequested string `xml:"DateRequested"`
	Source        string `xml:"Source"`
	Type          string `xml:"Type"`
	FirstName     string `xml:"CustomerInformation>FirstName"`
	LastName      string `xml:"CustomerInformation>LastName"`
	Company       string `xml:"CustomerInformation>Company"`
	Email         string `xml:"CustomerInformation>Email"`
	Phone         string `xml:"CustomerInformation>Phone"`
	Address       string `xml:"CustomerInformation>Address>Address1"`
	Address2      string `xml:"CustomerInformation>Address>Address2"`
	City          string `xml:"CustomerInformation>Address>City"`
	State         string `xml:"CustomerInformation>Address>State"`
	PostalCode    string `xml:"CustomerInformation>Address>PostalCode"`
	Country       string `xml:"CustomerInformation>Address>Country"`
}

//ConfirmCatalogRequest is the struct to send the confirm catalog database
type ConfirmCatalogRequest struct {
	XMLName          xml.Name `xml:"catalogRequestConfirm"`
	CatalogRequestID []string `xml:"CatalogRequestID"`
}
