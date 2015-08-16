package main

import "github.com/mithereal/go-ebay"
import "github.com/franela/goreq"
import "github.com/alecthomas/colour"
import "gopkg.in/alecthomas/kingpin.v2"
import "encoding/xml"
import "bytes"
import "io/ioutil"
import "os"

const (
	Version        = "1.0"
	EbayApiVersion = "933"
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
	//Credentials.SetToken(Token)

	OrdersRequest := ebay.GetOrdersRequest{
		Xmlns:                "urn:ebay:apis:eBLBaseComponents",
		CreateTimeFrom:       "2015-08-13T20:34:44.000Z",
		CreateTimeTo:         "2015-08-14T20:34:44.000Z",
		OrderRole:            "Seller",
		IncludeFinalValueFee: "true",
		OrderStatus:          "Completed",
		Version:              EbayApiVersion,
		RequesterCredentials: Credentials,
	}

	reqxml, err := xml.Marshal(OrdersRequest)

	if err != nil {
		colour.Println("^1 ERROR - xml.Marshal : " + err.Error())
		return
	}

	body := bytes.NewBuffer(reqxml)
	// <?xml version="1.0" encoding="utf-8"?> needs to be added to body
	req := goreq.Request{
		Method:      "POST",
		Uri:         "https://api.ebay.com/ws/api.dll",
		Body:        body,
		ContentType: "application/xml; charset=utf-8",
		UserAgent:   "go-ebay-fetch-orders",
		ShowDebug:   true,
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

	var response ebay.GetOrdersRequestResponse

	xml.Unmarshal(data, &response)

	e := response.Errors

	if e.ShortMessage != "" {
		colour.Println("^1 ERROR - " + e.ErrorCode + " : " + e.LongMessage)
		return
	}

}
