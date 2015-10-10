package ebay

import "encoding/xml"

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
	Xmlns         string `xml:"xmlns,attr"`
	Timestamp     string `xml:"Timestamp"`
	Ack           string `xml:"Ack"`
	Build         string `xml:"Build"`
	Errors        Errors `xml:"Errors"`
	HasMoreOrders bool   `xml:"HasMoreOrders"`

	CorrelationID         string `xml:"CorrelationID"`
	HardExpirationWarning string `xml:"HardExpirationWarning"`

	Paginations  Pagination `xml:"PaginationResult"`
	OrdersArray OrderArray `xml:"OrderArray"`

	OrdersPerPage            int `xml:"OrdersPerPage"`
	PageNumber               int `xml:"PageNumber"`
	ReturnedOrderCountActual int `xml:"ReturnedOrderCountActual"`
}

type Amount struct {
}
type BuyerPackageEnclosures struct {
}
type PaymentInstructionCode struct {
}
type TaxIdentifierAttributeCode struct {
}
type ValueTypeCode struct {
}
type OrderIDArray struct {
	OrderID     string
	BuyerUserID string
}

type OrderArray struct {
	XMLName xml.Name `xml:"OrderArray"`
	Orders  []Order  `xml:"Order"`
}

type Order struct {
	XMLName                             xml.Name `xml:"Order"`
	OrderID                             string
	OrderStatus                         string
	BuyerCheckoutMessage                string
	AmountSaved                         string
	BuyerPackageEnclosures              BuyerPackageEnclosure `xml:"BuyerPackageEnclosure"`
	BuyerTaxIdentifiers                 BuyerTaxIdentifier `xml:"BuyerTaxIdentifier"`
	BuyerUserID                         string
	CancelDetails                       CancelDetail `xml:"CancelDetail"`
	CancelReason                        string
	CancelReasonDetails                 string
	CancelStatus                        string
	CheckoutStatuss                     CheckoutStatus `xml:"CheckoutStatus"`
	ContainseBayPlusTransaction         bool
	CreatedTime                         string
	CreatingUserRole                    string
	EIASToken                           string
	ExtendedOrderID                     string
	ExternalTransactions                ExternalTransaction `xml:"ExternalTransaction"`
	IntegratedMerchantCreditCardEnabled string
	IsMultiLegShipping                  string
	LogisticsPlanType                   string
	MonetaryDetailss                    MonetaryDetails `xml:"MonetaryDetails"`
	MultiLegShippingDetailss            MultiLegShippingDetails `xml:"MultiLegShippingDetails"`
	PaidTime                            string
	PaymentHoldDetailss                 PaymentHoldDetails `xml:"PaymentHoldDetails"`
	PaymentHoldStatus                   string
	PaymentMethods                      string
	PickupDetailss                      PickupDetails `xml:"PickupDetails"`
	PickupMethodSelected
	RefundArrays    RefundArray `xml:"RefundArray"`
	SellerEIASToken string
	SellerEmail     string
	SellerUserID    string
	ShippedTime     string
	//	ShippingAddress
	ShippingConvenienceCharge string
	ShippingDetailss          ShippingDetails `xml:"ShippingDetails"`
	ShippingServiceSelected   string
	Subtotal                  string
	Total                     string
	TransactionArrays         TransactionArray `xml:"TransactionArray"`
}
type Transaction struct {

}
type TransactionArray struct {
	XMLName                             xml.Name `xml:"TransactionArray"`
	Transactions              []Transaction
}

type PickupMethodSelected struct {
	XMLName                             xml.Name `xml:"PickupMethodSelected"`
	MerchantPickupCode    string
	PickupFulfillmentTime string
	PickupLocationUUID    string
	PickupMethod          string
	PickupStatus          string
	PickupStoreID         string
}
type ShippingDetails struct {
	CalculatedShippingRate
	CODCost         string
	InsuranceFee    string
	InsuranceOption string
	InsuranceWanted string

	InternationalShippingServiceOption
	//SellingManagerSalesRecordNumber
	ShipmentTrackingDetails
	ShippingServiceOptions
	TaxTable
}
type TaxJurisdiction struct {
	JurisdictionID        string
	SalesTaxPercent       string
	ShippingIncludedInTax string
}
type TaxTable struct {
	TaxJurisdiction
}
type ShippingPackageInfo struct {
	ActualDeliveryTime       string
	ScheduledDeliveryTimeMax string
	ScheduledDeliveryTimeMin string
	ShippingTrackingEvent    string
	StoreID                  string
}
type ShippingServiceOptions struct {
	ExpeditedService              string
	ImportCharge                  string
	LogisticPlanType              string
	ShippingInsuranceCost         string
	ShippingService               string
	ShippingServiceAdditionalCost string
	ShippingServiceCost           string
	ShippingServicePriority       int
}
type ShipmentTrackingDetails struct {
	ShipmentTrackingNumber string
	ShippingCarrierUsed    string
}
type InternationalShippingServiceOption struct {
	//	ImportCharge  string
	//     ShippingInsuranceCost string
	//     ShippingService string
	//      ShippingServiceAdditionalCost  string
	//     ShippingServiceCost string
	//     ShippingServicePriority int
	//    ShipToLocation string
}
type CalculatedShippingRate struct {
	InternationalPackagingHandlingCosts string
	OriginatingPostalCode               string
	PackageDepth                        string
	PackageLength                       string
	PackageWidth                        string
	PackagingHandlingCosts              string
	ShippingIrregular                   string
	ShippingPackage                     string
	WeightMajor                         string
	WeightMinor                         string
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
	RequiredSellerActionArrays  RequiredSellerActionArray
}
type ShipToAddress struct {
	XMLName                             xml.Name `xml:"ShipToAddress"`
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
	//ReferenceID       string
	StateOrProvince string
	Street1         string
	Street2         string
}
type SellerShipmentToLogisticsProvider struct {
	ShippingServiceDetails
	ShipToAddress
	ShippingTimeMax int
	ShippingTimeMin int
}
type ShippingServiceDetails struct {
	ShippingService   string
	TotalShippingCost string
}
type MultiLegShippingDetails struct {
	SellerShipmentToLogisticsProvider
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
	//ReferenceID       string
	RefundAmount string
	RefundStatus string
	RefundTime   string
	RefundTo     string
	RefundType   string

	//RefundAmount string
	RefundFromSeller string
	//	RefundID string
	//	RefundStatus string
	//	RefundTime string
	TotalRefundToBuyer string
}

type SalesTax struct {
	SalesTaxAmount        string
	SalesTaxPercent       string
	SalesTaxState         string
	ShippingIncludedInTax string
}

type Payment struct {
	FeeOrCreditAmount  string
	Payee              string
	Payer              string
	PaymentAmount      string
	PaymentReferenceID string
	PaymentStatus      string
	PaymentTime        string
	//ReferenceID        string
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

type CheckoutStatus struct {
	eBayPaymentStatus                   string
	IntegratedMerchantCreditCardEnabled string
	LastModifiedTime                    string
	PaymentInstrument                   string
	PaymentMethod                       string
	Status                              string
}

func (o *OrderIDArray) GetOrderId(orderid string) {
	o.OrderID = orderid

}
