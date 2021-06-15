package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/raafvargas/gqlgen-demo/graph/generated"
	"github.com/raafvargas/gqlgen-demo/graph/model"
)

func (r *characterResolver) Comics(ctx context.Context, obj *model.Character, offset int, limit int) ([]*model.Comic, error) {
	res, err := r.ComicClient.Search(ctx,
		"characters", strconv.Itoa(obj.ID), "offset", strconv.Itoa(offset), "limit", strconv.Itoa(limit))

	if err != nil {
		return nil, err
	}

	comics := []*model.Comic{}

	for _, c := range res {
		comics = append(comics, model.MapComic(c))
	}

	return comics, nil
}

func (r *characterResolver) Stories(ctx context.Context, obj *model.Character, offset int, limit int) ([]*model.Story, error) {
	res, err := r.StoryClient.Search(ctx,
		"characters", strconv.Itoa(obj.ID), "offset", strconv.Itoa(offset), "limit", strconv.Itoa(limit))

	if err != nil {
		return nil, err
	}

	stories := []*model.Story{}

	for _, s := range res {
		stories = append(stories, model.MapStory(s))
	}

	return stories, nil
}

func (r *comicResolver) Stories(ctx context.Context, obj *model.Comic, offset int, limit int) ([]*model.Story, error) {
	res, err := r.StoryClient.Search(ctx,
		"comics", strconv.Itoa(obj.ID), "offset", strconv.Itoa(offset), "limit", strconv.Itoa(limit))

	if err != nil {
		return nil, err
	}

	stories := []*model.Story{}

	for _, s := range res {
		stories = append(stories, model.MapStory(s))
	}

	return stories, nil
}

func (r *comicResolver) Characters(ctx context.Context, obj *model.Comic, offset int, limit int) ([]*model.Character, error) {
	res, err := r.CharacterClient.Search(ctx,
		"stories", strconv.Itoa(obj.ID), "offset", strconv.Itoa(offset), "limit", strconv.Itoa(limit))

	if err != nil {
		return nil, err
	}

	characters := []*model.Character{}

	for _, c := range res {
		characters = append(characters, model.MapCharacter(c))
	}

	return characters, nil
}

func (r *queryResolver) Character(ctx context.Context, id int) (*model.Character, error) {
	character, err := r.CharacterClient.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	return model.MapCharacter(character), nil
}

func (r *queryResolver) Characters(ctx context.Context, offset int, limit int) ([]*model.Character, error) {
	res, err := r.CharacterClient.Search(ctx, "offset", strconv.Itoa(offset), "limit", strconv.Itoa(limit))

	if err != nil {
		return nil, err
	}

	characters := []*model.Character{}

	for _, c := range res {
		characters = append(characters, model.MapCharacter(c))
	}

	return characters, nil
}

func (r *queryResolver) Story(ctx context.Context, id int) (*model.Story, error) {
	story, err := r.StoryClient.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	return model.MapStory(story), nil
}

func (r *queryResolver) Stories(ctx context.Context, offset int, limit int) ([]*model.Story, error) {
	res, err := r.StoryClient.Search(ctx, "offset", strconv.Itoa(offset), "limit", strconv.Itoa(limit))

	if err != nil {
		return nil, err
	}

	stories := []*model.Story{}

	for _, s := range res {
		stories = append(stories, model.MapStory(s))
	}

	return stories, nil
}

func (r *queryResolver) Comic(ctx context.Context, id int) (*model.Comic, error) {
	res, err := r.ComicClient.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	return model.MapComic(res), nil
}

func (r *queryResolver) Comics(ctx context.Context, offset int, limit int) ([]*model.Comic, error) {
	res, err := r.ComicClient.Search(ctx, "offset", strconv.Itoa(offset), "limit", strconv.Itoa(limit))

	if err != nil {
		return nil, err
	}

	comics := []*model.Comic{}

	for _, c := range res {
		comics = append(comics, model.MapComic(c))
	}

	return comics, nil
}

func (r *storyResolver) Comics(ctx context.Context, obj *model.Story, offset int, limit int) ([]*model.Comic, error) {
	res, err := r.ComicClient.Search(ctx,
		"characters", strconv.Itoa(obj.ID), "offset", strconv.Itoa(offset), "limit", strconv.Itoa(limit))

	if err != nil {
		return nil, err
	}

	comics := []*model.Comic{}

	for _, c := range res {
		comics = append(comics, model.MapComic(c))
	}

	return comics, nil
}

func (r *storyResolver) Characters(ctx context.Context, obj *model.Story, offset int, limit int) ([]*model.Character, error) {
	res, err := r.CharacterClient.Search(ctx,
		"stories", strconv.Itoa(obj.ID), "offset", strconv.Itoa(offset), "limit", strconv.Itoa(limit))

	if err != nil {
		return nil, err
	}

	characters := []*model.Character{}

	for _, c := range res {
		characters = append(characters, model.MapCharacter(c))
	}

	return characters, nil
}

// Character returns generated.CharacterResolver implementation.
func (r *Resolver) Character() generated.CharacterResolver { return &characterResolver{r} }

// Comic returns generated.ComicResolver implementation.
func (r *Resolver) Comic() generated.ComicResolver { return &comicResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Story returns generated.StoryResolver implementation.
func (r *Resolver) Story() generated.StoryResolver { return &storyResolver{r} }

type characterResolver struct{ *Resolver }
type comicResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type storyResolver struct{ *Resolver }
