package controller

import (
	"fmt"
	menu "project/Menu"
	todolist "project/ToDoList"
)

func Control(comand string) {
	switch comand {
	case "help":
		todolist.Help()
		Home()
	case "add":
		todolist.Add()
		Home()
	case "list":
		todolist.List()
		Home()
	case "del":
		todolist.Del()
		Home()
	case "done":
		todolist.Done()
		Home()
	case "events":
		Home()
	case "exit":
		todolist.Exit()
		Home()
	}
}

func Home() {
	menu.PrintMenu()
	var c string
	fmt.Scanln(&c)
	Control(c)
}
