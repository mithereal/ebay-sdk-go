package main

import "github.com/mithereal/go-ebay"
import "github.com/franela/goreq"
import "github.com/alecthomas/colour"
import "github.com/davecgh/go-spew/spew"
import "gopkg.in/alecthomas/kingpin.v2"
import "encoding/xml"
import "io/ioutil"
import "os/user"
import "path"

const (
	Version        = "1.0"
	EbayApiVersion = "935"
)

func main() {

	kingpin.Parse()

	usr, err := user.Current()

	Config, _ := ebay.NewConfig(path.Join(usr.HomeDir, ".ebayapi"))

	Credentials := ebay.legacy.RequesterCredentials{
		RequestToken: Config.Token,
	}

	MemberMessagesRequest := ebay.legacy.GetMemberMessagesRequest{
		Xmlns:             "urn:ebay:apis:eBLBaseComponents",
		DisplayToPublic:   *DisplayToPublic,
		EndCreationTime:   *EndCreationTime,
		ItemID:            *ItemID,
		MailMessageType:   *MailMessageType,
		MessageStatus:     *MessageStatus,
		SenderID:          *SenderID,
		StartCreationTime: *StartCreationTime,
		ErrorLanguage:     *ErrorLanguage,
		MessageID:         *MessageID,
		OutputSelector:    *OutputSelector,

		Version:              EbayApiVersion,
		RequesterCredentials: Credentials,
	}

	MemberMessagesXml, err := xml.Marshal(MemberMessagesRequest)

	if err != nil {
		colour.Println("^1 ERROR - xml.Marshal : " + err.Error())
		return
	}

	xmlheader := []byte("<?xml version=\"1.0\" encoding=\"utf-8\"?>")

	body := append(xmlheader, MemberMessagesXml...)

	req := goreq.Request{
		Method:      "POST",
		Uri:         "https://api.ebay.com/ws/api.dll",
		Body:        body,
		ContentType: "application/xml; charset=utf-8",
		UserAgent:   "go-ebay-fetch-orders",
		ShowDebug:   false,
	}

	req.AddHeader("X-EBAY-API-CALL-NAME", "GetMemberMessages")
	req.AddHeader("X-EBAY-API-DEV-NAME", Config.DevID)
	req.AddHeader("X-EBAY-API-CERT-NAME", Config.CertID)
	req.AddHeader("X-EBAY-API-APP-NAME", Config.AppID)
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

	var Response ebay.GetMemberMessagesRequestResponse

	xml.Unmarshal(data, &Response)

	e := Response.Errors

	if e.ShortMessage != "" {
		colour.Println("^1 ERROR - " + e.ErrorCode + " : " + e.LongMessage)
		return
	}

	spew.Dump(Response)
}
