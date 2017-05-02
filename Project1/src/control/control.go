package control

import (
	"fmt"
)

type Control struct{
}

var(
	ctrl Control
)

func ControlInit() {
	fmt.Printf("Control initial...\n")
}

func GetControl()(Control) {
	return ctrl
}
