/*
Blake Ellis <blake@commercev3.com>

Usage:
      api := cv3go.NewApi()
      api.SetCredentials("user-name","password","api-service-id")
      api.GetProductSingle("43523")
      data := api.Execute()
      fmt.Printf(string(data))

*/

//Package cv3go is used to connect to the CV3 API
package cv3go

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"unicode/utf8"
)

const (
	cv3_endpoint = "https://service.commercev3.com/"
	soapEnvelope = "<SOAP-ENV:Envelope xmlns:SOAP-ENV=\"http://www.w3.org/2001/12/soap-envelope\" SOAP-ENV:encodingStyle=\"http://www.w3.org/2001/12/soap-encoding\">\n  <SOAP-ENV:Body>\n<m:CV3Data xmlns:m=\"http://soapinterop.org/\" SOAP-ENV:encodingStyle=\"http://schemas.xmlsoap.org/soap/encoding/\">\n<data xsi:type=\"xsd:string\">%v</data>\n</m:CV3Data>\n</SOAP-ENV:Body>\n</SOAP-ENV:Envelope>\n\n"
)

//Credentials struct
type Credentials struct {
	XMLName   xml.Name `xml:"authenticate"`
	User      string   `xml:"user"`
	Password  string   `xml:"pass"`
	ServiceID string   `xml:"serviceID"`
}

//RequestBody struct
type RequestBody struct {
	XMLName  xml.Name `xml:"request"`
	Auth     Credentials
	Requests []Request `xml:"requests"`
}

//Request struct
type Request struct {
	Request string `xml:",innerxml"`
}

//Confirm struct
type Confirm struct {
	Confirm string `xml:",innerxml"`
}

//OrderStatus struct
type OrderStatus struct {
	OrderStatus string `xml:",innerxml"`
}

//ProductCall struct
type ProductCall struct {
	ProductCall string `xml:",innerxml"`
}

//CV3Data struct
type CV3Data struct {
	// XMLName xml.Name `xml:"CV3Data"`
	CV3Data       RequestBody
	Confirms      []Confirm     `xml:"confirm"`
	OrderStatuses []OrderStatus `xml:"orders"`
	Products      []ProductCall `xml:"products"`
}

type response struct {
	XMLName xml.Name `xml:"Envelope"`
	Data    string   `xml:"Body>CV3DataResponse>return"`
}

type nopCloser struct {
	io.Reader
}

//Convert a string to Base64 encoded string
func toBase64(data string) string {
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	encoder.Write([]byte(data))
	encoder.Close()
	return buf.String()
}

//Api is the struct to send api calls
type Api struct {
	Debug       bool
	user        string
	pass        string
	serviceID   string
	request     string
	confirm     string
	product     string
	prodIgnore  bool
	orderStatus string
}

//NewApi Generate a new API
func NewApi() *Api {
	api := new(Api)
	return api
}

//SetCredentials Set the credentials of the API
func (self *Api) SetCredentials(username, password, serviceID string) {
	self.user = username
	self.pass = password
	self.serviceID = serviceID
	self.prodIgnore = false
}

//GetCustomerGroups Set the request to reqCustomerInformation
func (self *Api) GetCustomerGroups() {
	self.request = "<reqCustomerInformation members_only=\"false\"/>"
}

//GetProductSingle Set the request to reqProducts->reqProductSingle
//containing string(o) as the data
func (self *Api) GetProductSingle(o string) {
	self.request = "<reqProducts><reqProductSingle>" + o + "</reqProductSingle></reqProducts>"
}

//GetProductSKU Set the request to reqProducts->reqProductSKU
//containing string(o) as the data
func (self *Api) GetProductSKU(o string, t bool) {
	if t {
		self.request = "<reqProducts export_sku_only=\"true\"><reqProductSKU>" + o + "</reqProductSKU></reqProducts>"
	} else {
		self.request = "<reqProducts export_sku_only=\"false\"><reqProductSKU>" + o + "</reqProductSKU></reqProducts>"
	}
}

