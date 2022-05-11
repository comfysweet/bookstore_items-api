package elastic_search

import (
	"context"
	"fmt"
	"github.com/comfysweet/bookstore_utils-go/logger"
	"github.com/olivere/elastic"
	"time"
)

var (
	Client esClientInterface = &esClient{}
)

type esClient struct {
	client *elastic.Client
}

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string, string) (*elastic.GetResult, error)
	Search(string, elastic.Query) (*elastic.SearchResult, error)
}

func Init() {
	log := logger.GetLogger()
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
	)
	if err != nil {
		panic(err)
	}
	Client.setClient(client)
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(index string, docType string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().
		Index(index).
		Type(docType).
		BodyJson(doc).
		Do(ctx)
	if err != nil {
		logger.Err(fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Get(index string, docType string, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := c.client.Get().
		Index(index).
		Type(docType).
		Id(id).
		Do(ctx)
	if err != nil {
		logger.Err(fmt.Sprintf("error when trying to get id %s", id), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Search(index string, query elastic.Query) (*elastic.SearchResult, error) {
	ctx := context.Background()
	if err := c.client.Search(index).Query(query).Validate(); err != nil {
		fmt.Println("result" + err.Error())
		return nil, err
	}
	result, err := c.client.Search(index).Query(query).RestTotalHitsAsInt(true).Do(ctx)
	if err != nil {
		logger.Err(fmt.Sprintf("error when trying to search documents in index %s", index), err)
		return nil, err
	}
	return result, nil
}
