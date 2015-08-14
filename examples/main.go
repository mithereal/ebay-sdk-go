package main

import "github.com/mithereal/go-ebay"	
import "github.com/davecgh/go-spew/spew"
import "github.com/franela/goreq"
import "github.com/alecthomas/colour"
import "gopkg.in/alecthomas/kingpin.v2"
import "encoding/xml"
import "bytes"

const (
        
        Version = "1.0"
		EbayApiVersion = "933"
		EbayAppId =""
		Token = ""
        
)

var (
    verbose     = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	listingtype = kingpin.Flag("listingtype", "Type of listing (Auction).").Short('l').String()

	createtimefrom = kingpin.Flag("createtimefrom", "Time the first order was created.").Short('s').String()
	createtimeto   = kingpin.Flag("createtimeto", "Time the last order was created.").Short('e').String()
)

func main() {
	
 kingpin.Parse()
  
 OrdersRequest := &ebay.GetOrdersRequest  {
  		Xmlns:          "urn:ebay:apis:eBLBaseComponents",
		CreateTimeFrom: "2007-12-01T20:34:44.000Z",
		CreateTimeTo:   "2007-12-10T20:34:44.000Z",
		OrderRole:      "Seller",
		OrderStatus:    "Completed",
		Version:        EbayApiVersion,
   }
   
OrdersRequest.RequesterCredentials.SetToken(Token)


xml,_ := xml.Marshal(OrdersRequest)
body := bytes.NewBuffer(xml)


req := goreq.Request{
	Method: "POST",
	Uri: "https://api.ebay.com/ws/api.dll",
    Body: body,
    ContentType: "application/xml; charset=utf-8",
    UserAgent: "go-ebay-fetch-orders",
    ShowDebug:   true,
	}
	
	req.AddHeader("X-EBAY-API-CALL-NAME", "GetOrdersRequest")
	req.AddHeader("X-EBAY-API-APP-ID", EbayAppId)
	req.AddHeader("X-EBAY-API-VERSION", EbayApiVersion)
	req.AddHeader("X-EBAY-API-REQUEST-ENCODING", "XML")
	req.AddHeader("X-EBAY-API-RESPONSE-ENCODING", "XML")
	req.AddHeader("X-EBAY-API-SITE-ID", "0")
	
	req.AddHeader("Accept", "application/xml,application/xhtml+xml")
	req.AddHeader("X-Powered-By", "go-ebay (https://goo.gl/Zi7RMK)")
	req.AddHeader("X-Author", "Jason Clark (mithereal@gmail.com)")

	response, err := req.Do()
	
	if err != nil {
		colour.Println("ERROR - processUrl -> req.Do: " + err.Error())
		return
	}


spew.Dump(response.Body.ToString())

		}
		


