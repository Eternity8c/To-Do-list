package todolist

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/k0kubun/pp"
)

const fileName = "todo_list.json"

type Task struct {
	Flag bool      `json:"flag"`
	Name string    `json:"name"`
	Text string    `json:"text"`
	Time time.Time `json:"time"`
}

type ToDoList struct {
	Tasks []Task `json:"tasks"`
}

func createFile(fn string) error {
	file, err := os.Create(fn)
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

func (tl *ToDoList) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(tl, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

// Выводит информацию о меню
func Help() {
	fmt.Println("- add  - эта команда позволяет добавлять новые задачи в список задач")
	fmt.Println("- list - эта команда позволяет получить полный список всех задач")
	fmt.Println("- del - эта команда позволяет удалить задачу по её заголовку")
	fmt.Println("- done - эта команда позволяет отменить задачу как выполненную")
	fmt.Println("- events - эта команда позволяет получить список всех событий")
	fmt.Println("- exit - эта команда позволяет завершить выполнение программы")
}

func Add() {
	var n, t string
	fmt.Println("Введите название задачи")
	reader := bufio.NewReader(os.Stdin)
	n, _ = reader.ReadString('\n') // Читаем до символа новой строки
	// Удаляем символ новой строки в конце
	n = strings.TrimSpace(n)
	if n == "" {
		fmt.Println("Название не может быть пустым")
	}
	fmt.Println("Введите содержание задачи")
	t, _ = reader.ReadString('\n') // Читаем до символа новой строки
	t = strings.TrimSpace(t)
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
			Time: time.Now(),
		})
		if err := todo.SaveToFile(fileName); err != nil {
			fmt.Println("Ошибка записи в файл")
			return
		}
		fmt.Println("Задачи успешно сохранены в", fileName)
	}
}

func List() {
	data, err := os.ReadFile("todo_list.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла")
		return
	}

	todo := ToDoList{}
	json.Unmarshal(data, &todo)
	pp.Println(todo)
}

func Del() {
	var n string
	fmt.Println("Введите название задачи, которую хотите удалить")
	reader := bufio.NewReader(os.Stdin)
	n, _ = reader.ReadString('\n') // Читаем до символа новой строки
	// Удаляем символ новой строки в конце
	n = strings.TrimSpace(n)
	if n == "" {
		fmt.Println("Неверный ввод")
		return
	}
	data, err := os.ReadFile("todo_list.json")
	if err != nil {
		fmt.Println("Ошибка чтения файл")
	}

	todo := ToDoList{}
	json.Unmarshal(data, &todo)
	for i, name := range todo.Tasks {
		if name.Name == n {
			todo.Tasks = append(todo.Tasks[:i], todo.Tasks[i+1:]...)
			todo.SaveToFile(fileName)
			fmt.Println("Задача успешно удалена")
			return
		}
	}
	fmt.Println("Нету задачи с таким именем")
}

func Done() {
	var n string
	fmt.Println("Введите название задачи, которая выполнена")
	reader := bufio.NewReader(os.Stdin)
	n, _ = reader.ReadString('\n') // Читаем до символа новой строки
	// Удаляем символ новой строки в конце
	n = strings.TrimSpace(n)
	if n == "" {
		fmt.Println("Неверный ввод")
		return
	}
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Ошибка чтения файла")
		return
	}
	todo := ToDoList{}
	json.Unmarshal(data, &todo)
	for i, name := range todo.Tasks {
		if name.Name == n && !name.Flag {
			todo.Tasks[i] = Task{
				Flag: true,
				Name: name.Name,
				Text: name.Text,
				Time: time.Now(),
			}
			todo.SaveToFile(fileName)
			fmt.Println("Флаг изменен")
			return
		}
	}
	fmt.Println("Задача уже выполнена")
}

func Exit() {
	os.Exit(0)
}
