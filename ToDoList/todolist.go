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
	fmt.Println("- add  - эта команда позволяет добавлять новые задачи в список задач")
	fmt.Println("- list - эта команда позволяет получить полный список всех задач")
	fmt.Println("- del - эта команда позволяет удалить задачу по её заголовку")
	fmt.Println("- done - эта команда позволяет отменить задачу как выполненную")
	fmt.Println("- events - эта команда позволяет получить список всех событий")
	fmt.Println("- exit - эта команда позволяет завершить выполнение программы")
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
