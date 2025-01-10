package infrastructure

import (
	"server/db"
)

type Todo struct {
	ID   int
	Name string
}

func GetTodosFromDB() ([]Todo, error) {
	rows, err := db.DB.Query("SELECT id, name FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		todo := Todo{}
		if err := rows.Scan(&todo.ID, &todo.Name); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}
