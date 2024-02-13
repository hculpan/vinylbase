package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/hculpan/vinylbase/cmd/web/handlers"
	"github.com/hculpan/vinylbase/pkg/db"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("DB_NAME undefined")
	}

	tursoToken := os.Getenv("TURSO_TOKEN")
	if tursoToken == "" {
		log.Fatal("TURSO_TOKEN undefined")
	}

	if err := db.InitDb(dbName, tursoToken); err != nil {
		log.Fatal(err)
	}
	defer db.CloseDb()
	log.Default().Println("Connected to database")

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Port not configured, using default")
		port = "8080"
	}

	// Initialize router
	r := chi.NewRouter()

	handlers.InitSessionManager(r)

	setRoutes(r)

	// Start server
	log.Printf("Starting server on :%s", port)
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})
}
