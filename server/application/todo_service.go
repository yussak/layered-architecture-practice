package application

// application層の役割
// アプリのユースケースを書く
// リクエスト内容がアプリの仕様にあっているかの確認はここで行う

// todoをcreateする、likeをaddするなどのユースケースなどが書かれる

import (
	"server/domain"
)

func GetTodos() ([]domain.Todo, error) {
	return domain.GetTodos()
}

func CreateTodo(name string) (domain.Todo, error) {
	return domain.CreateTodo(name)
}

func DeleteTodo(id string) error {
	return domain.DeleteTodo(id)
}
