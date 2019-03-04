package ebay

import "encoding/xml"
import "github.com/franela/goreq"
import "io/ioutil"
import "github.com/alecthomas/colour"
import "os/user"
import "path"

type GetOrdersRequest struct {
	Xmlns                string `xml:"xmlns,attr"`

	RequesterCredentials Legacy_RequesterCredentials `xml:"RequesterCredentials"`

	CreateTimeFrom       string
	CreateTimeTo         string
	IncludeFinalValueFee string
	ListingType          string
	ModTimeFrom          string
	ModTimeTo            string
	NumberOfDays         string
	OrderRole            string
	OrderStatus          string
	SortingOrder         string
	DetailLevel          string
	ErrorLanguage        string
	MessageID            string
	OutputSelector       string
	Version              string
	WarningLevel         string
	Pagination           Legacy_Pagination
	OrderIDArray Legacy_OrderIDArray
}

type Legacy_GetOrdersRequestResponse struct {
	Xmlns                    string `xml:"xmlns,attr"`
	Timestamp                string `xml:"Timestamp"`
	Ack                      string `xml:"Ack"`
	Build                    string `xml:"Build"`
	Errors                   Errors `xml:"Errors"`
	HasMoreOrders            bool   `xml:"HasMoreOrders"`

	CorrelationID            string `xml:"CorrelationID"`
	HardExpirationWarning    string `xml:"HardExpirationWarning"`

	Paginations              Legacy_Pagination `xml:"PaginationResult"`
	OrdersArray              Legacy_OrderArray `xml:"OrderArray"`

	OrdersPerPage            int `xml:"OrdersPerPage"`
	PageNumber               int `xml:"PageNumber"`
	ReturnedOrderCountActual int `xml:"ReturnedOrderCountActual"`
}

type Legacy_Amount struct {
}
type Legacy_BuyerPackageEnclosures struct {
}
type Legacy_PaymentInstructionCode struct {
}
type Legacy_TaxIdentifierAttributeCode struct {
}
type Legacy_ValueTypeCode struct {
}
type Legacy_OrderIDArray struct {
	OrderID     string
	BuyerUserID string
}

type Legacy_OrderArray struct {
	XMLName xml.Name `xml:"OrderArray"`
	Orders  []Legacy_Order  `xml:"Order"`
}

type Legacy_Order struct {
	XMLName                             xml.Name `xml:"Order"`
	OrderID                             string
	OrderStatus                         string
	BuyerCheckoutMessage                string
	AmountSaved                         string
	BuyerPackageEnclosures              Legacy_BuyerPackageEnclosure `xml:"BuyerPackageEnclosure"`
	BuyerTaxIdentifiers                 Legacy_BuyerTaxIdentifier `xml:"BuyerTaxIdentifier"`
	BuyerUserID                         string
	CancelDetails                       Legacy_CancelDetail `xml:"CancelDetail"`
	CancelReason                        string
	CancelReasonDetails                 string
	CancelStatus                        string
	CheckoutStatuss                     Legacy_CheckoutStatus `xml:"CheckoutStatus"`
	ContainseBayPlusTransaction         bool
	CreatedTime                         string
	ModifiedTime                        string
	CreatingUserRole                    string
	EIASToken                           string
	ExtendedOrderID                     string
	ExternalTransactions                Legacy_ExternalTransaction `xml:"ExternalTransaction"`
	IntegratedMerchantCreditCardEnabled string
	IsMultiLegShipping                  string
	LogisticsPlanType                   string
	MonetaryDetailss                    Legacy_MonetaryDetails `xml:"MonetaryDetails"`
	MultiLegShippingDetailss            Legacy_MultiLegShippingDetails `xml:"MultiLegShippingDetails"`
	PaidTime                            string
	PaymentHoldDetailss                 Legacy_PaymentHoldDetails `xml:"PaymentHoldDetails"`
	PaymentHoldStatus                   string
	PaymentMethods                      string
	PickupDetailss                      Legacy_PickupDetails `xml:"PickupDetails"`
	PickupMethodSelected
	RefundArrays                        Legacy_RefundArray `xml:"RefundArray"`
	SellerEIASToken                     string
	SellerEmail                         string
	SellerUserID                        string
	ShippedTime                         string
	//	ShippingAddress
	ShippingConvenienceCharge           string
	ShippingDetailss                    Legacy_ShippingDetails `xml:"ShippingDetails"`
	ShippingServiceSelected             string
	Subtotal                            string
	Total                               float64
	TransactionArrays                   Legacy_TransactionArray `xml:"TransactionArray"`
}
type Legacy_Transaction struct {
	//Buyer  Buyer
	ShippingDetails     Legacy_ShippingDetails
	CreatedDate         Legacy_string
	Item                Legacy_Item
	QuantityPurchased   string
	//	Status Status
	TransactionID       string
	TransactionPrice    string
	//ShippingServiceSelected ShippingServiceSelected
	FinalValueFee       string
	TransactionSiteID   string
	Platform            string
	//Taxes Taxes
	ActualShippingCost  float32
	ActualHandlingCost  float32
	OrderLineItemID     string
	ExtendedOrderID     string
	eBayPlusTransaction bool

}
type Legacy_TransactionArray struct {
	XMLName      xml.Name `xml:"TransactionArray"`
	Transactions []Legacy_Transaction
}

