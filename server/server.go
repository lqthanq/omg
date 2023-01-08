package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/lqthanq/omg-simpler/app"
	"github.com/lqthanq/omg-simpler/graph"
	"github.com/lqthanq/omg-simpler/migrate"
)

const defaultPort = "9000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	app.InitDatabase()
	migrate.Migrate()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
