package control

import (
	"control/webServer"
	"fmt"
)

type Control struct {
}

var (
	Ctrl Control
)

func ControlInit() {
	fmt.Printf("Control initial...\n")
	webConf := webServer.GetWebConfig()
	fmt.Println(webConf)
	webServer.StartServer(webConf)
}

func GetControl() *Control {
	return &Ctrl
}
