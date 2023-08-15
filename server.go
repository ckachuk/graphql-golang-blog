package main

import (
	"blog-graphql/db"
	"blog-graphql/db/migrations"
	"blog-graphql/directives"
	"blog-graphql/env"
	"blog-graphql/graph"
	"blog-graphql/middlewares"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
)

const defaultPort = "8080"

func main() {
	env.InitEnv()
	db.DBConnection()
	migrations.Migrate()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router:= mux.NewRouter()
	router.Use(middlewares.AuthMiddleware)
	
	c := graph.Config{Resolvers: &graph.Resolver{}}
	c.Directives.Auth = directives.Auth
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
