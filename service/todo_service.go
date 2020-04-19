package service

import (
	"errors"

	"github.com/helloproclub/pride-boilerplate/domain"
	"github.com/helloproclub/pride-boilerplate/repository"
	uuid "github.com/satori/go.uuid"
)

type TodoService struct {
	repository repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return TodoService{repository: repo}
}

func (service *TodoService) CreateNewTodo(title, description string) (domain.Todo, error) {
	todo, err := service.repository.Save(domain.NewTodo(title, description))
	if err != nil {
		return todo, errors.New("failed to create new todo: " + err.Error())
	}

	return todo, nil
}

func (service *TodoService) GetById(id uuid.UUID) (domain.Todo, error) {
	todo, err := service.repository.GetById(id)
	if err != nil {
		return todo, errors.New("failed to get todo " + id.String() + " : " + err.Error())
	}

	return todo, nil
}

func (service *TodoService) FindAll(limit, offset int) ([]domain.Todo, error) {
	todos, err := service.repository.FindAll(limit, offset)
	if err != nil {
		return todos, errors.New("failed to find todos: " + err.Error())
	}

	return todos, nil
}

func (service *TodoService) DeleteByID(id uuid.UUID) (domain.Todo, error) {
	todo, err := service.repository.DeleteByID(id)
	if err != nil {
		return todo, errors.New("failed to delete todo" + id.String() + " : " + err.Error())
	}

	return todo, nil
}

func (service *TodoService) Update(todo domain.Todo) (domain.Todo, error) {
	todo, err := service.repository.Update(todo)
	if err != nil {
		return todo, errors.New("failed to update todo: " + err.Error())
	}

	return todo, nil
}
