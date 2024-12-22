package handlers

import "net/http"

// Message will be passed with context value... I guess
func ErrHandler(w http.ResponseWriter, r *http.Request) {
	
}
// Does not serve any path
// If error occurrs in any handler, they will call this handler, 
// that should return error.html template with message

// This handler does not do anything with waitgroup


