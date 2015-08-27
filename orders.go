package cv3go

// data structures for Orders
type ShipToProduct struct {
	SKU      string `xml:"SKU"`
	Quantity string `xml:"quantity"`
	Price    string `xml:"price"`
}
type ShipTo struct {
	Name           string          `xml:"name"`
	FirstName      string          `xml:"firstName"`
	LastName       string          `xml:"lastName"`
	Company        string          `xml:"company"`
	Title          string          `xml:"title"`
	Address        string          `xml:"address1"`
	Address2       string          `xml:"address2"`
	City           string          `xml:"city"`
	State          string          `xml:"state"`
	Zip            string          `xml:"zip"`
	Country        string          `xml:"country"`
	Phone          string          `xml:"phone"`
	Tax            string          `xml:"tax"`
	Shipping       string          `xml:"shipping"`
	Message        string          `xml:"message"`
	ShipMethod     string          `xml:"shipMethod"`
	ShipMethodCode string          `xml:"shipMethodCode"`
	ShipOn         string          `xml:"shipOn"`
	GiftWrap       string          `xml:"giftWrap"`
	ShipToProducts []ShipToProduct `xml:"shipToProducts>shipToProduct"`
}
type Order struct {
	OrderID             string             `xml:"orderID"`
	PriceCategory       string             `xml:"priceCategory"`
	TotalPrice          string             `xml:"totalPrice"`
	TotalShipping       string             `xml:"totalShipping"`
	TotalTax            string             `xml:"totalTax"`
	DateOrdered         string             `xml:"dateOrdered"`
	TimeOrdered         string             `xml:"timeOrdered"`
	PayMethod           string             `xml:"payMethod"`
	SourceCode          string             `xml:"sourceCode"`
	PromoCode           string             `xml:"promoCode"`
	Comments            string             `xml:"comments"`
	IP                  string             `xml:"IP"`
	BillingCompany      string             `xml:"billing>company"`
	BillingFirstName    string             `xml:"billing>firstName"`
	BillingLastName     string             `xml:"billing>lastName"`
	BillingTitle        string             `xml:"billing>title"`
	BillingAddress      string             `xml:"billing>address1"`
	BillingAddress2     string             `xml:"billing>address2"`
	BillingCity         string             `xml:"billing>city"`
	BillingState        string             `xml:"billing>state"`
	BillingZip          string             `xml:"billing>zip"`
	BillingCountry      string             `xml:"billing>country"`
	BillingEmail        string             `xml:"billing>email"`
	BillingPhone        string             `xml:"billing>phone"`
	BillingOptOut       string             `xml:"billing>optOut"`
	CCType              string             `xml:"billing>CCInfo>CCType"`
	CCName              string             `xml:"billing>CCInfo>CCName"`
	CCNum               string             `xml:"billing>CCInfo>CCNum"`
	CCExpM              string             `xml:"billing>CCInfo>CCExpM"`
	CCExpY              string             `xml:"billing>CCInfo>CCExpY"`
	AuthCode            string             `xml:"billing>CCInfo>authCode"`
	AuthAmount          string             `xml:"billing>CCInfo>authAmount"`
	RequestToken        string             `xml:"billing>CCInfo>token"`
	PurchaseOrder       string             `xml:"purchaseOrder"`
	PayPalBuyer         string             `xml:"payPalInfo>Buyer"`
	PayPalAmount        string             `xml:"payPalInfo>Amount"`
	PayPalTransactionID string             `xml:"payPalInfo>TransactionID"`
	ShipTos             []ShipTo           `xml:"shipTos>shipTo"`
	CustomFields        []CustomField      `xml:"customFields>customField"`
	TotalOrderDiscount  TotalOrderDiscount `xml:"totalOrderDiscount"`
}

type TotalOrderDiscount struct {
	Amount        string `xml:"amount"`
	TotalDiscount string `xml:"totalDiscount"`
	Type          string `xml:"type,attr"`
}

type Price struct {
}

type Orders struct {
	Orders []Order `xml:"orders>order"`
}

type CustomField struct {
	Value string `xml:",innerxml"`
}
