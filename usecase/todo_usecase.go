package usecase

import (
	"github.com/s-koyasu/todo-app/domain/entities"
	"github.com/s-koyasu/todo-app/domain/repository"
)

type TodoUsecase struct {
	repo repository.TodoRepository
}

func NewTodoUsecase(repo repository.TodoRepository) *TodoUsecase {
	return &TodoUsecase{
		repo: repo,
	}
}

func (u *TodoUsecase) GetAllTodos() ([]entities.Todo, error) {
	return u.repo.FindAll()
}

func (u *TodoUsecase) GetTodoByID(id string) (*entities.Todo, error) {
	return u.repo.FindByID(id)
}

func (u *TodoUsecase) CreateTodo(todo *entities.Todo) error {
	return u.repo.Save(todo)
}

func (u *TodoUsecase) UpdateTodo(todo *entities.Todo) error {
	return u.repo.Update(todo)
}

func (u *TodoUsecase) DeleteTodo(id string) error {
	return u.repo.Delete(id)
}
