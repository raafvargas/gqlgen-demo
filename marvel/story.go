package marvel

import (
	"context"
	"encoding/json"
)

type MarvelStoryClient struct {
	*MarvelClient
}

func NewMarvelStoryClient() *MarvelStoryClient {
	return &MarvelStoryClient{
		MarvelClient: NewMarvelClient("stories"),
	}
}

func (c *MarvelStoryClient) Get(ctx context.Context, id int) (*Story, error) {
	res, err := c.MarvelClient.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	stories := []*Story{}

	if err := json.Unmarshal(res, &stories); err != nil {
		return nil, err
	}

	return stories[0], nil
}

func (c *MarvelStoryClient) Search(ctx context.Context, keyValues ...string) ([]*Story, error) {
	res, err := c.MarvelClient.Search(ctx, keyValues...)

	if err != nil {
		return nil, err
	}

	stories := []*Story{}

	if err := json.Unmarshal(res, &stories); err != nil {
		return nil, err
	}

	return stories, nil
}
