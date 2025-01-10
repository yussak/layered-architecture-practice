package domain

import "server/infrastructure"

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
