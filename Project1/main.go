package main

import (
	"control"
	"fmt"
)

func init() {
	fmt.Printf("main initial...\n")
	//ctrl := control.GetControl()
	//fmt.Println(ctrl)

	control.ControlInit()
}

func main() {

}
