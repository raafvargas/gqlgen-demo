package marvel

import (
	"context"
	"encoding/json"
)

type MarvelComicClient struct {
	*MarvelClient
}

func NewMarvelComicClient() *MarvelComicClient {
	return &MarvelComicClient{
		MarvelClient: NewMarvelClient("comics"),
	}
}

func (c *MarvelComicClient) Get(ctx context.Context, id int) (*Comic, error) {
	res, err := c.MarvelClient.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	comics := []*Comic{}

	if err := json.Unmarshal(res, &comics); err != nil {
		return nil, err
	}

	return comics[0], nil
}

func (c *MarvelComicClient) Search(ctx context.Context, keyValues ...string) ([]*Comic, error) {
	res, err := c.MarvelClient.Search(ctx, keyValues...)

	if err != nil {
		return nil, err
	}

	comics := []*Comic{}

	if err := json.Unmarshal(res, &comics); err != nil {
		return nil, err
	}

	return comics, nil
}
