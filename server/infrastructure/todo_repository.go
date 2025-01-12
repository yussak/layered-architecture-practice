package infrastructure

// infrastructure層の役割
// データベース、外部サービスとのやり取りを行う

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

// TODO:責務が別れてるかもしれないので確認
func GetInsertedTodoID(name string) (int, error) {
	// TodosテーブルにINSERTして、INSERTしたレコードのIDを取得
	var insertedID int
	err := db.DB.QueryRow(
		"INSERT INTO Todos (name) VALUES ($1) RETURNING id",
		name,
	).Scan(&insertedID)

	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

func Delete(id string) error {
	_, err := db.DB.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
