package main 
import "flag"
import "fmt"

func getopt(){
timePtrf := flag.String("timefrom", "2007-12-01T20:34:44.000Z", "a string")
	timePtrt := flag.String("timeto", "2007-12-01T20:34:44.000Z", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
    IncludeFinalValueFeePtr := flag.Bool("IncludeFinalValueFee", false, "true/false")
    
	flag.Parse()
	
	fmt.Println("timefrom: ", *timePtrf)
	fmt.Println("timeto: ", *timePtrt)
    fmt.Println("numb: ", *numbPtr)
    fmt.Println("IncludeFinalValueFee: ", *IncludeFinalValueFeePtr)
    }
    
