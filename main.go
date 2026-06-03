package main

import (
	"TodoList/models"
	"fmt"
	"log"
)

func main() {
	var manager TaskManager = NewInMemoryTaskManager()

	newTask := models.NewTask(1, "Изучить Go", "Написать менеджер задач на Go", models.StatusNew)
	fmt.Printf("Создана задача: %s (ID: %d)\n", newTask.Title(), newTask.ID())

	err := manager.AddTask(newTask)
	if err != nil {
		log.Fatalf("Не удалось добавить задачу: %v", err)
	}

	foundTask, err := manager.GetTask(1)
	if err != nil {
		log.Fatalf("Ошибка при поиске задачи: %v", err)
	}
	fmt.Println("\n--- Результат поиска ---")
	fmt.Printf("Заголовок: %s,\nОписание: %s,\nСтатус: %v\n", foundTask.Title(), foundTask.Description(), foundTask.Status())

	_ = foundTask.SetStatus(models.StatusInProgress)
	fmt.Printf("Новый статус задачи: %v\n", foundTask.Status())
}
