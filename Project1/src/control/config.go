package control

import (
	"control/webServer"
)

type Control struct {
	Web webServer.WebServer
}

var (
	Ctrl Control
)

func init() {
}
