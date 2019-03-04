package ebay

import (
	"encoding/xml"

)

const (
	Version = "933"
)

type legacy.ErrorMessage struct {
	XmlName xml.Name `xml:"errorMessage"`
	Error   Error    `xml:"error"`
}

type legacy.Errors struct {
	XmlName      xml.Name `xml:"Errors"`
	ShortMessage string   `xml:"ShortMessage"`
	LongMessage  string   `xml:"LongMessage"`
	ErrorCode    string   `xml:"ErrorCode"`
}

type legacy.Error struct {
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

type legacy.EBay struct {
	ApplicationId string
}

type legacy.Item struct {
	ItemITitleD string
	ItemID string
	ListingDetails 
	SellingStatus 
}

type legacy.ListingDetails struct {
	EndTime       string
	StartTime       string
	}

type legacy.SellingStatus struct {
	CurrentPrice       string
	}

type legacy.Pagination struct {
	EntriesPerPage       string
	OrdersPerPage        string
	PageNumber           string
	TotalNumberOfEntries int
	TotalNumberOfPages   int
}

	
func legacy.New(application_id string) *legacy.EBay {
	e := legacy.EBay{}
	e.ApplicationId = application_id

	return &e
}

func (p *legacy.Pagination) legacy.Setpagination(entriesperpage, pagenumber string) {
	p.EntriesPerPage = entriesperpage
	p.PageNumber = pagenumber
}

