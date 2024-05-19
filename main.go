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
	mux.HandleFunc("GET /{$}", handleIndex(templates))
	mux.HandleFunc("GET /health", handleHealth())

	var handler http.Handler
	handler = logging(logger, mux)

	log.Fatal(http.ListenAndServe(":8000", handler))
}

func handleIndex(templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := []Peak{
			{
				Name:          "Scafell Pike",
				HeightMetres:  978.07,
				GridReference: "NY215072",
			},
			{
				Name:          "Scafell",
				HeightMetres:  963.9,
				GridReference: "NY206064",
			},
		}

		templates.Execute(w, data)
	}
}

func handleHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "200 ok")
	}
}
