package marvel

import (
	"context"
	"encoding/json"
)

type MarvelClient struct {
	path string
}

func NewMarvelClient(path string) *MarvelClient {
	return &MarvelClient{
		path: path,
	}
}

func (c *MarvelClient) Get(ctx context.Context, id int) (json.RawMessage, error) {
	res, err := doGet(ctx, c.path, id)

	if err != nil {
		return nil, err
	}

	return res.Data.Results, nil
}

func (c *MarvelClient) Search(ctx context.Context, keyValues ...string) (json.RawMessage, error) {
	res, err := doSearch(ctx, c.path, keyValues...)

	if err != nil {
		return nil, err
	}

	return res.Data.Results, nil
}
