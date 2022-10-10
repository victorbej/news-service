package checkHealthService

import (
	"context"
	api "github.com/victorbej/news-service/pkg/api"
)

type ContentCheckService struct {
	api.UnimplementedContentCheckServiceServer
}

func (s *ContentCheckService) CheckHealth(ctx context.Context, req *api.EmptyRequest) (res *api.HealthResponse, err error) {
	return &api.HealthResponse{ServiceName: "CheckHealthServer", ServiceStatus: "200"}, nil
}
