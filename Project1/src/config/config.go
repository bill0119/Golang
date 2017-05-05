package config

import (
	"fmt"
)

type WebConfig struct {
	HttpPort  int
	HttpsPort int
}

var (
	WebConf WebConfig
)

func init() {
	fmt.Printf("Initial Config\n")
	WebConf.HttpPort = 80
	WebConf.HttpsPort = 443
}

func GetWebConfig() *WebConfig {
	return &WebConf
}
