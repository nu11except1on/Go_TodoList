package models

import "errors"

type Task struct {
	id          int
	title       string
	description string
	status      Status
}

func NewTask(id int, title string, description string, status Status) *Task {
	return &Task{
		id:          id,
		title:       title,
		description: description,
		status:      status,
	}
}

func (t Task) Title() string {
	return t.title
}

func (t *Task) SetTitle(title string) error {
	if len(title) == 0 {
		return errors.New("Empty title")
	}
	t.title = title
	return nil
}

func (t Task) Description() string {
	return t.description
}

func (t *Task) SetDescription(description string) error {
	if len(description) == 0 {
		return errors.New("Empty description")
	}
	t.description = description
	return nil
}

func (t Task) ID() int {
	return t.id
}

func (t *Task) SetId(id int) error {
	if id < 0 {
		return errors.New("Invalid id")
	}
	t.id = id
	return nil
}

func (t Task) Status() Status {
	return t.status
}

func (t *Task) SetStatus(status Status) error {
	t.status = status
	return nil
}
