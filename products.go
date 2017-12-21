package cv3go

// Products data sructure for Products
type Products struct {
	Products []Product `xml:"product,omitempty"`
}

//Product is the struct used when unmarshaling inventory items
type Product struct {
	Inactive         string           `xml:"inactive,attr"`
	Sku              string           `xml:"SKU"`
	ProdId           string           `xml:"ProdID,omitempty"`
	Name             string           `xml:"Name,omitempty"`
	UrlName          string           `xml:"URLName,omitempty"`
	InventoryControl InventoryControl `xml:"InventoryControl,omitempty"`
	Retail           Retail           `xml:"Retail,omitempty"`
	SubProducts      []SubProduct     `xml:"SubProducts,omitempty>SubProduct,omitempty"`
	Description      string           `xml:"Description,omitempty"`
	Keywords         string           `xml:"Keywords,omitempty"`
	Meta             Meta             `xml:"Meta,omitempty"`
	Images           Images           `xml:"Images,omitempty"`
	Categories       ProdCategories   `xml:"Categories,omitempty"`
}

//InventoryControl struct for marshalling and unmarshalling cv3's xml node of InventoryControl
type InventoryControl struct {
	InventoryStatus          string `xml:"Status,omitempty"`
	InventoryOnHand          string `xml:"Inventory,omitempty"`
	OnOrder                  string `xml:"OnOrder,omitempty"`
	InventoryBackorderedDate string `xml:"InventoryBackorderedDate,omitempty"`
	OutOfStock               string `xml:"OutOfStockPoint,omitempty"`
}

//Retail is the struct for marshalling and unmarshalling cv3's retail node
type Retail struct {
	Price Pricing `xml:"Price,omitempty"`
}

//Pricing is the struct for marshalling and unmarshalling cv3's price node
type Pricing struct {
	StandardPrice string `xml:"StandardPrice,omitempty"`
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

//SubProduct TODO good description
type SubProduct struct {
	Inactive                 string `xml:"inactive,attr,omitempty"`
	Sku                      string `xml:"SKU,omitempty"`
	ProdId                   string `xml:"ProdID,omitempty"`
	Name                     string `xml:"Name,omitempty"`
	InventoryStatus          string `xml:"InventoryControl>Status,omitempty"`
	InventoryOnHand          string `xml:"InventoryControl>Inventory,omitempty"`
	OnOrder                  string `xml:"InventoryControl>OnOrder,omitempty"`
	InventoryBackorderedDate string `xml:"InventoryControl>InventoryBackorderedDate,omitempty"`
}

//ProductIDs struct to hold product IDs
type ProductIDs struct {
	ID []string `xml:"productIDs>ID"`
}

// type ProductId struct {
// 	Id string `xml:"ID"`
// }
