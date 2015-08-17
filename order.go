package ebay

type GetOrdersRequest struct {
	Xmlns string `xml:"xmlns,attr"`

	RequesterCredentials RequesterCredentials `xml:"RequesterCredentials"`

	CreateTimeFrom       string
	CreateTimeTo         string
	IncludeFinalValueFee string
	//  ListingType          string
	//  ModTimeFrom string
	//	ModTimeTo            string
	//	NumberOfDays         string
	OrderRole   string
	OrderStatus string
	//	SortingOrder         string
	DetailLevel string
	//	ErrorLanguage        string
	//	MessageID            string
	//	OutputSelector       string
	Version string
	//	WarningLevel         string
	//	Pagination
	//	OrderIDArray
}

type GetOrdersRequestResponse struct {
	Xmlns                    string `xml:"xmlns,attr"`
	Timestamp                string
	Ack                      string
	Build                    string
	Errors                   Errors
	HasMoreOrders            bool
	ReturnedOrderCountActual int
	CorrelationID            string
	HardExpirationWarning    string

	Pagination
	OrderArray OrderArray
}

type Pagination struct {
	EntriesPerPage       string
	OrdersPerPage        string
	PageNumber           string
	TotalNumberOfEntries int
	TotalNumberOfPages   int
}

type OrderIDArray struct {
	OrderID     string
	BuyerUserID string
}

type OrderArray struct {
	Order Order
}

type Order struct {
	BuyerCheckoutMessage string
	AmountSaved          string
	BuyerPackageEnclosure
	BuyerTaxIdentifier
	BuyerUserID string
	CancelDetail
	CancelReasonDetails string
	CancelStatus        string
	CheckoutStatus
}

type BuyerPackageEnclosure struct {
	BuyerPackageEnclosureType string
}
type BuyerTaxIdentifier struct {
	Attribute string
	ID        string
	Type      string
}

type CancelDetail struct {
	CancelCompleteDate  string
	CancelIntiationDate string
	CancelIntiator      string
	CancelReason        string
	CancelReasonDetails string
}

type CheckoutStatus struct{}

func (p *Pagination) Setpagination(entriesperpage, pagenumber string) {
	p.EntriesPerPage = entriesperpage
	p.PageNumber = pagenumber
}

func (o *OrderIDArray) GetOrderId(orderid string) {
	o.OrderID = orderid

}
