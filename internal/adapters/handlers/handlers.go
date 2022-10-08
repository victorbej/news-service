package handlers

import (
	"github.com/victorbej/news-service/internal/adapters/handlers/checkService"
	api "github.com/victorbej/news-service/pkg/api"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

func startServer(wg *sync.WaitGroup) {
	defer wg.Done()

	grpcServer := grpc.NewServer()
	server := &checkService.ServiceServer{}
	api.RegisterContentCheckServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func HandleServices() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go startServer(wg)
	go checkService.HandleServiceClient(wg)

	wg.Wait()
}
