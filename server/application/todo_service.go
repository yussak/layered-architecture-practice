package application

// application層の役割
// アプリのユースケースを書く
// リクエスト内容がアプリの仕様にあっているかの確認はここで行う
// todoをcreateする、likeをaddするなどのユースケースなどが書かれる

import (
	"server/domain"
)

// 直接ui, applicationで依存させず、間にinterfaceを噛ませることでそれを変更してモックしやすくなる
// ui - interface - application という関係になるので、interfaceを変えやすくなるということか

// GetTodosという名前で、 ([]domain.Todo, error)を返す関数ならこのインターフェースを満たすことを指定
// 具体的な実装は持っていないため、実装を切り替えできる
type TodoService interface {
	GetTodos() ([]domain.Todo, error)
}

// 依存性の注入時に&application.TodoServiceImplで呼ぶことで、TodoServiceImpl構造体の具体的な関数を呼べる
type TodoServiceImpl struct {
	// Domain側も依存性注入可能にしている
	Domain domain.TodoDomain
}

// TodoServiceImpl構造体の関数であると示す
func (s *TodoServiceImpl) GetTodos() ([]domain.Todo, error) {
	return s.Domain.GetTodos()
}

func CreateTodo(name string) (domain.Todo, error) {
	return domain.CreateTodo(name)
}

func DeleteTodo(id string) error {
	return domain.DeleteTodo(id)
}
