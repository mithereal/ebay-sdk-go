package ebay

import "encoding/xml"
import "github.com/franela/goreq"
import "io/ioutil"
import "github.com/alecthomas/colour"
import  "os/user"
import "path"

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
	ModifiedTime                         string
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
	//Buyer  Buyer
	ShippingDetails ShippingDetails
	CreatedDate string
	Item Item
	QuantityPurchased string
//	Status Status
	TransactionID string
	TransactionPrice string
	//ShippingServiceSelected ShippingServiceSelected
	FinalValueFee string
	TransactionSiteID string
	Platform string
	//Taxes Taxes
	ActualShippingCost float32
	ActualHandlingCost float32
	OrderLineItemID string
	ExtendedOrderID string
	eBayPlusTransaction bool

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


func  (o *GetOrdersRequest) FetchOrders(c Config) GetOrdersRequestResponse {

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
		return GetOrdersRequestResponse{}
		}


	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		colour.Println("^1 ERROR - ioutil.ReadAll : " + err.Error())
		return GetOrdersRequestResponse{}
	}

	Response := GetOrdersRequestResponse{

	}

	Orders := OrderArray{}


	xml.Unmarshal(data, &Response)
	xml.Unmarshal(data, &Orders)

	e := Response.Errors

	if e.ShortMessage != "" {
		colour.Println("^1 ERROR - " + e.ErrorCode + " : " + e.LongMessage)
		return GetOrdersRequestResponse{}
	}

	return Response
}

func  (o *GetOrdersRequest) _FetchOrders() GetOrdersRequestResponse {

	usr, err := user.Current()

	Config, _ := NewConfig(path.Join(usr.HomeDir, ".ebayapi"))

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
		return GetOrdersRequestResponse{}
	}


	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		colour.Println("^1 ERROR - ioutil.ReadAll : " + err.Error())
		return GetOrdersRequestResponse{}
	}

	Response := GetOrdersRequestResponse{

	}

	Orders := OrderArray{}


	xml.Unmarshal(data, &Response)
	xml.Unmarshal(data, &Orders)

	e := Response.Errors

	if e.ShortMessage != "" {
		colour.Println("^1 ERROR - " + e.ErrorCode + " : " + e.LongMessage)
		return GetOrdersRequestResponse{}
	}

	return Response
}