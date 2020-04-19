package repository

import (
	"github.com/helloproclub/pride-boilerplate/domain"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type postgresRepository struct {
	DB *gorm.DB
}

func NewPostgresTodoRepository(DB *gorm.DB) TodoRepository {
	return postgresRepository{
		DB,
	}
}

func (repo postgresRepository) Save(todo domain.Todo) (domain.Todo, error) {
	if err := repo.DB.Model(&todo).Create(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}

func (repo postgresRepository) GetById(id uuid.UUID) (domain.Todo, error) {
	var todo domain.Todo

	if err := repo.DB.Model(&todo).Where("id = ?", id).Find(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}

func (repo postgresRepository) FindAll(limit, offset int) ([]domain.Todo, error) {
	var todos []domain.Todo
	if err := repo.DB.Model(&todos).
		Limit(limit).
		Offset(offset).
		Find(&todos).Error; err != nil {
		return todos, err
	}

	return todos, nil
}

func (repo postgresRepository) Update(todo domain.Todo) (domain.Todo, error) {
	if err := repo.DB.Model(&todo).Where("id = ?", todo.Id).Update(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}

func (repo postgresRepository) DeleteByID(id uuid.UUID) (domain.Todo, error) {
	var todo domain.Todo

	if err := repo.DB.Model(&todo).Where("id = ?", id).Delete(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}
