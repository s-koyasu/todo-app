package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/s-koyasu/todo-app/domain/entities"
	"github.com/s-koyasu/todo-app/infrastructure/persistence/inmemory"
	"github.com/s-koyasu/todo-app/usecase"
)

func main() {
	// 依存関係のワイヤリング
	todoRepo := inmemory.NewTodoRepository()
	todoUsecase := usecase.NewTodoUsecase(todoRepo)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command (create, list, update, delete, exit): ")
		input, _ := reader.ReadString('\n')
		command := strings.TrimSpace(input)

		switch command {
		case "create":
			fmt.Print("Enter ToDo title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			todo := &entities.Todo{
				Title: title,
			}
			err := todoUsecase.CreateTodo(todo)
			if err != nil {
				fmt.Println("Error creating ToDo:", err)
			} else {
				fmt.Println("ToDo created:", todo.ID)
			}
		case "list":
			todos, err := todoUsecase.GetAllTodos()
			if err != nil {
				fmt.Println("Error getting ToDos:", err)
			} else {
				fmt.Println("ToDos:")
				for _, todo := range todos {
					fmt.Printf("  %s: %s\n", todo.ID, todo.Title)
				}
			}
		case "update":
			fmt.Print("Enter ToDo ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Enter new ToDo title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			todo := &entities.Todo{
				ID:    id,
				Title: title,
			}
			err := todoUsecase.UpdateTodo(todo)
			if err != nil {
				fmt.Println("Error updating ToDo:", err)
			} else {
				fmt.Println("ToDo updated:", todo.ID)
			}
		case "delete":
			fmt.Print("Enter ToDo ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			err := todoUsecase.DeleteTodo(id)
			if err != nil {
				fmt.Println("Error deleting ToDo:", err)
			} else {
				fmt.Println("ToDo deleted:", id)
			}
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid command.")
		}
	}
}
