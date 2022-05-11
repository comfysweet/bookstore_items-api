package items

import (
	"encoding/json"
	"fmt"
	"github.com/comfysweet/bookstore_items-api/clients/elastic_search"
	"github.com/comfysweet/bookstore_items-api/domain/queries"
	"github.com/comfysweet/bookstore_utils-go/errors"
	"strings"
)

const (
	indexItems = "items"
	typeItem   = "_doc"
)

func (i *Item) Save() errors.RestErr {
	result, err := elastic_search.Client.Index(indexItems, typeItem, i)
	if err != nil {
		return errors.NewInternalServiceError("error when trying to save item", errors.NewError("database error"))
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() errors.RestErr {
	itemId := i.Id
	result, err := elastic_search.Client.Get(indexItems, typeItem, i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return errors.NewNotFoundError(fmt.Sprintf("no item found with %s", i.Id))
		}
		return errors.NewInternalServiceError("error when trying to get item", errors.NewError("database error"))
	}

	bytes, errM := result.Source.MarshalJSON()
	if errM != nil {
		return errors.NewInternalServiceError("error when trying to get item", errors.NewError("database error"))
	}
	if err := json.Unmarshal(bytes, i); err != nil {
		return errors.NewInternalServiceError("error when trying to get item", errors.NewError("database error"))
	}
	i.Id = itemId
	return nil
}

func (i *Item) Search(query queries.EsQuery) ([]Item, errors.RestErr) {
	result, err := elastic_search.Client.Search(indexItems, query.Build())
	if err != nil {
		return nil, errors.NewInternalServiceError("error when trying to search item", errors.NewError("database error"))
	}
	fmt.Println(result)

	items := make([]Item, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, errors.NewInternalServiceError("error when trying to parse json", errors.NewError("database error"))
		}
		item.Id = hit.Id
		items[index] = item
	}

	if len(items) == 0 {
		return nil, errors.NewNotFoundError("no items found")
	}
	return items, nil
}
