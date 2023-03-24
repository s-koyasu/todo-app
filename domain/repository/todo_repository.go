package repository

import "github.com/s-koyasu/todo-app/domain/entities"

type TodoRepository interface {
	FindAll() ([]entities.Todo, error)
	FindByID(id string) (*entities.Todo, error)
	Save(todo *entities.Todo) error
	Update(todo *entities.Todo) error
	Delete(id string) error
}
