package checkHealthService

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/victorbej/news-service/pkg/api"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"sync"
)

const (
	grpcHostPort = "0.0.0.0:9090"
)

func RunCheckHealthService(wg *sync.WaitGroup) {
	defer wg.Done()

	grpcServer := grpc.NewServer()
	server := &ServiceServer{}
	listen, err := net.Listen("tcp", grpcHostPort)
	if err != nil {
		log.Fatal(err)
	}

	api.RegisterContentCheckServiceServer(grpcServer, server)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err = api.RegisterContentCheckServiceHandlerFromEndpoint(context.Background(), mux, grpcHostPort, opts)
	if err != nil {
		log.Fatal(err)
	}

	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() (err error) {
		return grpcServer.Serve(listen)
	})
	g.Go(func() (err error) {
		return http.ListenAndServe(":8081", mux)
	})

	err = g.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
