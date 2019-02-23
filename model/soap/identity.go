// Code generated by gowsdl DO NOT EDIT.

package ebay.soap

import (
	"encoding/xml"
	"github.com/hooklift/gowsdl/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

//
// Used to specify whether a token is passed by value or by reference.
// Applications should expect it to be passed by value, and should pass it by
// value. Later implementations will allow tokens to be passed by reference, to
// accomodate lengthy assertions in case there is a need to pass encrypted tokens in
// URLs.
//
type OpeneBayTokenType string

const (

	//
	// The token is passed by reference.
	//
	OpeneBayTokenTypeReference OpeneBayTokenType = "Reference"

	//
	// The token is passed by value.
	//
	OpeneBayTokenTypeValue OpeneBayTokenType = "Value"
)

//
// Indicates whether the error is a severe error (causing the request to fail) or an
// informational error (a warning) that should be communicated to the user.
//

type AuthenticateRequest OpeneBayAuthenticateRequestType

type AuthenticateResponse OpeneBayAuthenticateResponseType

type LoginRequest OpeneBayLoginRequestType

type LoginResponse OpeneBayLoginResponseType

type OpeneBayPrincipalIdentifierType struct {

	//
	// AppID, or fully-qualified name of an application as required in the deployment
	// descriptor. This value must be structured as
	// ApplicationName.PlatformName.TopDomain.
	//
	AppId string `xml:"appId,attr,omitempty"`
}

type OpeneBayAcceptType struct {

	//
	// Container for specifying whether a token is passed by reference or by value.
	//
	By *OpeneBayTokenType `xml:"by,attr,omitempty"`
}

type OpeneBayLoginCredentialsType struct {
	*OpeneBayPrincipalIdentifierType

	//
	// The certID is generated by eBay and provided to third parties who use Selling Manager Applications. Third
	// parties use the certId to log into eBay through the EIDP service.
	//
	CertId string `xml:"certId,omitempty"`
}

type OpeneBaySecurityTokenType struct {
	*OpeneBayBaseSecurityTokenType

	//
	// The EIDP token identifying the third-party application. Obtained using
	// the login request.
	//
	TokenValue string `xml:"tokenValue,omitempty"`

	//
	// Reserved for future use.
	//
	Signature string `xml:"signature,omitempty"`
}

type OpeneBayBaseSecurityTokenType struct {

	//
	// Base type for tokens.
	//
	Type *OpeneBayTokenType `xml:"type,attr,omitempty"`

	//
	// Issuing authority for the token. Value is eBay.
	//
	IssuingAuthority string `xml:"issuingAuthority,attr,omitempty"`

	//
	// The time and date that the EIDP token expires (typically, 24 hours after
	// issuance).
	//
	ExpirationDate time.Time `xml:"expirationDate,attr,omitempty"`
}


type OpeneBayAuthenticateRequestType struct {
	XMLName xml.Name `xml:"http://www.ebay.com/marketplace/services authenticateRequest"`

	*BaseServiceRequest

	//
	// The EIDP token to be validated. The EIDP token is included in the HTTP header
	// to validate calls from your application to eBay, such as OEAIS service
	// operations.
	//
	TokenToValidate string `xml:"tokenToValidate,omitempty"`
}

type OpeneBayAuthenticateResponseType struct {
	XMLName xml.Name `xml:"http://www.ebay.com/marketplace/services authenticateResponse"`

	*BaseServiceResponse

	//
	// Internal use only.
	//
	Attributes *OpeneBayAssertionAttributesType `xml:"attributes,omitempty"`
}

type OpeneBayAssertionAttributeType struct {

	//
	// Contains the values of the identity assertion, usually a DevID, used in the login
	// call.
	//
	AttributeValue string `xml:"attributeValue,omitempty"`

	//
	// Specifies which identity assertion is used in the login call. For example, DevID.
	//
	Name string `xml:"name,attr,omitempty"`
}

type OpeneBayAssertionAttributesType struct {

	//
	// Container for the identity assertions used in the login call.
	//
	Attribute []*OpeneBayAssertionAttributeType `xml:"attribute,omitempty"`
}

type OpeneBayLoginRequestType struct {
	XMLName xml.Name `xml:"http://www.ebay.com/marketplace/services loginRequest"`

	*BaseServiceRequest

	//
	// Specifies the type of credentials offered in the login request.
	//
	Credential *OpeneBayLoginCredentialsType `xml:"credential,omitempty"`

	//
	// Specifies whether acceptance will be based on a token passed by value or
	// by reference.
	//
	Accept *OpeneBayAcceptType `xml:"accept,omitempty"`

	//
	// The identity assertions submitted at login include CertID and AppID.
	//
	Attributes *OpeneBayAssertionAttributesType `xml:"attributes,omitempty"`
}

type OpeneBayLoginResponseType struct {
	XMLName xml.Name `xml:"http://www.ebay.com/marketplace/services loginResponse"`

	*BaseServiceResponse

	//
	// Container for EIDP token, returned in the login response. This token accompanies all subsequent calls made to
	// eBay by Selling Manager applications.
	//
	SecurityToken *OpeneBaySecurityTokenType `xml:"securityToken,omitempty"`
}

type OpeneBayIdentityProviderServicePort interface {
	Authenticate(request *OpeneBayAuthenticateRequestType) (*OpeneBayAuthenticateResponseType, error)

	Login(request *OpeneBayLoginRequestType) (*OpeneBayLoginResponseType, error)
}

type openeBayIdentityProviderServicePort struct {
	client *soap.Client
}

func NewOpeneBayIdentityProviderServicePort(client *soap.Client) OpeneBayIdentityProviderServicePort {
	return &openeBayIdentityProviderServicePort{
		client: client,
	}
}

func (service *openeBayIdentityProviderServicePort) Authenticate(request *OpeneBayAuthenticateRequestType) (*OpeneBayAuthenticateResponseType, error) {
	response := new(OpeneBayAuthenticateResponseType)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *openeBayIdentityProviderServicePort) Login(request *OpeneBayLoginRequestType) (*OpeneBayLoginResponseType, error) {
	response := new(OpeneBayLoginResponseType)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}