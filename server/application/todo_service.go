package application

// application層の役割
// アプリのユースケースを書く
// リクエスト内容がアプリの仕様にあっているかの確認はここで行う

import (
	"errors"
	"server/domain"
)

func GetTodos() ([]domain.Todo, error) {
	return domain.GetTodos()
}

func CreateTodo(name string) (domain.Todo, error) {
	if name == "" {
		return domain.Todo{}, errors.New("nameが空です")
	}

	return domain.CreateTodo(name)
}

func DeleteTodo(id string) error {
	return domain.DeleteTodo(id)
}
