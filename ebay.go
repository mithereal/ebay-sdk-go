package ebay

import (
	//"fmt"
	"encoding/xml"
	//"strconv"
	//"net/url"
	//"errors"
	//"github.com/franela/goreq"
)

const (
	Version = "933"
)

type ErrorMessage struct {
	XmlName xml.Name `xml:"errorMessage"`
	Error   Error    `xml:"error"`
}

type Errors struct {
	XmlName xml.Name `xml:"Errors"`
	Error   Error    `xml:"error"`
}

type Error struct {
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

type EBay struct {
	ApplicationId string
}

func New(application_id string) *EBay {
	e := EBay{}
	e.ApplicationId = application_id

	return &e
}
