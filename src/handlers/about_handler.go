package handlers

import (
	"net/http"
)

// GET: serves on both / and /about paths, returns simple HTML static page, not even a template
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/about.html")
}