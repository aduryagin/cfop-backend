package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/99designs/gqlgen/handler"
	cfop "github.com/aduryagin/cfop-backend"
	DB "github.com/aduryagin/cfop-backend/db"
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

	router.Handle("/api", handler.GraphQL(cfop.NewExecutableSchema(cfop.Config{Resolvers: &cfop.Resolver{}})))

	// static	files serve

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "static")
	FileServer(router, "/static", http.Dir(filesDir))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground!", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
