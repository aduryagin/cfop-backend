package main

import (
	"log"
	"net/http"
	"os"
	cfop "github.com/aduryagin/cfop-backend"
	DB "github.com/aduryagin/cfop-backend/db"
	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	DB.Init()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowCredentials: true,
	}).Handler)

	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", handler.GraphQL(cfop.NewExecutableSchema(cfop.Config{Resolvers: &cfop.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
