package cv3go

// OrdStatus holds our data to update orders with new tracking
// info. It is named badly because I can't figure out what the
// existing OrderStatus struct is doing.
type OrdStatus struct {
	OrderID                  string
	Status                   string
	Tracking                 string
	CustomerNumber           string
	SendTrackingNotification bool
}
