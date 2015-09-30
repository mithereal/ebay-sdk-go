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
	//	Pagination  		Pagination
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

type OrderIDArray struct {
	OrderID     string
	BuyerUserID string
}

type OrderArray struct {
	Order []Order
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
	ContainseBayPlusTransaction bool
	CreatedTime                 string
	CreatingUserRole            string
	EIASToken                   string
	ExtendedOrderID             string
	ExternalTransaction
	IntegratedMerchantCreditCardEnabled string
	IsMultiLegShipping                  string
	LogisticsPlanType                   string
	MonetaryDetails
	OrderID     string
	OrderStatus string
	PaidTime    string
	PaymentHoldDetails
	PaymentHoldStatus string
	PaymentMethods    string
	PickupDetails
}
type PickupMethodSelected struct {
	MerchantPickupCode    string
	PickupFulfillmentTime string
	PickupLocationUUID    string
	PickupMethod          string
	PickupStatus          string
	PickupStoreID         string
}
type PickupOptions struct {
	PickupMethod   string
	PickupPriority string
}

type RefundArray struct{}
type PickupDetails struct {
	PickupOptions
}

type RequiredSellerActionArray struct {
	RequiredSellerAction string
}
type PaymentHoldDetails struct {
	ExpectedReleaseDate   string
	NumOfReqSellerActions string
	PaymentHoldReason     string
	RequiredSellerActionArray
}
type ShipToAddress struct {
	AddressAttribute  string
	AddressID         string
	AddressOwner      string
	CityName          string
	Country           string
	CountryName       string
	ExternalAddressID string
	Name              string
	Phone             string
	PostalCode        string
	ReferenceID       string
	StateOrProvince   string
	Street1           string
	Street2           string
}
type SellerShipmentToLogisticsProvider struct {
	ShippingServiceDetails
	ShipToAddress
}
type ShippingServiceDetails struct {
	ShippingService   string
	TotalShippingCost string
}
type MultiLegShippingDetails struct {
	SellerShipmentToLogisticsProvider
	ShippingTimeMax int
	ShippingTimeMin int
	ShipToAddress
}
type MonetaryDetails struct {
	Payments
	Refunds
}

type Payments struct {
	Payment []Payment
}

type Refunds struct {
	Refund
}
type Refund struct {
	FeeOrCreditAmount string
	ReferenceID       string
	RefundAmount      string
	RefundStatus      string
	RefundTime        string
	RefundTo          string
	RefundType        string
}

type Payment struct {
	FeeOrCreditAmount  string
	Payee              string
	Payer              string
	PaymentAmount      string
	PaymentReferenceID string
	PaymentStatus      string
	PaymentTime        string
	ReferenceID        string
}

type ExternalTransaction struct {
	ExternalTransactionID     string
	ExternalTransactionStatus string
	ExternalTransactionTime   string
	FeeOrCreditAmount         string
	PaymentOrRefundAmount     string
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



func (o *OrderIDArray) GetOrderId(orderid string) {
	o.OrderID = orderid

}