//GetProductSKUs gets the product skus
func (self *Api) GetProductSKUs(o []string, t bool) {
	req := "<reqProducts export_sku_only=\""
	if t {
		req += "true"
	} else {
		req += "false"
	}
	req += "\">"
	for i := 0; i < len(o); i++ {
		req += "<reqProductSKU>" + o[i] + "</reqProductSKU>"
	}
	req += "</reqProducts>"
	self.request = req
}

//GetProductRange Set the request to reqProducts->reqProductRange
//using start and end to dictate the range
func (self *Api) GetProductRange(start string, end string) {
	self.request = "<reqProducts><reqProductRange start=\"" + start + "\" end =\"" + end + "\" /></reqProducts>"
}

//GetProductIds Set the request to reqProductIDs
func (self *Api) GetProductIds() ProductIDs {
	self.request = "<reqProductIDs />"
	data := self.Execute()
	p := ProductIDs{}
	err := xml.Unmarshal(data, &p)
	if err != nil {
		fmt.Printf("can't get products: %v", err)
	}
	return p
}

//GetProductSkus Set the request to reqProductSKU
func (self *Api) GetProductSkus() {
	self.request = "<reqProductSKU />"
}

//GetCatalogRequestsNew Set the request to reqCatalogRequests->reqNew
func (self *Api) GetCatalogRequestsNew() CatalogRequests {
	self.request = "<reqCatalogRequests><reqNew/></reqCatalogRequests>"
	catalogs := CatalogRequests{}
	n := self.Execute()
	err := xml.Unmarshal(n, &catalogs)
	if err != nil {
		fmt.Printf("can't get catalog requests: %v", err)
	}
	return catalogs
}

//GetOrdersNew Set the request to reqOrders->reqOrderNew
func (self *Api) GetOrdersNew() {
	self.request = "<reqOrders><reqOrderNew/></reqOrders>"
}

//GetOrdersRange Set the request to reqOrders->reqOrderOutOfStockPointRange from o to p
func (self *Api) GetOrdersRange(o string, p string) {
	self.request = "<reqOrders><reqOrderOutOfStockPointRange start=\"" + o + "\" end=\"" + p + "\" /></reqOrders>"
}

//OrderConfirm Set request to orderConfirm->orderConf
//using string o as contents
func (self *Api) OrderConfirm(o string) {
	self.confirm = "  <orderConfirm><orderConf>" + o + "</orderConf></orderConfirm>"
}

//UpdateOrderStatus Set request to status->[orderID(o),status(p),tracking(q)]
func (self *Api) UpdateOrderStatus(o string, p string, q string) {
	self.orderStatus = "  <status><orderID>" + o + "</orderID><status>" + p + "</status><tracking>" + q + "</tracking></status>"
}

//CatalogRequest Set request to catalogRequestConfirm->CatalogRequestID(o)
func (self *Api) CatalogRequestConfirm(o string) {
	self.confirm = "  <catalogRequestConfirm><CatalogRequestID>" + o + "</CatalogRequestID></catalogRequestConfirm>"
}

//PushInventory Set the request to an inventory update call
//using o as the data
func (self *Api) PushInventory(o string, t bool) {
	self.product = o
	fmt.Printf("Should we ignore inventory? %+v\n", t)
	if t {
		self.prodIgnore = true
	}
}

//UnmarshalOrders Convert an XML response containing order to an Orders object
func (self *Api) UnmarshalOrders(n []byte) Orders {
	orders := Orders{}
	err := xml.Unmarshal(n, &orders)
	if err != nil {
		fmt.Printf("can't get orders: %v", err)
	}
	return orders
}

//UnmarshalInventory Convert an XML response containing Inventory to a Products object
func (self *Api) UnmarshalInventory(n []byte) Products {
	n = CheckUTF8(n)
	products := Products{}
	err := xml.Unmarshal(n, &products)
	if err != nil {
		fmt.Printf("can't get products: %v", err)
	}
	return products
}

//UnmarshalProduct Convert an XML response containing a single product to a Product object
func (self *Api) UnmarshalProduct(n []byte) Product {
	product := Product{}
	err := xml.Unmarshal(n, &product)
	if err != nil {
		fmt.Printf("can't get product: %v", err)
	}
	return product
}

