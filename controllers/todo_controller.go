package controllers

import (
	"log"
	"net/http"
	"server/db"
	"server/models"
	"text/template"
)

// TODO:add todo実装後荷表示されるかを確認する
func ListTodos(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, name FROM todos")
	if err != nil {
		http.Error(w, "データベースエラー", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var todo struct {
			ID   int
			Name string
		}
		if err := rows.Scan(&todo.ID, &todo.Name); err != nil {
			http.Error(w, "データ取得エラー", http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}

	// HTMLテンプレートを読み込む
	tmpl, err := template.ParseFiles("views/index.html")
	if err != nil {
		http.Error(w, "テンプレートエラー", http.StatusInternalServerError)
		log.Printf("テンプレートエラー: %v", err)
		return
	}

	// テンプレートにTODOリストを渡してレンダリング
	if err := tmpl.Execute(w, todos); err != nil {
		http.Error(w, "テンプレートレンダリングエラー", http.StatusInternalServerError)
		log.Printf("テンプレートレンダリングエラー: %v", err)
	}
}
