package handlers

import (
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	
}

// GET: serves on both / and /about paths, returns simple HTML static page, not even a template

// All handlers here add 1 to waitgroup and do .Done() in the end