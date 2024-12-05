package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestSaveTasksToFile(t *testing.T) {
	// Подготовка тестовых данных
	tasks := []string{"Задача 1", "Задача 2", "Задача 3"}
	testFile := "test_tasks.json"

	// Вызов функцию сохранению
	saveTaskToFile(testFile)

	// Проверка создания файла
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Fatalf("Файл %s не был создан", testFile)
	}

	// Чтение содержимого файла
	content, err := ioutil.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Не удалочь прочитать файл %s: %v", testFile, err)
	}

	// Проверка содержимого файла
	expected := `["Задача 1","Задача 2", "Задача 3"]`
	if string(content) != expected {
		t.Errorf("Ожидалось: %s, Получено: %s", expected, string(content))
	}

	// Удаление тестового файла после тест
	os.Remove(testFile)
}