type Legacy_PickupMethodSelected struct {
	XMLName               xml.Name `xml:"PickupMethodSelected"`
	MerchantPickupCode    string
	PickupFulfillmentTime string
	PickupLocationUUID    string
	PickupMethod          string
	PickupStatus          string
	PickupStoreID         string
}
type Legacy_ShippingDetails struct {
	CalculatedShippingRate Legacy_CalculatedShippingRate
	CODCost         string
	InsuranceFee    string
	InsuranceOption string
	InsuranceWanted string

	InternationalShippingServiceOption  Legacy_InternationalShippingServiceOption
	//SellingManagerSalesRecordNumber
	ShipmentTrackingDetails Legacy_ShipmentTrackingDetails
	ShippingServiceOptions Legacy_ShippingServiceOptions
	TaxTable Legacy_TaxTable
}
type Legacy_TaxJurisdiction struct {
	JurisdictionID        string
	SalesTaxPercent       string
	ShippingIncludedInTax string
}
type Legacy_TaxTable struct {
	TaxJurisdiction Legacy_TaxJurisdiction
}
type Legacy_ShippingPackageInfo struct {
	ActualDeliveryTime       string
	ScheduledDeliveryTimeMax string
	ScheduledDeliveryTimeMin string
	ShippingTrackingEvent    string
	StoreID                  string
}
type Legacy_ShippingServiceOptions struct {
	ExpeditedService              string
	ImportCharge                  string
	LogisticPlanType              string
	ShippingInsuranceCost         string
	ShippingService               string
	ShippingServiceAdditionalCost string
	ShippingServiceCost           string
	ShippingServicePriority       int
}
type Legacy_ShipmentTrackingDetails struct {
	ShipmentTrackingNumber string
	ShippingCarrierUsed    string
}
type Legacy_InternationalShippingServiceOption struct {
	//	ImportCharge  string
	//     ShippingInsuranceCost string
	//     ShippingService string
	//      ShippingServiceAdditionalCost  string
	//     ShippingServiceCost string
	//     ShippingServicePriority int
	//    ShipToLocation string
}


