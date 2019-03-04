package ebay

type legacy.GetMemberMessagesRequest struct {
	Xmlns string `xml:"xmlns,attr"`

	RequesterCredentials legacy.RequesterCredentials `xml:"RequesterCredentials"`

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

	Pagination legacy.Pagination
}

type legacy.GetMemberMessagesRequestResponse struct {
	Xmlns                 string `xml:"xmlns,attr"`
	Timestamp             string
	Ack                   string
	Version               string
	Build                 string
	Errors                Errors
	HasMoreMessages       bool
	CorrelationID         string
	HardExpirationWarning string

	Pagination	legacy.Pagination
	MemberMessageExchange legacy.MemberMessageExchange
}

type legacy.MemberMessageExchange struct {
	Item             []Item
	LastModifiedDate string
	MessageStatus    string
	Response         string
	MessageMedia     []legacy.MessageMedia
	Question
}

type legacy.MemberMessage struct {
	MediaName string
	MediaURL string
}

type legacy.MessageMedia struct {
	MemberMessageExchange []legacy.MemberMessageExchange
}

type legacy.Question struct {
	Body            string
	DisplayToPublic string
	MessageID       string
	MessageMedia    []legacy.MessageMedia
	MessageType     string
	QuestionType    string
	RecipientID     string
	SenderEmail     string
	SenderID        string
	Subject         string
}
