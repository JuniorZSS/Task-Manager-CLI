package main

import "fmt"

var tasks []string

func addTast() {
	fmt.Print("Добавь новую задачу, введи её название: ")
	var task string
	fmt.Scanln(&task)
	tasks = append(tasks, task)
	fmt.Println("Задача добавлена!")
}

func main() {
	fmt.Println("Приветствую, это начало разработки Task Manager CLI!")
}
