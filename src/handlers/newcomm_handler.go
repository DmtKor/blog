package handlers

import "net/http"

func NewCommHandler(w http.ResponseWriter, r *http.Request) {
	
}
// POST handler /comment?postid=... (postid not optional) - creates new comment, 
// redirects to updated /doc?id=postid page
// Errors -> err_handler

// All handlers here add 1 to waitgroup and do .Done() in the end