//Execute Sends the request, return the response
//Note, one of the above requests must
//be set up first, and the credentials must be
//set up for this to work
func (self *Api) Execute() (n []byte) {
	//  var pre_n []byte
	w := Credentials{User: self.user, Password: self.pass, ServiceID: self.serviceID}
	x := Request{Request: self.request}
	y := Confirm{Confirm: self.confirm}
	o := OrderStatus{OrderStatus: self.orderStatus}
	z := ProductCall{ProductCall: self.product}
	t := RequestBody{Auth: w, Requests: []Request{x}}
	v := CV3Data{CV3Data: t, Products: []ProductCall{z}, Confirms: []Confirm{y}, OrderStatuses: []OrderStatus{o}}
	xmlbytes, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Println(err)
	}
	xmlstring := string(xmlbytes)
	xmlstring = strings.Replace(xmlstring, "<CV3Data>", "<CV3Data version=\"2.0\">", -1)
	if self.prodIgnore {
		xmlstring = strings.Replace(xmlstring, "<products>", `<products ignore_new_products="true">`, -1)
	}
	if self.Debug == true {
		fmt.Printf("Printing request string: ")
		fmt.Printf(xmlstring)
	}
	encodedString := toBase64(xmlstring)
	xmlstring = xml.Header + fmt.Sprintf(soapEnvelope, encodedString)
	if err == nil {

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		//client := &http.Client{}
		body := nopCloser{bytes.NewBufferString(xmlstring)}
		if err == nil {
			req, err := http.NewRequest("POST", cv3_endpoint, body)
			if err == nil {
				req.Header.Add("Accept", "text/xml")
				req.Header.Add("Content-Type", "text/xml; charset=utf-8")
				req.Header.Add("SOAPAction", "http://service.commercev3.com/index.php/CV3Data")
				req.ContentLength = int64(len(string(xmlstring)))
				//preq, _ := ioutil.ReadAll(req.Body)
				resp, err := client.Do(req)
				if err != nil {
					fmt.Printf("Request error: %v", err)
					os.Exit(1)
					return
				}
				res, err := ioutil.ReadAll(resp.Body)
				resp.Body.Close()
				if err != nil {
					fmt.Printf("Read Response Error: %v", err)
					os.Exit(1)
					return
				}
				y := response{}
				err = xml.Unmarshal([]byte(res), &y)
				if err != nil {
					fmt.Println(res)
					fmt.Printf("Unmarshal error: %v", err)
					os.Exit(1)
					return
				}
				n, err = base64.StdEncoding.DecodeString(y.Data)
				if err != nil {
					fmt.Printf("Decoding error: %v", err)
					os.Exit(1)
					return
				}
			}
		}
	}
	if self.Debug == true {
		fmt.Printf(string(n))
	}
	if strings.Contains(string(n), "<error>") {
		start := strings.Index(string(n), "<error>") + 7
		end := strings.Index(string(n), "</error>")
		fmt.Println("\nAn error occured: " + string(n[start:end]))
	}
	return
}

//CheckUTF8 converts []byte to []rune to string to []byte to make sure only utf8 characters are used.
func CheckUTF8(b []byte) []byte {
	//if b does not cantain valid utf8
	if !utf8.Valid(b) {
		//convert b into []rune then string then back into []byte
		return []byte(string(bytes.Runes(b)))
	} // else b is valid utf8
	return b
}

//GetTopLevelCategories uses reqCategoryRange with `top_level_only="true"` and  no end set
func (self *Api) GetTopLevelCategories() {
	self.request =
		`<reqCategories top_level_only="true">
	    <reqCategoryRange   start="0"/>
	    </reqCategories>`
}

//GetAllCategories uses reqCategoryRange with no end set
func (self *Api) GetAllCategories() {
	self.request =
		`<reqCategories >
	    <reqCategoryRange  start="0"/>
	    </reqCategories>`
}

/*
<reqCategories>
 <reqCategorySingle>1</reqCategorySingle>
 </reqCategories
*/

//UnmarshalCategories
func (self *Api) UnmarshalCategories(n []byte) Categories {
	//n = CheckUTF8(n)
	categories := Categories{}
	err := xml.Unmarshal(n, &categories)
	if err != nil {
		fmt.Printf("can't get categories: %v", err)
	}
	return categories
}
