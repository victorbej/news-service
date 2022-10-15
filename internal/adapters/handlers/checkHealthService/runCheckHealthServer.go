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

func RunCheckHealthServer(wg *sync.WaitGroup) {
	defer wg.Done()

	grpcServer := grpc.NewServer()
	server := &ContentCheckService{}
	listen, err := net.Listen(cfg.TCP, cfg.GRPC)
	if err != nil {
		log.Fatal(err)
	}

	api.RegisterContentCheckServiceServer(grpcServer, server)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err = api.RegisterContentCheckServiceHandlerFromEndpoint(context.Background(), mux, cfg.GRPC, opts)
	if err != nil {
		log.Fatal(err)
	}

	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() (err error) {
		return grpcServer.Serve(listen)
	})
	g.Go(func() (err error) {
		return http.ListenAndServe(cfg.Port, mux)
	})

	err = g.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
