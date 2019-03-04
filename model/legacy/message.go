package ebay

type Legacy_GetMemberMessagesRequest struct {
	Xmlns string `xml:"xmlns,attr"`

	RequesterCredentials Legacy_RequesterCredentials `xml:"RequesterCredentials"`

	DisplayToPublic   string
	EndCreationTime   string
//	ItemID            string
	MailMessageType   string
	MemberMessageID   string
	MessageStatus     string
	SenderID          string
	StartCreationTime string
	ErrorLanguage     string
	MessageID         string
	OutputSelector    string
	Version           string
	WarningLevel      string

	Pagination Legacy_Pagination
}

type Legacy_GetMemberMessagesRequestResponse struct {
	Xmlns                 string `xml:"xmlns,attr"`
	Timestamp             string
	Ack                   string
	Version               string
	Build                 string
	Errors                Errors
	HasMoreMessages       bool
	CorrelationID         string
	HardExpirationWarning string

	Pagination	Legacy_Pagination
	MemberMessageExchange Legacy_MemberMessageExchange
}

type Legacy_MemberMessageExchange struct {
	Item             []Legacy_Item
	LastModifiedDate string
	MessageStatus    string
	Response         string
	MessageMedia     []Legacy_MessageMedia
	Question Legacy_Question
}

type Legacy_MemberMessage struct {
	MediaName string
	MediaURL string
}

type Legacy_MessageMedia struct {
	MemberMessageExchange []Legacy_MemberMessageExchange
}

type Legacy_Question struct {
	Body            string
	DisplayToPublic string
	MessageID       string
	MessageMedia    []Legacy_MessageMedia
	MessageType     string
	QuestionType    string
	RecipientID     string
	SenderEmail     string
	SenderID        string
	Subject         string
}
