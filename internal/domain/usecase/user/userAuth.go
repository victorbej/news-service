package user

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/victorbej/news-service/internal/domain/entity/app"
	"github.com/victorbej/news-service/internal/domain/entity/controllers"
	"github.com/victorbej/news-service/pkg/config"
	"net/http"
	"sync"
)

func AuthUser(wg *sync.WaitGroup) {
	defer wg.Done()

	var cfg = config.NewConfig()

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	err := http.ListenAndServe(cfg.Port, router)
	if err != nil {
		fmt.Print(err)
	}
}
