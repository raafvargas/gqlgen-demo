package marvel

import (
	"context"
	"encoding/json"
)

type MarvelCharacterClient struct {
	*MarvelClient
}

func NewMarvelCharacterClient() *MarvelCharacterClient {
	return &MarvelCharacterClient{
		MarvelClient: NewMarvelClient("characters"),
	}
}

func (c *MarvelCharacterClient) Get(ctx context.Context, id int) (*Character, error) {
	res, err := c.MarvelClient.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	characters := []*Character{}

	if err := json.Unmarshal(res, &characters); err != nil {
		return nil, err
	}

	return characters[0], nil
}

func (c *MarvelCharacterClient) Search(ctx context.Context, keyValues ...string) ([]*Character, error) {
	res, err := c.MarvelClient.Search(ctx, keyValues...)

	if err != nil {
		return nil, err
	}

	characters := []*Character{}

	if err := json.Unmarshal(res, &characters); err != nil {
		return nil, err
	}

	return characters, nil
}
