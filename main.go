package main

import (
	"fmt"
	controller "project/Conroller"
	menu "project/Menu"
)

func main() {
	menu.PrintMenu()
	var c string
	fmt.Scanln(&c)
	controller.Control(c)
}
