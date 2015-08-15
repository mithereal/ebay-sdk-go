package main

import "github.com/mithereal/go-ebay"
import "github.com/davecgh/go-spew/spew"
import "github.com/franela/goreq"
import "github.com/alecthomas/colour"
import "gopkg.in/alecthomas/kingpin.v2"
import "encoding/xml"
import "bytes"
import "os"

const (
	Version        = "1.0"
	EbayApiVersion = "933"
)

func main() {

	kingpin.Parse()

	EbayAppId := os.Getenv("EBAY_PRODUCTION_APP_ID")
	Token := os.Getenv("EBAY_PRODUCTION_TOKEN")

	RequesterCredentials := ebay.RequesterCredentials{}
	RequesterCredentials.SetToken(Token)

	OrdersRequest := &ebay.GetOrdersRequest{
		Xmlns:                "urn:ebay:apis:eBLBaseComponents",
		RequesterCredentials: RequesterCredentials,
		CreateTimeFrom:       "2015-08-14T20:34:44.000Z",
		CreateTimeTo:         "2015-08-13T20:34:44.000Z",
		OrderRole:            "Seller",
		OrderStatus:          "Completed",
		Version:              EbayApiVersion,
	}

	xml, _ := xml.Marshal(OrdersRequest)
	body := bytes.NewBuffer(xml)

	req := goreq.Request{
		Method:      "POST",
		Uri:         "https://api.ebay.com/ws/api.dll",
		Body:        body,
		ContentType: "application/xml; charset=utf-8",
		UserAgent:   "go-ebay-fetch-orders",
		ShowDebug:   true,
	}

	req.AddHeader("X-EBAY-API-CALL-NAME", "GetOrdersRequest")
	req.AddHeader("X-EBAY-API-APP-ID", EbayAppId)
	req.AddHeader("X-EBAY-API-VERSION", EbayApiVersion)
	req.AddHeader("X-EBAY-API-REQUEST-ENCODING", "XML")
	req.AddHeader("X-EBAY-API-RESPONSE-ENCODING", "XML")
	req.AddHeader("X-EBAY-API-SITE-ID", "0")
	req.AddHeader("X-EBAY-API-COMPATIBILITY-LEVEL HTTP", EbayApiVersion)

	req.AddHeader("Accept", "application/xml,application/xhtml+xml")
	req.AddHeader("X-Powered-By", "go-ebay (https://goo.gl/Zi7RMK)")
	req.AddHeader("X-Author", "Jason Clark (mithereal@gmail.com)")

	res, err := req.Do()

	if err != nil {
		colour.Println("^1 ERROR - processUrl -> req.Do: " + err.Error())
		return
	}

	response := ebay.GetOrdersRequestResponse{}

	response_xml, _ := xml.Unmarshal(res.Body, &response)

	spew.Dump(res.Body.ToString())

}
