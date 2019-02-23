package main

import "flag"
import "fmt"
import "gopkg.in/alecthomas/kingpin.v2"

func getopt() {
	timePtrf := flag.String("timefrom", "2007-12-01T20:34:44", "a string")
	timePtrt := flag.String("timeto", "2007-12-01T20:34:44", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	IncludeFinalValueFeePtr := flag.Bool("IncludeFinalValueFee", false, "true/false")

	flag.Parse()

	fmt.Println("timefrom: ", *timePtrf)
	fmt.Println("timeto: ", *timePtrt)
	fmt.Println("numb: ", *numbPtr)
	fmt.Println("IncludeFinalValueFee: ", *IncludeFinalValueFeePtr)
}

var (
	verbose     = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	listingtype = kingpin.Flag("listingtype", "Type of listing (Auction).").Short('l').String()

	createtimefrom = kingpin.Flag("createtimefrom", "Time the first order was created.").Short('s').String()
	createtimeto   = kingpin.Flag("createtimeto", "Time the last order was created.").Short('e').String()
)
