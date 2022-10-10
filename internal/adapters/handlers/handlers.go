package handlers

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/victorbej/news-service/internal/adapters/handlers/checkService"
	"github.com/victorbej/news-service/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"sync"
)

func startServer(wg *sync.WaitGroup) {
	defer wg.Done()

	grpcServer := grpc.NewServer()
	server := &checkService.ServiceServer{}

	api.RegisterContentCheckServiceServer(grpcServer, server)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := api.RegisterContentCheckServiceHandlerFromEndpoint(context.Background(), mux, "localhost:9090", opts)
	if err != nil {
		log.Fatal(err)
	}

	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func HandleServices() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go startServer(wg)
	//go checkService.HandleServiceClient(wg)

	wg.Wait()
}
