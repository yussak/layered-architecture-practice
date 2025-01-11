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

type TodoResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetNewTodo(name string) (Todo, error) {
	insertedID, err := infrastructure.GetInsertedTodoID(name)
	if err != nil {
		return Todo{}, err
	}

	// 登録したTODOをJSONで返す
	newTodo := Todo{
		ID:   insertedID,
		Name: name,
	}

	return newTodo, nil
}

func DeleteTodo(id string) error {
	err := infrastructure.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
