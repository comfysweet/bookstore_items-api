package services

import (
	"github.com/comfysweet/bookstore_items-api/domain/items"
	"github.com/comfysweet/bookstore_items-api/domain/queries"
	"github.com/comfysweet/bookstore_utils-go/errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, errors.RestErr)
	Get(string) (*items.Item, errors.RestErr)
	Search(queries.EsQuery) ([]items.Item, errors.RestErr)
}

type itemsService struct {
}

func (s *itemsService) Create(item items.Item) (*items.Item, errors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Get(id string) (*items.Item, errors.RestErr) {
	item := items.Item{Id: id}
	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Search(query queries.EsQuery) ([]items.Item, errors.RestErr) {
	dao := items.Item{}
	return dao.Search(query)
}
