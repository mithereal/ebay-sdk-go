package main

import "github.com/mithereal/go-ebay"
import "github.com/franela/goreq"
import "github.com/alecthomas/colour"
import "github.com/davecgh/go-spew/spew"
import "gopkg.in/alecthomas/kingpin.v2"
import "encoding/xml"
import "io/ioutil"
import "os"

const (
	Version        = "1.0"
	EbayApiVersion = "935"
)

func main() {

	kingpin.Parse()

	EbayAppId := os.Getenv("EBAY_APP_ID")
	EbayDevName := os.Getenv("EBAY_DEV_ID")
	EbayCertName := os.Getenv("EBAY_CERT_ID")
	Token := os.Getenv("EBAY_TOKEN")

	Credentials := ebay.RequesterCredentials{
		RequestToken: Token,
	}

	OrdersRequest := ebay.GetOrdersRequest{
		Xmlns:                "urn:ebay:apis:eBLBaseComponents",
		CreateTimeFrom:       *createtimefrom,
		CreateTimeTo:         *createtimeto,
		DetailLevel:          "ReturnAll",
		OrderRole:            "Seller",
		IncludeFinalValueFee: "true",
		OrderStatus:          "Completed",
		Version:              EbayApiVersion,
		RequesterCredentials: Credentials,
	}

	OrdersXml, err := xml.Marshal(OrdersRequest)

	if err != nil {
		colour.Println("^1 ERROR - xml.Marshal : " + err.Error())
		return
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
	req.AddHeader("X-EBAY-API-DEV-NAME", EbayDevName)
	req.AddHeader("X-EBAY-API-CERT-NAME", EbayCertName)
	req.AddHeader("X-EBAY-API-APP-NAME", EbayAppId)
	req.AddHeader("X-EBAY-API-VERSION", EbayApiVersion)
	req.AddHeader("X-EBAY-API-COMPATIBILITY-LEVEL", EbayApiVersion)
	req.AddHeader("X-EBAY-API-REQUEST-ENCODING", "XML")
	req.AddHeader("X-EBAY-API-RESPONSE-ENCODING", "XML")
	req.AddHeader("X-EBAY-API-SITEID", "0")
	req.AddHeader("X-EBAY-API-COMPATIBILITY-LEVEL HTTP", EbayApiVersion)

	req.AddHeader("Accept", "application/xml,application/xhtml+xml")
	req.AddHeader("X-Powered-By", "go-ebay (https://goo.gl/Zi7RMK)")
	req.AddHeader("X-Author", "Jason Clark (mithereal@gmail.com)")

	res, err := req.Do()

	if err != nil {
		colour.Println("^1 ERROR - processUrl -> req.Do: " + err.Error())
		return
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		colour.Println("^1 ERROR - ioutil.ReadAll : " + err.Error())
		return
	}

	var Response ebay.GetOrdersRequestResponse

	xml.Unmarshal(data, &Response)

	e := Response.Errors

	if e.ShortMessage != "" {
		colour.Println("^1 ERROR - " + e.ErrorCode + " : " + e.LongMessage)
		return
	}

	spew.Dump(Response)
}
