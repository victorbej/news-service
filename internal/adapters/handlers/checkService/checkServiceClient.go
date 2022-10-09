package checkService

import (
	"context"
	"fmt"
	api "github.com/victorbej/news-service/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"sync"
)

func HandleServiceClient(wg *sync.WaitGroup) {
	defer wg.Done()

	conn, err := grpc.Dial(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := api.NewContentCheckServiceClient(conn)
	response, err := client.CheckHealth(context.Background(), &api.EmptyRequest{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response.GetServiceName(), response.GetServiceStatus())
}
