package domain

import (
	uuid "github.com/satori/go.uuid"
)

type Todo struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

func NewTodo(title, description string) Todo {
	id := uuid.NewV4()

	return Todo{
		Id:          id,
		Title:       title,
		Description: description,
	}
}
