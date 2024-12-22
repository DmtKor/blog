package handlers

import "net/http"

func NewPostHandlerGET(w http.ResponseWriter, r *http.Request) {
	
}

func NewPostHandlerPOST(w http.ResponseWriter, r *http.Request) {
	
}

// GET handler returns HTML page with POST form (/newdoc)

// POST handler adds new post to DB and redirects to created post
// If error occurrs, redirect to err_handler with message

// All handlers here add 1 to waitgroup and do .Done() in the end