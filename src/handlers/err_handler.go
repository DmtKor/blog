package handlers

import (
	"html/template"
	"net/http"
)

// Does not serve any path
// If error occurrs in any handler, they will call this handler, 
// that should return error.html template with message
// Message will be passed with context value.
// This handler does not do anything with waitgroup
func ErrHandler(w http.ResponseWriter, r *http.Request) {
	errmsg := r.Context().Value("errormsg").(string)
    // Serve data
	tmpl, _ := template.ParseFiles("templates/error.html")
    tmpl.Execute(w, errmsg)
}


