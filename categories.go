package cv3go

// Categories data sructure for Categories
type Categories struct {
	Categories []Category `xml:"categories>category"`
}

//Category is the struct used when unmarshaling categories
type Category struct {
	Invisible          string        `xml:"invisible,attr"`
	TopLevel           string        //not sent in xml, set in cv3Intigration.getCategories.go
	Featured           string        `xml:"featured,attr"`
	Name               string        `xml:"Name,omitempty"`
	ID                 string        `xml:"ID,omitempty"`
	URLName            string        `xml:"URLName,omitempty"`
	Description        string        `xml:"Description,omitempty"`
	MetaTitle          string        `xml:"MetaTitle,omitempty"`
	MetaDescription    string        `xml:"MetaDescription,omitempty"`
	MetaKeywords       string        `xml:"MetaKeyword,omitempty"`
	Template           string        `xml:"Template,omitempty"`
	NumProductsPerPage string        `xml:"NumProductsPerPage,omitempty"`
	Products           []string      `xml:"Products>SKU,omitempty"`
	FeaturedProducts   []string      `xml:"FeaturedProducts>SKU,omitempty"`
	CustomFields       []Custom      `xml:"Custom,omitempty"`
	SubCategories      []SubCategory `xml:"SubCategories>SubCategory"`
}

//SubCategory is the struct to hold the Subcategories
type SubCategory struct {
	Name      string `xml:"Name"`
	ID        string `xml:"ID"`
	Invisible string `xml:"invisible,attr"`
}

//Custom is the struct to hold the custom fields
type Custom struct {
	ID     string `xml:"id,attr"`
	Custom string `xml:",chardata"`
}
