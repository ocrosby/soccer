package handlers

import (
	"net/http"
	"path"
	"strings"
)

type SwaggerHandler struct {
}

type SwaggerUIHandler struct {
	StaticPath string
	IndexPath  string
}

func NewSwaggerHandler() *SwaggerHandler {
	return &SwaggerHandler{}
}

func (h *SwaggerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "swagger.yaml")
}

func NewSwaggerUIHandler(staticPath, indexPath string) *SwaggerUIHandler {
	return &SwaggerUIHandler{
		StaticPath: staticPath,
		IndexPath:  indexPath,
	}
}

func (h *SwaggerUIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, h.IndexPath)
		return
	}

	// Construct the file name based on the request URL
	filePath := path.Join(h.StaticPath, r.URL.Path)

	// Serve a file diretly if it exists and is not a directory
	if strings.HasSuffix(r.URL.Path, ".html") ||
		strings.HasSuffix(r.URL.Path, ".css") ||
		strings.HasSuffix(r.URL.Path, ".js") ||
		strings.HasSuffix(r.URL.Path, ".png") {
		http.ServeFile(w, r, filePath)
		return
	}

	// Fallback to serving the Swagger YAML file
	http.ServeFile(w, r, path.Join(h.StaticPath, "swagger.yaml"))
}
