package control

import (
	"config"
	"fmt"
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
}

func GetControl() *Control {
	return &Ctrl
}
