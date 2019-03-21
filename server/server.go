package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	cfop "github.com/aduryagin/cfop/backend"
	DB "github.com/aduryagin/cfop/backend/db"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const defaultPort = "8080"

func main() {
	DB.Init()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(cfop.NewExecutableSchema(cfop.Config{Resolvers: &cfop.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
