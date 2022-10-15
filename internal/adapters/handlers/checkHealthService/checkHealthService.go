package checkHealthService

import (
	"context"
	"github.com/victorbej/news-service/pkg/api"
	"github.com/victorbej/news-service/pkg/config"
)

var cfg = config.NewConfig()

type ContentCheckService struct {
	api.UnimplementedContentCheckServiceServer
}

func (s *ContentCheckService) CheckHealth(ctx context.Context, req *api.EmptyRequest) (res *api.HealthResponse, err error) {
	return &api.HealthResponse{ServiceName: cfg.ServiceName, ServiceStatus: cfg.ServiceStatus}, nil
}
