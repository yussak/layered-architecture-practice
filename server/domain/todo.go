package domain

// domain層の役割
// アプリのビジネルスールや中心となるロジックを管理する

import (
	"errors"
	"server/infrastructure"
)

type Todo struct {
	ID   int
	Name string
}

func GetTodos() ([]Todo, error) {
	infraTodos, err := infrastructure.GetTodosFromDB()
	if err != nil {
		return nil, err
	}

	// infraのTodo型からdomainのTodo型に変換
	var todos []Todo
	for _, infraTodo := range infraTodos {
		todos = append(todos, Todo{
			ID:   infraTodo.ID,
			Name: infraTodo.Name,
		})
	}

	return todos, nil
}

func CreateTodo(name string) (Todo, error) {
	if name == "" {
		return Todo{}, errors.New("nameが空です")
	}

	insertedID, err := infrastructure.GetInsertedTodoID(name)
	if err != nil {
		return Todo{}, err
	}

	todo := Todo{
		ID:   insertedID,
		Name: name,
	}

	return todo, nil
}

func DeleteTodo(id string) error {
	err := infrastructure.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
