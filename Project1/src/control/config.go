package control

import (
	"control/webServer"
)

type ControlConfig struct {
	WebConfig webServer.WebConfig
}

var (
	CtrlConf ControlConfig
)