type Legacy_CalculatedShippingRate struct {
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

type Legacy_PickupOptions struct {
	PickupMethod   string
	PickupPriority string
}

type Legacy_RefundArray struct{}
type Legacy_PickupDetails struct {
	PickupOptions Legacy_PickupOptions
}

type Legacy_RequiredSellerActionArray struct {
	RequiredSellerAction string
}
type Legacy_PaymentHoldDetails struct {
	ExpectedReleaseDate        string
	NumOfReqSellerActions      string
	PaymentHoldReason          string
	RequiredSellerActionArrays Legacy_RequiredSellerActionArray
}
type Legacy_ShipToAddress struct {
	XMLName           xml.Name `xml:"ShipToAddress"`
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
	StateOrProvince   string
	Street1           string
	Street2           string
}
type Legacy_SellerShipmentToLogisticsProvider struct {
	ShippingServiceDetails Legacy_ShippingServiceDetails
	ShipToAddress Legacy_ShipToAddress
	ShippingTimeMax int
	ShippingTimeMin int
}
type Legacy_ShippingServiceDetails struct {
	ShippingService   string
	TotalShippingCost string
}
type Legacy_MultiLegShippingDetails struct {
	SellerShipmentToLogisticsProvider Legacy_SellerShipmentToLogisticsProvider
}
type Legacy_MonetaryDetails struct {
	Payments Legacy_Payments
	Refunds lefacy.Refunds
}

type Legacy_Payments struct {
	Payment []Legacy_Payment
}

type Legacy_Refunds struct {
	Refund Legacy_Refund
}
type Legacy_Refund struct {
	FeeOrCreditAmount  string
	//ReferenceID       string
	RefundAmount       string
	RefundStatus       string
	RefundTime         string
	RefundTo           string
	RefundType         string

	//RefundAmount string
	RefundFromSeller   string
	//	RefundID string
	//	RefundStatus string
	//	RefundTime string
	TotalRefundToBuyer string
}

type Legacy_SalesTax struct {
	SalesTaxAmount        string
	SalesTaxPercent       string
	SalesTaxState         string
	ShippingIncludedInTax string
}

type Legacy_Payment struct {
	FeeOrCreditAmount  string
	Payee              string
	Payer              string
	PaymentAmount      string
	PaymentReferenceID string
	PaymentStatus      string
	PaymentTime        string
	//ReferenceID        string
}

type Legacy_ExternalTransaction struct {
	ExternalTransactionID     string
	ExternalTransactionStatus string
	ExternalTransactionTime   string
	FeeOrCreditAmount         string
	PaymentOrRefundAmount     string
}
type Legacy_BuyerPackageEnclosure struct {
	BuyerPackageEnclosureType string
}

type Legacy_BuyerTaxIdentifier struct {
	Attribute string
	ID        string
	Type      string
}

type Legacy_CancelDetail struct {
	CancelCompleteDate  string
	CancelIntiationDate string
	CancelIntiator      string
	CancelReason        string
	CancelReasonDetails string
}

type Legacy_CheckoutStatus struct {
	eBayPaymentStatus                   string
	IntegratedMerchantCreditCardEnabled string
	LastModifiedTime                    string
	PaymentInstrument                   string
	PaymentMethod                       string
	Status                              string
}

func (o *Legacy_OrderIDArray) Legacy_GetOrderId(orderid string) {
	o.OrderID = orderid

}


func (o *Legacy_GetOrdersRequest) Legacy_FetchOrders(c Config) Legacy_GetOrdersRequestResponse {
	o.Xmlns = "urn:ebay:apis:eBLBaseComponents"
	o.Version = Version
	OrdersXml, err := xml.Marshal(o)

	if err != nil {
		colour.Println("^1 ERROR - xml.Marshal : " + err.Error())
		return GetOrdersRequestResponse{}
	}

	xmlheader := []byte("<?xml version=\"1.0\" encoding=\"utf-8\"?>")

	body := append(xmlheader, OrdersXml...)

	req := goreq.Request{
		Method:      "POST",
		Uri:         "https://api.ebay.com/ws/api.dll",
		Body:        body,
		ContentType: "application/xml; charset=utf-8",
		UserAgent:   "go-ebay-fetch-orders",
		ShowDebug:   false,
	}

	req.AddHeader("X-EBAY-API-CALL-NAME", "GetOrders")
	req.AddHeader("X-EBAY-API-DEV-NAME", c.DevID)
	req.AddHeader("X-EBAY-API-CERT-NAME", c.CertID)
	req.AddHeader("X-EBAY-API-APP-NAME", c.AppID)
	req.AddHeader("X-EBAY-API-VERSION", o.Version)
	req.AddHeader("X-EBAY-API-COMPATIBILITY-LEVEL", o.Version)
	req.AddHeader("X-EBAY-API-REQUEST-ENCODING", "XML")
	req.AddHeader("X-EBAY-API-RESPONSE-ENCODING", "XML")
	req.AddHeader("X-EBAY-API-SITEID", "0")
	req.AddHeader("X-EBAY-API-COMPATIBILITY-LEVEL HTTP", o.Version)

	req.AddHeader("Accept", "application/xml,application/xhtml+xml")
	req.AddHeader("X-Powered-By", "go-ebay (https://goo.gl/Zi7RMK)")
	req.AddHeader("X-Author", "Jason Clark (mithereal@gmail.com)")

	res, err := req.Do()

	if err != nil {
		colour.Println("^1 ERROR - processUrl -> req.Do: " + err.Error())
		return Legacy_GetOrdersRequestResponse{}
	}


	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		colour.Println("^1 ERROR - ioutil.ReadAll : " + err.Error())
		return Legacy_GetOrdersRequestResponse{}
	}

	Response := Legacy_GetOrdersRequestResponse{

	}

	Orders := Legacy_OrderArray{}


	xml.Unmarshal(data, &Response)
	xml.Unmarshal(data, &Orders)

	e := Response.Errors

	if e.ShortMessage != "" {
		colour.Println("^1 ERROR - " + e.ErrorCode + " : " + e.LongMessage)
		return Legacy_GetOrdersRequestResponse{}
	}

	return Response
}

