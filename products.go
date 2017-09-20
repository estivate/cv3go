package cv3go

// Products data sructure for Products
type Products struct {
	Products []Product `xml:"products>product"`
}

//Product is the struct used when unmarshaling inventory items
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
	Description              string       `xml:"Description,omitempty"`
	Keywords                 string       `xml:"Keywords,omitempty"`
	MetaKeywords             string       `xml:"Meta>Keyword,omitempty"`
	MetaTitle                string       `xml:"Meta>Title,omitempty"`
	MetaDescription          string       `xml:"Meta>Description,omitempty"`
	ImageSetThumb1           string       `xml:"Images>Image>Thumbnail,omitempty"`
	CategoryIDs              []CategoryID `xml:"Categories,omitempty"`
}

/*
type SearchItem struct {
	SKU               string `gorm:"primary_key"`
	ProdName          string
	ProdDescription   string
	ChildImage        string
	RetailPrice1      string
	ProductURLName    string
	Keywords          string
	MetaKeywords      string
	MetaTitle         string
	MetaDescription   string
	ImageSetThumb1    string
	CategoryIDs       string
	DefaultCategoryID string
	SearchKeys        string
	KeyWords          string
}*/

//SubProduct TODO good description
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

//ProductIDs struct to hold product IDs
type ProductIDs struct {
	ID []string `xml:"productIDs>ID"`
}

//CategoryID used in the product struct
type CategoryID struct {
	ID string `xml:"ID"`
}

// type ProductId struct {
// 	Id string `xml:"ID"`
// }
