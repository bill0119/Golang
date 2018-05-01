package webServer

type WebConfig struct {
	HttpPort  int
	HttpsPort int
}

type Data struct {
	Email string
}

var (
	WebConf WebConfig
	WebData Data
)

func init() {
	WebConf.HttpPort = 80
	WebConf.HttpsPort = 443
	WebData.Email = "bill0119@gmail.com"
}

func GetWebConfig() *WebConfig {
	return &WebConf
}

func GetWebData() *Data {
	return &WebData
}
