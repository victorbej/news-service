package newsService

import (
	"context"
	"github.com/victorbej/news-service/pkg/api"
)

type NewsService struct {
	api.UnimplementedNewsServiceServer
}

func (n *NewsService) GetNews(ctx context.Context, in *api.NewsRequestParams) (*api.NewsList, error) {
	return &api.NewsList{News: nil, Total: 1}, nil
}

func ConvertEntityNewsToNewsObject(news *api.NewsRequestParams) *api.NewsObject {
	return &api.NewsObject{
		Id:          "",
		Title:       "",
		Author:      "",
		Active:      false,
		ActiveFrom:  0,
		Text:        "",
		TextJson:    "",
		UserId:      "",
		Tags:        nil,
		FilesInfo:   nil,
		IsImportant: false,
	}
}
