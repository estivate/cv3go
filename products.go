package cv3go

// data sructure for Products
type Products struct {
	Products []Product `xml:"products>product"`
}
type Product struct {
	Inactive                 string       `xml:"inactive,attr"`
	Sku                      string       `xml:"SKU"`
	ProdId                   string       `xml:"ProdID"`
	Name                     string       `xml:"Name"`
	UrlName                  string       `xml:"URLName"`
	InventoryStatus          string       `xml:"InventoryControl>Status"`
	InventoryOnHand          string       `xml:"InventoryControl>Inventory"`
	Price                    string       `xml:"Retail>Price>StandardPrice"`
	OutOfStock               string       `xml:"InventoryControl>OutOfStockPoint"`
	InventoryBackorderedDate string       `xml:"InventoryControl>InventoryBackorderedDate"`
	SubProducts              []SubProduct `xml:"SubProducts>SubProduct"`
}
type SubProduct struct {
	Inactive                 string `xml:"inactive,attr"`
	Sku                      string `xml:"SKU"`
	ProdId                   string `xml:"ProdID"`
	Name                     string `xml:"Name"`
	InventoryStatus          string `xml:"InventoryControl>Status"`
	InventoryOnHand          string `xml:"InventoryControl>Inventory"`
	InventoryBackorderedDate string `xml:"InventoryControl>InventoryBackorderedDate"`
}

type ProductIDs struct {
	ID []string `xml:"productIDs>ID"`
}

// type ProductId struct {
// 	Id string `xml:"ID"`
// }
