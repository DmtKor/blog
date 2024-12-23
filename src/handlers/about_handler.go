package handlers

import (
	"fmt"
	"net/http"
)

// GET: serves on both / and /about paths, returns simple HTML static page, not even a template
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[DEBUG] Started AboutHandler")
	http.ServeFile(w, r, "static/about.html")
}