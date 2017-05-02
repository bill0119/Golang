package main

import (
	"control"
	"fmt"
)

func init() {
	fmt.Printf("main initial...\n")
	control.ControlInit()
	ctrl := control.GetControl()
	fmt.Println(ctrl)
}

func main() {

}
