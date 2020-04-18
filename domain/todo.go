package domain

import (
	uuid "github.com/satori/go.uuid"
)

type Todo struct {
	Id          uuid.UUID
	Title       string
	Description string
}

func NewTodo(title, description string) Todo {
	id := uuid.NewV4()

	return Todo{
		Id:          id,
		Title:       title,
		Description: description,
	}
}
