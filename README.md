# cv3go

cv3go is a go library for working with the [CV3 API](http://www.commercev3.com)

It has been developed to accomplish specific integrations, so certain API functions and fields
may not be included but should be trivial to add since the core connection and parsing is already
being done.

## Install

Run `go get` to download via git.

	go get github.com/ericchiang/pup

## Quick start

Use this library from your Go program like this but with your real credentials:

```go
package main

import (
	"fmt"
	"log"

	"github.com/cv3/cv3go"
)

func main() {
	cv3username, cv3password, cv3apiKey := "**********", "*********", "***********"

	log.Printf("Getting product ids from CV3\r\n")
	api := cv3go.NewApi()
	api.Debug = true
	api.SetCredentials(cv3username, cv3password, cv3apiKey)
	api.GetProductIds()
	data := api.Execute()
	fmt.Printf(string(data))
}
```

## Other Useful Snippits 

```go
	api.Debug = false // turn off all the XML printing and such
```

```go
	// get new orders, or a range or orders
	api.GetOrdersNew()
	// api.GetOrdersRange("20010156", "20010158")
	data := api.Execute()
	orders := api.UnmarshalOrders(data)
	log.Printf("Found %v pending orders\r\n", len(orders.Orders))
	for i := range orders.Orders {
		// ... do stuff with each order ...
	} 
```

```go
	// get catalog requests
	catalogs := api.GetCatalogRequestsNew()
	log.Printf("Found %v pending requests to process\r\n", len(catalogs.CatalogRequests))
	for i := range catalogs.CatalogRequests {
		log.Printf("Catalog Request #%v importing\r\n", catalogs.CatalogRequests[i].CatalogId)
	}
```

