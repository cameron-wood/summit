package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"os"
)

var (
	//go:embed all:templates
	templateFS embed.FS

	//go:embed all:static
	staticFS embed.FS
)

func main() {
	templates, err := template.ParseFS(templateFS, "templates/*.html")
	if err != nil {
		log.Fatalf("failed to parse templates: %v", err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(staticFS))
	mux.HandleFunc("GET /{$}", handleRoot(templates))
	mux.HandleFunc("GET /health", handleHealth())

	var handler http.Handler
	handler = logging(logger, mux)

	log.Fatal(http.ListenAndServe(":8000", handler))
}

func handleRoot(templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templates.Execute(w, nil)
	}
}

func handleHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "200 ok")
	}
}
