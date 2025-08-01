package controller

import (
	todolist "project/ToDoList"
)

func Control(comand string) {
	switch comand {
	case "help":
		todolist.Help()
	case "add":
		todolist.Add()
	case "list":
		todolist.List()
	case "del":
	case "done":
	case "events":
	case "exit":
	}
}
