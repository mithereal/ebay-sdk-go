package main

import "flag"
import "fmt"
import "gopkg.in/alecthomas/kingpin.v2"

func getopt() {
	
	DisplayToPublic := flag.String("DisplayToPublic", "2007-12-01T20:34:44", "a string")
	EndCreationTime := flag.String("EndCreationTime", "2007-12-01T20:34:44", "a string")
	MailMessageType := flag.String("MailMessageType", "2007-12-01T20:34:44", "a string")
	MessageStatus := flag.String("MessageStatus", "2007-12-01T20:34:44", "a string")
	SenderID := flag.String("SenderID", "2007-12-01T20:34:44", "a string")
	StartCreationTime := flag.String("StartCreationTime", "2007-12-01T20:34:44", "a string")
	ErrorLanguage := flag.String("ErrorLanguage", "2007-12-01T20:34:44", "a string")
	MessageID := flag.String("MessageID", "2007-12-01T20:34:44", "a string")
	OutputSelector := flag.String("OutputSelector", "2007-12-01T20:34:44", "a string")
	ItemID := flag.Int("ItemID", 42, "an int")
	
	
	flag.Parse()

	fmt.Println("DisplayToPublic: ", *DisplayToPublic)
	fmt.Println("EndCreationTime: ", *EndCreationTime)
	fmt.Println("ItemID: ", *ItemID)
	fmt.Println("MailMessageType: ", *MailMessageType)
	fmt.Println("MessageStatus: ", *MessageStatus)
	fmt.Println("SenderID: ", *SenderID)
	fmt.Println("StartCreationTime: ", *StartCreationTime)
	fmt.Println("ErrorLanguage: ", *ErrorLanguage)
	fmt.Println("MessageID: ", *MessageID)
	fmt.Println("OutputSelector: ", *OutputSelector)
	fmt.Println("ItemID: ", *ItemID)
}

var (
	DisplayToPublic     = kingpin.Flag("DisplayToPublic", "DisplayToPublic .").Short('d').String()
	EndCreationTime = kingpin.Flag("EndCreationTime", ".").Short('e').String()
	MailMessageType = kingpin.Flag("MailMessageType", ".").Short('t').String()
	MessageStatus = kingpin.Flag("MessageStatus", ".").Short('f').String()
	SenderID = kingpin.Flag("SenderID", ".").Short('s').String()
	StartCreationTime = kingpin.Flag("StartCreationTime", ".").Short('b').String()
	ErrorLanguage = kingpin.Flag("ErrorLanguage", ".").Short('l').String()
	MessageID = kingpin.Flag("MessageID", "Message ID.").Short('m').String()
	OutputSelector = kingpin.Flag("OutputSelector", ".").Short('o').String()

	ItemID = kingpin.Flag("ItemID", ".").Short('s').String()
	
)
