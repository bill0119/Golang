package control

import (
	"config"
	"fmt"
	"webServer"
)

type Control struct {
}

var (
	Ctrl Control
)

func ControlInit() {
	fmt.Printf("Control initial...\n")
	conf := config.GetWebConfig()
	fmt.Println(conf)
	webServer.StartServer(conf.HttpPort)
}

func GetControl() *Control {
	return &Ctrl
}
