package inmemory

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/s-koyasu/todo-app/domain/entities"
	"github.com/s-koyasu/todo-app/domain/repository"
)

type todoRepository struct {
	todos map[string]entities.Todo
}

func NewTodoRepository() repository.TodoRepository {
	return &todoRepository{
		todos: make(map[string]entities.Todo),
	}
}

func (r *todoRepository) FindAll() ([]entities.Todo, error) {
	todoList := make([]entities.Todo, 0, len(r.todos))
	for _, todo := range r.todos {
		todoList = append(todoList, todo)
	}
	return todoList, nil
}

func (r *todoRepository) FindByID(id string) (*entities.Todo, error) {
	todo, ok := r.todos[id]
	if !ok {
		return nil, errors.New("ToDo not found")
	}
	return &todo, nil
}

func (r *todoRepository) Save(todo *entities.Todo) error {
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	newID := uuid.New().String()
	todo.ID = newID
	r.todos[newID] = *todo

	return nil
}

func (r *todoRepository) Update(todo *entities.Todo) error {
	_, ok := r.todos[todo.ID]
	if !ok {
		return errors.New("ToDo not found")
	}
	todo.UpdatedAt = time.Now()
	r.todos[todo.ID] = *todo
	return nil
}

func (r *todoRepository) Delete(id string) error {
	_, ok := r.todos[id]
	if !ok {
		return errors.New("ToDo not found")
	}
	delete(r.todos, id)
	return nil
}
