package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/victorbej/news-service/internal/domain/entity/app"
	"github.com/victorbej/news-service/internal/domain/entity/controllers"
	"net/http"
)

func main() {
	//handlers.HandleServices()

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	err := http.ListenAndServe(":8080", router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
