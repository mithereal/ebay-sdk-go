package ebay

import (
	"encoding/json"
	"os"
)

type Config struct {
	AppID     string `json:"appID"`       // Application ID
	CertID    string `json:"certID"`      // Application Certification
	DevID     string `json:"devID"`       // Developer ID
	RuName    string `json:"ruName"`      // RuName
	Endpoint  string `json:"ApiEndpoint"` // API Endpoint to use
	IsSandbox bool   `json:"isSandbox"`   // List if channels the bot will join after the connect
}

func NewConfig(cfgName string) (error, *Config) {
	cfgFile, err := os.Open(cfgName)
	if err != nil {
		return err, nil
	}
	decoder := json.NewDecoder(cfgFile)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return err, nil
	}
	return nil, &cfg
}

type RequesterCredentials struct {
	RequestToken string `xml:"eBayAuthToken"`
}

func (r *RequesterCredentials) SetToken(token string) {
	r.RequestToken = token
}
