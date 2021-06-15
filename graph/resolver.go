package graph

import "github.com/raafvargas/gqlgen-demo/marvel"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ComicClient     *marvel.MarvelComicClient
	StoryClient     *marvel.MarvelStoryClient
	CharacterClient *marvel.MarvelCharacterClient
}
