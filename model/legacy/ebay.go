package ebay

import (
	"encoding/xml"

)

const (
	Version = "933"
)

type Legacy_ErrorMessage struct {
	XmlName xml.Name `xml:"errorMessage"`
	Error   Error    `xml:"error"`
}

type Legacy_Errors struct {
	XmlName      xml.Name `xml:"Errors"`
	ShortMessage string   `xml:"ShortMessage"`
	LongMessage  string   `xml:"LongMessage"`
	ErrorCode    string   `xml:"ErrorCode"`
}

type Legacy_Error struct {
	ErrorId      string `xml:"errorId"`
	Domain       string `xml:"domain"`
	SeverityCode string `xml:"SeverityCode"`
	Severity     string `xml:"Severity"`
	Line         string `xml:"Line"`
	Column       string `xml:"Column"`
	Category     string `xml:"Category"`
	Message      string `xml:"ShortMessage"`
	SubDomain    string `xml:"subdomain"`
}

type Legacy_EBay struct {
	ApplicationId string
}

type Legacy_Item struct {
	ItemITitleD string
	ItemID string
	ListingDetails 
	SellingStatus 
}

type Legacy_ListingDetails struct {
	EndTime       string
	StartTime       string
	}

type Legacy_SellingStatus struct {
	CurrentPrice       string
	}

type Legacy_Pagination struct {
	EntriesPerPage       string
	OrdersPerPage        string
	PageNumber           string
	TotalNumberOfEntries int
	TotalNumberOfPages   int
}

type Legacy_RequesterCredentials struct {
	RequestToken string `xml:"eBayAuthToken"`
}

func (r *Legacy_RequesterCredentials) Legacy_SetToken(token string) {
	r.RequestToken = token
}


	
func Legacy_New(application_id string) *Legacy_EBay {
	e := Legacy_EBay{}
	e.ApplicationId = application_id

	return &e
}

func (p *Legacy_Pagination) Legacy_Setpagination(entriesperpage, pagenumber string) {
	p.EntriesPerPage = entriesperpage
	p.PageNumber = pagenumber
}

