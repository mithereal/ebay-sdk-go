package ebay

import "encoding/xml"
import "github.com/franela/goreq"
import "io/ioutil"
import "github.com/alecthomas/colour"
import "os/user"
import "path"

type GetOrdersRequest struct {
	Xmlns                string `xml:"xmlns,attr"`

	RequesterCredentials legacy.RequesterCredentials `xml:"RequesterCredentials"`

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
	Pagination           legacy.Pagination
	OrderIDArray legacy.OrderIDArray
}

type legacy.GetOrdersRequestResponse struct {
	Xmlns                    string `xml:"xmlns,attr"`
	Timestamp                string `xml:"Timestamp"`
	Ack                      string `xml:"Ack"`
	Build                    string `xml:"Build"`
	Errors                   Errors `xml:"Errors"`
	HasMoreOrders            bool   `xml:"HasMoreOrders"`

	CorrelationID            string `xml:"CorrelationID"`
	HardExpirationWarning    string `xml:"HardExpirationWarning"`

	Paginations              legacy.Pagination `xml:"PaginationResult"`
	OrdersArray              legacy.OrderArray `xml:"OrderArray"`

	OrdersPerPage            int `xml:"OrdersPerPage"`
	PageNumber               int `xml:"PageNumber"`
	ReturnedOrderCountActual int `xml:"ReturnedOrderCountActual"`
}

type legacy.Amount struct {
}
type legacy.BuyerPackageEnclosures struct {
}
type legacy.PaymentInstructionCode struct {
}
type legacy.TaxIdentifierAttributeCode struct {
}
type legacy.ValueTypeCode struct {
}
type legacy.OrderIDArray struct {
	OrderID     string
	BuyerUserID string
}

type legacy.OrderArray struct {
	XMLName xml.Name `xml:"OrderArray"`
	Orders  []legacy.Order  `xml:"Order"`
}

