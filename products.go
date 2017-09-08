package cv3go

// data sructure for Products
type Products struct {
	Products []Product `xml:"products>product"`
}
type Product struct {
	Inactive                 string       `xml:"inactive,attr"`
	Sku                      string       `xml:"SKU"`
	ProdId                   string       `xml:"ProdID,omitempty"`
	Name                     string       `xml:"Name,omitempty"`
	UrlName                  string       `xml:"URLName,omitempty"`
	InventoryStatus          string       `xml:"InventoryControl>Status,omitempty"`
	InventoryOnHand          string       `xml:"InventoryControl>Inventory,omitempty"`
	OnOrder                  string       `xml:"InventoryControl>OnOrder,omitempty"`
	Price                    string       `xml:"Retail>Price>StandardPrice,omitempty"`
	OutOfStock               string       `xml:"InventoryControl>OutOfStockPoint,omitempty"`
	InventoryBackorderedDate string       `xml:"InventoryControl>InventoryBackorderedDate,omitempty"`
	SubProducts              []SubProduct `xml:"SubProducts>SubProduct,omitempty"`
}
type SubProduct struct {
	Inactive                 string `xml:"inactive,attr"`
	Sku                      string `xml:"SKU"`
	ProdId                   string `xml:"ProdID"`
	Name                     string `xml:"Name"`
	InventoryStatus          string `xml:"InventoryControl>Status"`
	InventoryOnHand          string `xml:"InventoryControl>Inventory"`
	OnOrder                  string `xml:"InventoryControl>OnOrder,omitempty"`
	InventoryBackorderedDate string `xml:"InventoryControl>InventoryBackorderedDate"`
}

type ProductIDs struct {
	ID []string `xml:"productIDs>ID"`
}

// type ProductId struct {
// 	Id string `xml:"ID"`
// }
