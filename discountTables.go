package cv3go

import "encoding/xml"

//
//ProductGroups structs below
//

//ProductGroups is the struct to hold data for the cv3 webservice productGroup calls
type ProductGroups struct {
	ProductGroup []ProductGroup `xml:"ProductGroup"`
}

//ProductGroup is the struct to hold product group data for the cv3 webservice calls
type ProductGroup struct {
	Delete        string `xml:"delete,attr"`
	Name          string `xml:"Name"`
	Description   string `xml:"Description"`
	DiscountTable string `xml:"DiscountTable"`
	Type          string `xml:"Type"`
	Products      []SKU  `xml:"Products"`
}

//SKU is the struct to hold skus
type SKU struct {
	SKU string `xml:"SKU"`
}

//
//DiscountTables structs below
//

//ReqPricing is the struct to hold various pricing requests from the cv3 web serviceID
type ReqPricing struct {
	XMLName xml.Name    `xml:"reqPricing"`
	Data    interface{} `xml:",innerxml"`
}

//DiscountTables is the struct to hold discountTable data for CV3's webservice
type DiscountTables struct {
	ReplaceExisting string          `xml:"replace_existing,attr"`
	DiscountTable   []DiscountTable `xml:"DiscountTable"`
}

//DiscountTable is the struct to hold discount table data for the cv3 CV3WebService
type DiscountTable struct {
	Delete             string    `xml:"delete,attr"`
	Name               string    `xml:"Name"`
	Description        string    `xml:"Description"`
	Discounts          Discounts `xml:"Discounts"`
	AdditionalDiscount Discount  `xml:"AdditionalDiscount"`
}

//Discounts is the struct to hold discount table data for the cv3 RunWebServiceProductChunkFunction
type Discounts struct {
	Discount []Discount `xml:"Discount"`
}

//Discount is the struct to hold data for the discount tables for the cv3 CV3WebService
type Discount struct {
	Delete string `xml:"delete,attr,omitempty"`
	Qty    string `xml:"Qty"`
	Amount string `xml:"Amount"`
	Type   string `xml:"Type"`
}

//ReqDiscountTables is the struct for requesting discount tables from the cv3 web service
type ReqDiscountTables struct {
	XMLName   xml.Name `xml:"reqDiscountTables"`
	ReturnAll string   `xml:"return_all,attr"`
	ByName    *ByName  `xml:"ByName,omitempty"`
	ByID      *ByID    `xml:"ByID,omitempty"`
}

//ReqProductGroups is the struct to hold data for the product group request to cv3's WebServiceChunkFunc
type ReqProductGroups struct {
	XMLName   xml.Name `xml:"reqProductGroups"`
	ReturnAll string   `xml:"return_all,attr"`
	ByName    *ByName  `xml:"ByName,omitempty"`
	ByID      *ByID    `xml:"ByIDomitempty"`
}

//ByName holds name request data for pricing calls to cv3 WebServiceChunkFunc
type ByName struct {
	Name []string `xml:"Name,omitempty"`
}

//ByID holds ID request data for pricing calls to cv3 WebServiceChunkFunc
type ByID struct {
	ID []string `xml:"ID,omitemtpy"`
}
