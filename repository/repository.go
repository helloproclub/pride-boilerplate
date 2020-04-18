package repository

import (
	"github.com/helloproclub/pride-boilerplate/domain"
	uuid "github.com/satori/go.uuid"
)

type TodoRepository interface {
	Save(todo domain.Todo) (domain.Todo, error)
	GetById(id uuid.UUID) (domain.Todo, error)
	FindAll(limit, offset int) ([]domain.Todo, error)
	DeleteByID(id uuid.UUID) (domain.Todo, error)
	Update(todo domain.Todo) (domain.Todo, error)
}
