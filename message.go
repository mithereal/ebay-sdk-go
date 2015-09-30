package ebay

type GetMemberMessagesRequest struct {
	Xmlns string `xml:"xmlns,attr"`

	RequesterCredentials RequesterCredentials `xml:"RequesterCredentials"`

	DisplayToPublic   string
	EndCreationTime   string
	ItemID            string
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

	Pagination
}

type GetMemberMessagesRequestResponse struct {
	Xmlns                 string `xml:"xmlns,attr"`
	Timestamp             string
	Ack                   string
	Version               string
	Build                 string
	Errors                Errors
	HasMoreMessages       bool
	CorrelationID         string
	HardExpirationWarning string

	Pagination
	MemberMessageExchange 
}

type MemberMessageExchange struct {
	Item             []Item
	LastModifiedDate string
	MessageStatus    string
	Response         string
	MessageMedia     []MessageMedia
	Question
}

type MemberMessage struct {
	MediaName string
	MediaURL string
}

type MessageMedia struct {
	MemberMessageExchange []MemberMessageExchange
}

type Question struct {
	Body            string
	DisplayToPublic string
	MessageID       string
	MessageMedia    []MessageMedia
	MessageType     string
	QuestionType    string
	RecipientID     string
	SenderEmail     string
	SenderID        string
	Subject         string
}
