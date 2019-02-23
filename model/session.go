package ebay

import (
	"encoding/json"
	"os"
)

type Config struct {
	AppID  string `json:"appID"`  // Application ID
	CertID string `json:"certID"` // Application Certification
	DevID  string `json:"devID"`  // Developer ID
	RuName string `json:"ruName"` // RuName
	Token  string `json:"Token"`  // Token
}

func NewConfig(cfgName string) (*Config, error) {
	cfgFile, err := os.Open(cfgName)
	if err != nil {
		cfg := Config{
			AppID:  os.Getenv("EBAY_APP_ID"),
			DevID:  os.Getenv("EBAY_DEV_ID"),
			CertID: os.Getenv("EBAY_CERT_ID"),
			Token:  os.Getenv("EBAY_TOKEN"),
		}
		return &cfg, nil

	}
	decoder := json.NewDecoder(cfgFile)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

