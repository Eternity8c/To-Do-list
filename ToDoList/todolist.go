package todolist

import (
	"encoding/json"
	"fmt"
	"os"
)

type ToDoList struct {
	Tasks []Task `json:"tasks"`
}

func (tl *ToDoList) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(tl, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func Help() {

}

func List() {
	data, err := os.ReadFile("todo_list.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла")
		return
	}

	todo := ToDoList{}
	json.Unmarshal(data, &todo)
	fmt.Println(todo)
}
