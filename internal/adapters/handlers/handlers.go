package handlers

import (
	"github.com/victorbej/news-service/internal/adapters/handlers/checkHealthService"
	"sync"
)

func HandleServices() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go checkHealthService.RunCheckHealthService(wg)

	wg.Wait()
}
