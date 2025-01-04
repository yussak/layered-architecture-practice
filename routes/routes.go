package routes

import (
	"net/http"
	"server/controllers"
)

func SetupRoutes() {
	http.HandleFunc("/", controllers.ListTodos)
}
