package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

var tasks []string

type Task struct {
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func addTask() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Добавь новую задачу, введи её название: ")
	task, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}
	task = task[:len(task)-1] // Удаляем символ новой строки (\n)
	tasks = append(tasks, task)
	fmt.Println("Задача добавлена!")
}

func main() {
	fmt.Println("Приветствую, это начало разработки Task Manager CLI!")
	for {
		fmt.Println("1. Добавить задачу")
		fmt.Println("2. Просмотреть задачи")
		fmt.Println("3. Удалить задачу")
		fmt.Println("4. Завершить задачу")
		fmt.Println("5. Выйти")
		fmt.Print("Выберите действие: ")

		var choice int
		if _, err := fmt.Scanln(&choice); err != nil { // Добавляем ссылку на переменную choice
			fmt.Println("Ошибка ввода:", err)
			return
		}

		switch choice {
		case 1:
			addTask()
		case 2:
			viewTasks()
		case 3:
			deleteTask()
		case 4:
			completeTask()
		case 5:
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}

func viewTasks() {
	if len(tasks) == 0 {
		fmt.Println("Пусто, задач нет")
		return
	}
	fmt.Println("Список задач:")
	for i, task := range tasks {
		fmt.Println("%d. %s\n", i+1, task)
	}

}

func deleteTask() {
	var taskIndex int
	fmt.Print("Введите номер задачи для удаления: ")
	if _, err := fmt.Scanln(&taskIndex); err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	if taskIndex < 1 || taskIndex > len(tasks) {
		fmt.Println("Некорректный номер задачи.")
		return
	}

	tasks = append(tasks[:taskIndex-1], tasks[taskIndex:]...)
	fmt.Println("Задача удалена.")
}

func completeTask() {
	var taskIndex int
	fmt.Print("Введите номер задачи для завершения: ")
	if _, err := fmt.Scanln(&taskIndex); err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	if taskIndex < 1 || taskIndex > len(tasks) {
		fmt.Println("Задачи с таким номером не существует")
		return
	}

	tasks[taskIndex-1] = tasks[taskIndex-1] + " (Завершена)"
	fmt.Println("Задача завершена.")
}

func saveTaskToFile(filename string) {
	// Создаём или перезаписывем задачи
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Ошибка при создании файла: ", err)
		return
	}
	defer file.Close()

	// Преобразование задачи в json
	data, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("Ошибка при создании json файла: ", err)
		return
	}

	// Сохраняем задачи в файл
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Ошибка при записи в файл: ", err)
		return
	}

	fmt.Println("Задачи успешно сохранены!")
}