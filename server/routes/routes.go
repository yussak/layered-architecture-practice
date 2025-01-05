package routes

import (
	"net/http"
	"server/controllers"
)

func SetupRoutes() {
	http.HandleFunc("/", controllers.ListTodos)
	http.HandleFunc("/add", controllers.AddTodo)
	http.HandleFunc("/delete", controllers.DeleteTodo)
}
