package main

import "github.com/mithereal/go-ebay"	
import "github.com/davecgh/go-spew/spew"
import "github.com/franela/goreq"
//import "github.com/alecthomas/colour"
import "gopkg.in/alecthomas/kingpin.v2"
import "log"
import "encoding/xml"

const (
        
        Version = "1.0"
        
)

var (
  verbose = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
  listingtype    = kingpin.Flag("listingtype", "Type of listing (Auction).").Short('t').String()
)

func main() {
	
  kingpin.Parse()
  //colour.Printf(" %v,  %s\n", *verbose, *listingtype)
  
   OrdersRequest := ebay.GetOrdersRequest  {
    Xmlns:           "urn:ebay:apis:eBLBaseComponents",
    ListingType:  *listingtype,
    CreateTimeFrom:   "2007-12-01T20:34:44.000Z" ,
    CreateTimeTo:   "2007-12-10T20:34:44.000Z",
	OrderRole:   "Seller" ,
	OrderStatus:   "Completed" ,

   }
   
   OrdersRequest.RequesterCredentials.SetToken("token")

xml,_ := xml.Marshal(OrdersRequest)


req := goreq.Request{
	Method: "POST",
	Uri: "https://api.ebay.com/ws/api.dll",
    Body: xml,
    ContentType: "application/xml; charset=utf-8",
    UserAgent: "go-ebay-fetchorders",
    ShowDebug:   true,
	}
	
	req.AddHeader("Accept", "application/xml,application/xhtml+xml")


	response, err := req.Do()
	
	if err != nil {
		log.Println("ERROR - processUrl -> req.Do: " + err.Error())
		return
	}


spew.Dump(response.Body.ToString())

		}
		


