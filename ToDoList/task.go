package todolist

import (
	"encoding/json"
	"fmt"
	"os"
)

const fileName = "todo_list.json"

type Task struct {
	Flag bool   `json:"flag"`
	Name string `json:"name"`
	Text string `json:"text"`
}

func createFile(fn string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("[]")
	if err != nil {
		return err
	}

	return nil
}

func Add() {
	var n, t string
	fmt.Println("Введите название задачи")
	fmt.Scanln(&n)
	if n == "" {
		fmt.Println("Название не может быть пустым")
	}
	fmt.Println("Введите содержание задачи")
	fmt.Scanln(&t)
	if t == "" {
		fmt.Println("Содержание не может быть пустым")
	}
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		//создаем новый файл
		if err := createFile(fileName); err != nil {
			fmt.Println("Ошибка при создании файла")
			return
		}
		fmt.Println("Файл успешно создан")
	} else if err != nil {
		fmt.Println("Произошла ошибка припроверке файла")
		return
	} else {
		fmt.Println("Файл найден")

		todo := ToDoList{}

		data, err := os.ReadFile("todo_list.json")
		if err != nil {
			fmt.Println("Файл не удалось прочитать")
			return
		}

		json.Unmarshal(data, &todo)

		todo.Tasks = append(todo.Tasks, Task{
			Flag: false,
			Name: n,
			Text: t,
		})
		if err := todo.SaveToFile(fileName); err != nil {
			fmt.Println("Ошибка записи в файл")
			return
		}
		fmt.Println("Задачи успешно сохранены в", fileName)
	}
}
