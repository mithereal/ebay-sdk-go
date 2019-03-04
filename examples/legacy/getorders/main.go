package main

import "github.com/mithereal/go-ebay"
import "github.com/alecthomas/colour"
import "github.com/davecgh/go-spew/spew"
import "gopkg.in/alecthomas/kingpin.v2"
import "os/user"
import "path"
import "strconv"


const (
	Version        = "1.0"

)

func main() {

	kingpin.Parse()

	usr, err := user.Current()

	if err != nil {
		colour.Println("^1 ERROR - processUser -> user.Current: " + err.Error())
		return
	}

	//Config is * optional *, its for use with custom config data stores ie. db
	//example shown is for custom config data stored in  ~/.ebayapi or in environment variables


	Config, _ := ebay.NewConfig(path.Join(usr.HomeDir, ".ebayapi"))

	Credentials := ebay.Legacy_RequesterCredentials{
		RequestToken: Config.Token,
	}


	OrdersRequest := ebay.Legacy_GetOrdersRequest{
		CreateTimeFrom:       *createtimefrom,
		CreateTimeTo:         *createtimeto,
		DetailLevel:          "ReturnAll",
		OrderRole:            "Seller",
		IncludeFinalValueFee: "true",
		OrderStatus:          "Completed",
		// Credentials are used for custom config data store
		RequesterCredentials: Credentials,
	}
	// Response := OrdersRequest._FetchOrders()  for use with default config data store
	Response := OrdersRequest.FetchOrders(*Config)



	// do something with the orders //
	spew.Dump(Response.OrdersArray)

	switch Response.Ack {
	case "Success":
		colour.Printf(" ^6 Response: ^2" + Response.Ack +" ^R \n")

		orderactualcount := strconv.Itoa(Response.ReturnedOrderCountActual)
		pagecount := strconv.Itoa(Response.Paginations.TotalNumberOfPages)
		pagenumber :=strconv.Itoa(Response.PageNumber)
		totalentries := strconv.Itoa(Response.Paginations.TotalNumberOfEntries)
		hasMoreOrders := strconv.FormatBool(Response.HasMoreOrders)



		colour.Printf(" ^6 Total Pages: ^2" +  pagecount +" ^R \n")
		colour.Printf(" ^6 Current Page: ^2" +  pagenumber +" ^R \n")
		colour.Printf(" ^6 Page Entries: ^2" +  orderactualcount +" ^R \n")
		colour.Printf(" ^6 Total Orders: ^2" +  totalentries +" ^R \n")
		colour.Printf(" ^6 Has more orders: ^2" +  hasMoreOrders +" ^R \n")

	case "Failure":
		colour.Printf(" ^6 Response: ^1" + Response.Ack +" ^R ")
	default:
		colour.Printf(" ^6 Response: ^3" + Response.Ack +" ^R ")
	}


}
