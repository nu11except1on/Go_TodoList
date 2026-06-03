package main

import (
	"TodoList/models"
	"errors"
)

type InMemoryTaskManager struct {
	taskHashMap map[int]*models.Task
}

func NewInMemoryTaskManager() TaskManager {
	return &InMemoryTaskManager{
		taskHashMap: make(map[int]*models.Task),
	}
}

func (m *InMemoryTaskManager) AddTask(task *models.Task) error {
	if task == nil {
		return errors.New("task cannot be nil")
	}
	m.taskHashMap[task.ID()] = task
	return nil
}

func (m *InMemoryTaskManager) GetTask(id int) (*models.Task, error) {
	task, exists := m.taskHashMap[id]
	if !exists {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (m *InMemoryTaskManager) DeleteTask(id int) error {
	_, exists := m.taskHashMap[id]
	if !exists {
		return errors.New("task not found")
	}
	delete(m.taskHashMap, id)
	return nil
}

func (m *InMemoryTaskManager) UpdateTask(id int, newTitle string, newDescription string) error {
	task, exists := m.taskHashMap[id]
	if !exists {
		return errors.New("task not found")
	}
	if err := task.SetTitle(newTitle); err != nil {
		return err
	}
	if err := task.SetDescription(newDescription); err != nil {
		return err
	}
	return nil
}

func (m *InMemoryTaskManager) UpdateTaskStatus(id int, newStatus string) error {
	task, exists := m.taskHashMap[id]
	if !exists {
		return errors.New("task not found")
	}
	if err := task.SetStatus(models.Status(newStatus)); err != nil {
		return err
	}
	return nil
}
