package handlers

import (
	"net/http"
)

func DelCommHandlerGET(w http.ResponseWriter, r *http.Request) {
	
}

func DelCommHandlerDELETE(w http.ResponseWriter, r *http.Request) {
	
}
// DELETE /comment?commid=... (commid not optional), deletes comment, is called by link/button on 
// GET /doc?id=... page
// Errors -> err_handler.go

// All handlers here add 1 to waitgroup and do .Done() in the end