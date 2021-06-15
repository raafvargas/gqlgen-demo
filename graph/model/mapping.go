package model

import "github.com/raafvargas/gqlgen-demo/marvel"

func MapCharacter(c *marvel.Character) *Character {
	return &Character{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
	}
}

func MapStory(s *marvel.Story) *Story {
	return &Story{
		ID:          s.ID,
		Title:       s.Title,
		Description: s.Description,
	}
}

func MapComic(s *marvel.Comic) *Comic {
	return &Comic{
		ID:          s.ID,
		Title:       s.Title,
		Description: s.VariantDescription,
	}
}
