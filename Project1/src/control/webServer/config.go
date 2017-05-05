package webServer

type WebConfig struct {
	HttpPort  int
	HttpsPort int
}

var (
	WebConf WebConfig
)

func init() {
	WebConf.HttpPort = 80
	WebConf.HttpsPort = 443
}

func GetWebConfig() *WebConfig {
	return &WebConf
}
