package cv3go

import "encoding/xml"

//C s
type C struct {
	// XMLName xml.Name `xml:"CV3Data"`
	CV3Data       RequestBody
	Confirms      []Confirm     `xml:"confirm"`
	OrderStatuses []OrderStatus `xml:"orders"`
	Products      Products      `xml:"products"`
}

// Products data sructure for Products
type Products struct {
	XMLName  xml.Name  `xml:"products"`
	Products []Product `xml:"product,omitempty"`
}

//Product is the struct used when unmarshaling inventory items
type Product struct {
	XMLName          xml.Name         `xml:"product"`
	Inactive         string           `xml:"inactive,attr"`
	Sku              string           `xml:"SKU"`
	ProdId           string           `xml:"ProdID,omitempty"`
	Name             string           `xml:"Name,omitempty"`
	UrlName          string           `xml:"URLName,omitempty"`
	InventoryControl InventoryControl `xml:"InventoryControl,omitempty"`
	Retail           Retail           `xml:"Retail,omitempty"`
	SubProducts      SubProducts      `xml:"SubProducts,omitempty"`
	Description      string           `xml:"Description,omitempty"`
	Keywords         string           `xml:"Keywords,omitempty"`
	Meta             Meta             `xml:"Meta,omitempty"`
	Images           Images           `xml:"Images,omitempty"`
	Categories       ProdCategories   `xml:"Categories,omitempty"`
}

//InventoryControl struct for marshalling and unmarshalling cv3's xml node of InventoryControl
type InventoryControl struct {
	InventoryControlExempt   string `xml:"inventory_control_exempt,attr,omitempty"`
	InventoryStatus          string `xml:"Status,omitempty"`
	InventoryOnHand          string `xml:"Inventory,omitempty"`
	OnOrder                  string `xml:"OnOrder,omitempty"`
	InventoryBackorderedDate string `xml:"InventoryBackorderedDate,omitempty"`
	OutOfStock               string `xml:"OutOfStockPoint,omitempty"`
}

//Retail is the struct for marshalling and unmarshalling cv3's retail node
type Retail struct {
	Active string  `xml:"active,attr,omitempty"`
	Price  Pricing `xml:"Price,omitempty"`
}

//WholeSale hold pricing info for a wholesale item
type WholeSale struct {
	Active        string `xml:"active,attr,omitempty"`
	StandardPrice string `xml:"StandardPrice,omitempty"`
	Qty           string `xml:"Qty,omitempty"`
}

//Special hold the pricing information for a special sale
type Special struct {
	Ongoing string `xml:"ongoing,attr,omitempty"`
	Start   string `xml:"Start,omitempty"`
	End     string `xml:"End,omitempty"`
	Text    string `xml:"Text,omitempty"`
}

//Weight hold the information about the weight
type Weight struct {
	ShipWeight    string `xml:"ShipWeight,omitempty"`
	DisplayWeight string `xml:"DisplayWeight,omitempty"`
	DisplayUnit   string `xml:"DisplayUnit,omitempty"`
}

//Shipping hold the products shipping information
type Shipping struct {
	ShipPreference string  `xml:"ShipPreference,omitempty"`
	FixedRate      string  `xml:"FixedRate,omitempty"`
	Package        Package `xml:"Package,omitempty"`
}

//Package hold the shipping package information
type Package struct {
	ShipsInOwnBox string `xml:"ships_in_own_box,attr,omitempty"`
	Length        string `xml:"Length,omitempty"`
	Width         string `xml:"Width,omitempty"`
	Height        string `xml:"Height,omitempty"`
}

//Pricing is the struct for marshalling and unmarshalling cv3's price node
type Pricing struct {
	PriceCategory string `xml:"price_category,attr,omitempty"`
	StandardPrice string `xml:"StandardPrice,omitempty"`
}

//GiftCertificate hold information about gift certificates
type GiftCertificate struct {
	Active        string `xml:"active,attr,omitempty"`
	DaysAvailable string `xml:"DaysAvailable,omitempty"`
	Value         string `xml:"Value,omitempty"`
}

//Subscription hold the product subscription information
type Subscription struct {
	Active string `xml:"active,attr,omitempty"`
	Price  string `xml:"Price,omitempty"`
}

//ElectronicDelivery holds the information about the products electronic delivery
type ElectronicDelivery struct {
	Active   string   `xml:"active,attr,omitempty"`
	Document Document `xml:"Document,omitempty"`
}

//Document holds information dealing with the electronic delivery
type Document struct {
	DaysAvailable string `xml:"DaysAvailable,omitempty"`
	Description   string `xml:"Description,omitempty"`
}

//Attribute holds product attribute information
type Attribute struct {
	Active string   `xml:"active,attr,omitempty"`
	Values []string `xml:"Value,omitempty"`
}

//Meta is the struct for marshalling and unmarshalling the cv3's Meta node
type Meta struct {
	MetaKeywords    string `xml:"Keyword,omitempty"`
	MetaTitle       string `xml:"Title,omitempty"`
	MetaDescription string `xml:"Description,omitempty"`
}

//Images is the struct for marshalling and unmarshalling cv3's Images node
type Images struct {
	Images []Image `xml:"Image,omitempty"`
}

//Image is the struct for marshalling and unmarshalling cv3's Images node
type Image struct {
	ImageSetThumb1 string `xml:"Thumbnail,omitempty"`
}

//ProdCategories is the struct for marshalling and unmarshalling cv3's Categories node
type ProdCategories struct {
	IDs []string `xml:"ID,omitempty"`
}

//SubProducts is an array of type SubProduct
type SubProducts struct {
	SubProducts []SubProduct `xml:"SubProduct,omitempty"`
	Active      string       `xml:"active,attr"`
}

//SubProduct TODO good description
type SubProduct struct {
	XMLName              xml.Name           `xml:"SubProduct"`
	Inactive             string             `xml:"inactive,attr,omitempty"`
	OutOfSeason          string             `xml:"out_of_season,attr,omitempty"`
	TaxExempt            string             `xml:"tax_exempt,attr,omitempty"`
	GoogleCheckoutExempt string             `xml:"google_checkout_exempt,attr,omitempty"`
	Sku                  string             `xml:"SKU,omitempty"`
	AltID                string             `xml:"AltID,omitempty"`
	ProdId               string             `xml:"ProdID,omitempty"`
	Name                 string             `xml:"Name,omitempty"`
	Image                string             `xml:"Image,omitempty"`
	Retail               Retail             `xml:"Retail,omitempty"`
	WholeSale            WholeSale          `xml:"WholeSale,omitempty"`
	Special              Special            `xml:"Special,omitempty"`
	Weight               Weight             `xml:"Weight,omitempty"`
	Shipping             Shipping           `xml:"Shipping,omitempty"`
	GiftCertificate      GiftCertificate    `xml:"GiftCertificate,omitempty"`
	Subscription         Subscription       `xml:"Subscription,omitempty"`
	ElectronicDelivery   ElectronicDelivery `xml:"ElectronicDelivery,omitempty"`
	Attribute            Attribute          `xml:"Attribute,omitempty"`
	InventoryControl     InventoryControl   `xml:"InventoryControl,omitempty"`
}

//DependancyProducts holds the sku of the product that is depended upon
type DependancyProducts struct {
	Type string   `xml:"type,attr,omitempty"`
	SKUs []string `xml:"SKU,omitempty"`
}

//ProductIDs struct to hold product IDs
type ProductIDs struct {
	ID []string `xml:"productIDs>ID"`
}

// type ProductId struct {
// 	Id string `xml:"ID"`
// }
