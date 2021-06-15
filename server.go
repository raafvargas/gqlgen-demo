package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/raafvargas/gqlgen-demo/graph"
	"github.com/raafvargas/gqlgen-demo/graph/generated"
	"github.com/raafvargas/gqlgen-demo/marvel"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		ComicClient:     marvel.NewMarvelComicClient(),
		StoryClient:     marvel.NewMarvelStoryClient(),
		CharacterClient: marvel.NewMarvelCharacterClient(),
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
