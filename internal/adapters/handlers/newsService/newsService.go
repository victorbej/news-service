package newsService

import (
	"github.com/victorbej/news-service/pkg/api"
)

type NewsService struct {
	api.UnimplementedNewsServiceServer
}

//func (n *NewsService) GetNews(ctx context.Context, in *api.NewsRequestParams) (*api.NewsList, error) {
//
//}
