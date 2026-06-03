package main

import "TodoList/models"

type TaskManager interface {
	AddTask(task *models.Task) error

	GetTask(id int) (*models.Task, error)

	DeleteTask(id int) error

	UpdateTask(id int, newTitle string, newDescription string) error

	UpdateTaskStatus(id int, newStatus string) error
}
