package handlers

import (
	"github.com/victorbej/news-service/internal/adapters/handlers/checkHealthService"
	"github.com/victorbej/news-service/internal/domain/usecase/user"
	"sync"
)

func HandleServices() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go checkHealthService.RunCheckHealthServer(wg)
	go user.AuthUser(wg)

	wg.Wait()
}
