package main

import (
	"fmt"
	"net/http"
	"server/db"
	"server/routes"

	_ "github.com/lib/pq"
)

func main() {
	db.Init()
	routes.SetupRoutes()

	fmt.Println("Server running on port :8080")
	http.ListenAndServe(":8080", nil)
}
