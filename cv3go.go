/*
Blake Ellis <blake@commercev3.com>

Usage:
      api := cv3go.NewApi()
      api.SetCredentials("user-name","password","api-service-id")
      api.GetProductSingle("43523")
      data := api.Execute()
      fmt.Printf(string(data))

*/

package cv3go

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	cv3_endpoint = "https://service.commercev3.com/"
	soapEnvelope = "<SOAP-ENV:Envelope xmlns:SOAP-ENV=\"http://www.w3.org/2001/12/soap-envelope\" SOAP-ENV:encodingStyle=\"http://www.w3.org/2001/12/soap-encoding\">\n  <SOAP-ENV:Body>\n<m:CV3Data xmlns:m=\"http://soapinterop.org/\" SOAP-ENV:encodingStyle=\"http://schemas.xmlsoap.org/soap/encoding/\">\n<data xsi:type=\"xsd:string\">%v</data>\n</m:CV3Data>\n</SOAP-ENV:Body>\n</SOAP-ENV:Envelope>\n\n"
)

type Credentials struct {
	XMLName   xml.Name `xml:"authenticate"`
	User      string   `xml:"user"`
	Password  string   `xml:"pass"`
	ServiceID string   `xml:"serviceID"`
}

type RequestBody struct {
	XMLName  xml.Name `xml:"request"`
	Auth     Credentials
	Requests []Request `xml:"requests"`
}

type Request struct {
	Request string `xml:",innerxml"`
}

type Confirm struct {
	Confirm string `xml:",innerxml"`
}

type OrderStatus struct {
	OrderStatus string `xml:",innerxml"`
}

type ProductCall struct {
	ProductCall string `xml:",innerxml"`
}

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

func toBase64(data string) string {
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	encoder.Write([]byte(data))
	encoder.Close()
	return buf.String()
}

type Api struct {
	Debug       bool
	user        string
	pass        string
	serviceID   string
	request     string
	confirm     string
	product     string
	orderStatus string
}

func NewApi() *Api {
	api := new(Api)
	return api
}

func (self *Api) SetCredentials(username, password, serviceID string) {
	self.user = username
	self.pass = password
	self.serviceID = serviceID
}

func (self *Api) GetCustomerGroups() {
	self.request = "<reqCustomerInformation members_only=\"false\"/>"
}

func (self *Api) GetProductSingle(o string) {
	self.request = "<reqProducts><reqProductSingle>" + o + "</reqProductSingle></reqProducts>"
}

func (self *Api) GetProductSKU(o string) {
	self.request = "<reqProducts><reqProductSKU>" + o + "</reqProductSKU></reqProducts>"
}

func (self *Api) GetProductRange(start string, end string) {
	self.request = "<reqProducts><reqProductRange start=\"" + start + "\" end =\"" + end + "\" /></reqProducts>"
}

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

func (self *Api) GetProductSkus() {
	self.request = "<reqProductSKU />"
}

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

func (self *Api) GetOrdersNew() {
	self.request = "<reqOrders><reqOrderNew/></reqOrders>"
}

func (self *Api) GetOrdersRange(o string, p string) {
	self.request = "<reqOrders><reqOrderRange start=\"" + o + "\" end=\"" + p + "\" /></reqOrders>"
}

func (self *Api) OrderConfirm(o string) {
	self.confirm = "  <orderConfirm><orderConf>" + o + "</orderConf></orderConfirm>"
}

func (self *Api) UpdateOrderStatus(o string, p string, q string) {
	self.orderStatus = "  <status><orderID>" + o + "</orderID><status>" + p + "</status><tracking>" + q + "</tracking></status>"
}

func (self *Api) CatalogRequestConfirm(o string) {
	self.confirm = "  <catalogRequestConfirm><CatalogRequestID>" + o + "</CatalogRequestID></catalogRequestConfirm>"
}

func (self *Api) PushInventory(o string) {
	self.product = o
}

func (self *Api) UnmarshalOrders(n []byte) Orders {
	orders := Orders{}
	err := xml.Unmarshal(n, &orders)
	if err != nil {
		fmt.Printf("can't get orders: %v", err)
	}
	return orders
}

func (self *Api) UnmarshalInventory(n []byte) Products {
	products := Products{}
	err := xml.Unmarshal(n, &products)
	if err != nil {
		fmt.Printf("can't get products: %v", err)
	}
	return products
}

func (self *Api) UnmarshalProduct(n []byte) Product {
	product := Product{}
	err := xml.Unmarshal(n, &product)
	if err != nil {
		fmt.Printf("can't get product: %v", err)
	}
	return product
}

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
	xmlstring := string(xmlbytes)
	xmlstring = strings.Replace(xmlstring, "<CV3Data>", "<CV3Data version=\"2.0\">", -1)
	if self.Debug == true {
		fmt.Printf(xmlstring)
	}
	encodedString := toBase64(xmlstring)
	xmlstring = xml.Header + fmt.Sprintf(soapEnvelope, encodedString)
	if err == nil {
		client := &http.Client{}
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
					return
				}
				res, err := ioutil.ReadAll(resp.Body)
				resp.Body.Close()
				if err != nil {
					fmt.Printf("Read Response Error: %v", err)
					return
				}
				y := response{}
				err = xml.Unmarshal([]byte(res), &y)
				if err != nil {
					fmt.Printf("Unmarshal error: %v", err)
					return
				}
				n, err = base64.StdEncoding.DecodeString(y.Data)
				if err != nil {
					fmt.Printf("Decoding error: %v", err)
					return
				}
			}
		}
	}
	if self.Debug == true {
		fmt.Printf(string(n))
	}
	return
}