type legacy.Order struct {
	XMLName                             xml.Name `xml:"Order"`
	OrderID                             string
	OrderStatus                         string
	BuyerCheckoutMessage                string
	AmountSaved                         string
	BuyerPackageEnclosures              legacy.BuyerPackageEnclosure `xml:"BuyerPackageEnclosure"`
	BuyerTaxIdentifiers                 legacy.BuyerTaxIdentifier `xml:"BuyerTaxIdentifier"`
	BuyerUserID                         string
	CancelDetails                       legacy.CancelDetail `xml:"CancelDetail"`
	CancelReason                        string
	CancelReasonDetails                 string
	CancelStatus                        string
	CheckoutStatuss                     legacy.CheckoutStatus `xml:"CheckoutStatus"`
	ContainseBayPlusTransaction         bool
	CreatedTime                         string
	ModifiedTime                        string
	CreatingUserRole                    string
	EIASToken                           string
	ExtendedOrderID                     string
	ExternalTransactions                legacy.ExternalTransaction `xml:"ExternalTransaction"`
	IntegratedMerchantCreditCardEnabled string
	IsMultiLegShipping                  string
	LogisticsPlanType                   string
	MonetaryDetailss                    legacy.MonetaryDetails `xml:"MonetaryDetails"`
	MultiLegShippingDetailss            legacy.MultiLegShippingDetails `xml:"MultiLegShippingDetails"`
	PaidTime                            string
	PaymentHoldDetailss                 legacy.PaymentHoldDetails `xml:"PaymentHoldDetails"`
	PaymentHoldStatus                   string
	PaymentMethods                      string
	PickupDetailss                      legacy.PickupDetails `xml:"PickupDetails"`
	PickupMethodSelected
	RefundArrays                        legacy.RefundArray `xml:"RefundArray"`
	SellerEIASToken                     string
	SellerEmail                         string
	SellerUserID                        string
	ShippedTime                         string
	//	ShippingAddress
	ShippingConvenienceCharge           string
	ShippingDetailss                    legacy.ShippingDetails `xml:"ShippingDetails"`
	ShippingServiceSelected             string
	Subtotal                            string
	Total                               float64
	TransactionArrays                   legacy.TransactionArray `xml:"TransactionArray"`
}
type legacy.Transaction struct {
	//Buyer  Buyer
	ShippingDetails     legacy.ShippingDetails
	CreatedDate         legacy.string
	Item                legacy.Item
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
type legacy.TransactionArray struct {
	XMLName      xml.Name `xml:"TransactionArray"`
	Transactions []legacy.Transaction
}

type legacy.PickupMethodSelected struct {
	XMLName               xml.Name `xml:"PickupMethodSelected"`
	MerchantPickupCode    string
	PickupFulfillmentTime string
	PickupLocationUUID    string
	PickupMethod          string
	PickupStatus          string
	PickupStoreID         string
}
type legacy.ShippingDetails struct {
	CalculatedShippingRate legacy.CalculatedShippingRate
	CODCost         string
	InsuranceFee    string
	InsuranceOption string
	InsuranceWanted string

	InternationalShippingServiceOption  legacy.InternationalShippingServiceOption
	//SellingManagerSalesRecordNumber
	ShipmentTrackingDetails legacy.ShipmentTrackingDetails
	ShippingServiceOptions legacy.ShippingServiceOptions
	TaxTable legacy.TaxTable
}
type legacy.TaxJurisdiction struct {
	JurisdictionID        string
	SalesTaxPercent       string
	ShippingIncludedInTax string
}
type legacy.TaxTable struct {
	TaxJurisdiction legacy.TaxJurisdiction
}
type legacy.ShippingPackageInfo struct {
	ActualDeliveryTime       string
	ScheduledDeliveryTimeMax string
	ScheduledDeliveryTimeMin string
	ShippingTrackingEvent    string
	StoreID                  string
}
type legacy.ShippingServiceOptions struct {
	ExpeditedService              string
	ImportCharge                  string
	LogisticPlanType              string
	ShippingInsuranceCost         string
	ShippingService               string
	ShippingServiceAdditionalCost string
	ShippingServiceCost           string
	ShippingServicePriority       int
}
type legacy.ShipmentTrackingDetails struct {
	ShipmentTrackingNumber string
	ShippingCarrierUsed    string
}
type legacy.InternationalShippingServiceOption struct {
	//	ImportCharge  string
	//     ShippingInsuranceCost string
	//     ShippingService string
	//      ShippingServiceAdditionalCost  string
	//     ShippingServiceCost string
	//     ShippingServicePriority int
	//    ShipToLocation string
}


type legacy.CalculatedShippingRate struct {
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

type legacy.PickupOptions struct {
	PickupMethod   string
	PickupPriority string
}

type legacy.RefundArray struct{}
type legacy.PickupDetails struct {
	PickupOptions legacy.PickupOptions
}

type legacy.RequiredSellerActionArray struct {
	RequiredSellerAction string
}
type legacy.PaymentHoldDetails struct {
	ExpectedReleaseDate        string
	NumOfReqSellerActions      string
	PaymentHoldReason          string
	RequiredSellerActionArrays legacy.RequiredSellerActionArray
}
type legacy.ShipToAddress struct {
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
type legacy.SellerShipmentToLogisticsProvider struct {
	ShippingServiceDetails legacy.ShippingServiceDetails
	ShipToAddress legacy.ShipToAddress
	ShippingTimeMax int
	ShippingTimeMin int
}
type legacy.ShippingServiceDetails struct {
	ShippingService   string
	TotalShippingCost string
}
type legacy.MultiLegShippingDetails struct {
	SellerShipmentToLogisticsProvider legacy.SellerShipmentToLogisticsProvider
}
type legacy.MonetaryDetails struct {
	Payments legacy.Payments
	Refunds lefacy.Refunds
}

type legacy.Payments struct {
	Payment []legacy.Payment
}

type legacy.Refunds struct {
	Refund legacy.Refund
}
type legacy.Refund struct {
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

type legacy.SalesTax struct {
	SalesTaxAmount        string
	SalesTaxPercent       string
	SalesTaxState         string
	ShippingIncludedInTax string
}

type legacy.Payment struct {
	FeeOrCreditAmount  string
	Payee              string
	Payer              string
	PaymentAmount      string
	PaymentReferenceID string
	PaymentStatus      string
	PaymentTime        string
	//ReferenceID        string
}

type legacy.ExternalTransaction struct {
	ExternalTransactionID     string
	ExternalTransactionStatus string
	ExternalTransactionTime   string
	FeeOrCreditAmount         string
	PaymentOrRefundAmount     string
}
type legacy.BuyerPackageEnclosure struct {
	BuyerPackageEnclosureType string
}

type legacy.BuyerTaxIdentifier struct {
	Attribute string
	ID        string
	Type      string
}

type legacy.CancelDetail struct {
	CancelCompleteDate  string
	CancelIntiationDate string
	CancelIntiator      string
	CancelReason        string
	CancelReasonDetails string
}

type legacy.CheckoutStatus struct {
	eBayPaymentStatus                   string
	IntegratedMerchantCreditCardEnabled string
	LastModifiedTime                    string
	PaymentInstrument                   string
	PaymentMethod                       string
	Status                              string
}

func (o *legacy.OrderIDArray) legacy.GetOrderId(orderid string) {
	o.OrderID = orderid

}


func (o *legacy.GetOrdersRequest) legacy.FetchOrders(c Config) legacy.GetOrdersRequestResponse {
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
		return legacy.GetOrdersRequestResponse{}
	}


	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		colour.Println("^1 ERROR - ioutil.ReadAll : " + err.Error())
		return legacy.GetOrdersRequestResponse{}
	}

	Response := legacy.GetOrdersRequestResponse{

	}

	Orders := legacy.OrderArray{}


	xml.Unmarshal(data, &Response)
	xml.Unmarshal(data, &Orders)

	e := Response.Errors

	if e.ShortMessage != "" {
		colour.Println("^1 ERROR - " + e.ErrorCode + " : " + e.LongMessage)
		return legacy.GetOrdersRequestResponse{}
	}

	return Response
}

func (o *legacy.GetOrdersRequest) _legacy.FetchOrders() legacy.GetOrdersRequestResponse {
	o.Xmlns = "urn:ebay:apis:eBLBaseComponents"
	o.Version = Version

	usr, err := user.Current()

	Config, _ := NewConfig(path.Join(usr.HomeDir, ".ebayapi"))

	OrdersXml, err := xml.Marshal(o)

	if err != nil {
		colour.Println("^1 ERROR - xml.Marshal : " + err.Error())
		return legacy.GetOrdersRequestResponse{}
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
		return legacy.GetOrdersRequestResponse{}
	}


	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		colour.Println("^1 ERROR - ioutil.ReadAll : " + err.Error())
		return legacy.GetOrdersRequestResponse{}
	}

	Response := legacy.GetOrdersRequestResponse{

	}

	Orders := legacy.OrderArray{}


	xml.Unmarshal(data, &Response)
	xml.Unmarshal(data, &Orders)

	e := Response.Errors

	if e.ShortMessage != "" {
		colour.Println("^1 ERROR - " + e.ErrorCode + " : " + e.LongMessage)
		return legacy.GetOrdersRequestResponse{}
	}

	return Response
}