func (o *Legacy_GetOrdersRequest) _Legacy_FetchOrders() Legacy_GetOrdersRequestResponse {
	o.Xmlns = "urn:ebay:apis:eBLBaseComponents"
	o.Version = Version

	usr, err := user.Current()

	Config, _ := NewConfig(path.Join(usr.HomeDir, ".ebayapi"))

	OrdersXml, err := xml.Marshal(o)

	if err != nil {
		colour.Println("^1 ERROR - xml.Marshal : " + err.Error())
		return Legacy_GetOrdersRequestResponse{}
	}

	xmlheader := []byte("<?xml version=\"1.0\" encoding=\"utf-8\"?>")

	body := append(xmlheader, OrdersXml...)

	req := goreq.Request{
		Method:      "POST",
		Uri:         "https://api.ebay.com/ws/api.dll",
		Body:        body,
		ContentType: "application/xml; charset=utf-8",
		UserAgent:   "go-ebay-fetch-orders",
		ShowDebug:   false,
	}

	req.AddHeader("X-EBAY-API-CALL-NAME", "GetOrders")
	req.AddHeader("X-EBAY-API-DEV-NAME", Config.DevID)
	req.AddHeader("X-EBAY-API-CERT-NAME", Config.CertID)
	req.AddHeader("X-EBAY-API-APP-NAME", Config.AppID)
	req.AddHeader("X-EBAY-API-VERSION", o.Version)
	req.AddHeader("X-EBAY-API-COMPATIBILITY-LEVEL", o.Version)
	req.AddHeader("X-EBAY-API-REQUEST-ENCODING", "XML")
	req.AddHeader("X-EBAY-API-RESPONSE-ENCODING", "XML")
	req.AddHeader("X-EBAY-API-SITEID", "0")
	req.AddHeader("X-EBAY-API-COMPATIBILITY-LEVEL HTTP", o.Version)

	req.AddHeader("Accept", "application/xml,application/xhtml+xml")
	req.AddHeader("X-Powered-By", "go-ebay (https://goo.gl/Zi7RMK)")
	req.AddHeader("X-Author", "Jason Clark (mithereal@gmail.com)")

	res, err := req.Do()

	if err != nil {
		colour.Println("^1 ERROR - processUrl -> req.Do: " + err.Error())
		return Legacy_GetOrdersRequestResponse{}
	}


	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		colour.Println("^1 ERROR - ioutil.ReadAll : " + err.Error())
		return Legacy_GetOrdersRequestResponse{}
	}

	Response := Legacy_GetOrdersRequestResponse{

	}

	Orders := Legacy_OrderArray{}


	xml.Unmarshal(data, &Response)
	xml.Unmarshal(data, &Orders)

	e := Response.Errors

	if e.ShortMessage != "" {
		colour.Println("^1 ERROR - " + e.ErrorCode + " : " + e.LongMessage)
		return Legacy_GetOrdersRequestResponse{}
	}

	return Response
}