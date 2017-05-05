package webServer

type WebConfig struct {
	HttpPort  int
	HttpsPort int
}

var (
	WebConf WebConfig
)

func GetWebConfig() *WebConfig {
	return &WebConf
}
