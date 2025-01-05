package controllers

import (
	"log"
	"net/http"
	"server/db"
	"server/models"
	"strconv"
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

func AddTodo(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("todo")
	if name == "" {
		http.Error(w, "TODO名が空です", http.StatusBadRequest)
		return
	}
	_, err := db.DB.Exec("INSERT INTO Todos (name) VALUES ($1)", name)
	if err != nil {
		http.Error(w, "データベースエラー", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "IDが空です", http.StatusBadRequest)
		return
	}
	// IDを整数に変換
	intID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "無効なID形式", http.StatusBadRequest)
		log.Printf("ID変換エラー: %v", err)
		return
	}

	// データベースから削除
	_, err = db.DB.Exec("DELETE FROM todos WHERE id = $1", intID)
	if err != nil {
		http.Error(w, "データベースエラー", http.StatusInternalServerError)
		log.Printf("削除エラー: %v", err)
		return
	}

	// 成功したらリストページにリダイレクト
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
