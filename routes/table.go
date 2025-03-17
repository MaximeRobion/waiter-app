package routes

import(
	"net/http"
	"waiter-app/handlers"
)

func RegisterRoutes(){
	http.HandleFunc("POST /tables", handlers.CreateTable)
	http.HandleFunc("GET /tables", handlers.GetAllTables)
	http.HandleFunc("GET /tables/", handlers.GetTable)
	http.HandleFunc("PUT /tables/", handlers.UpdateTable)
	http.HandleFunc("DELETE /tables/", handlers.DeleteTable)
}