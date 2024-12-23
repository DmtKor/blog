package handlers

import (
	"fmt"
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
	fmt.Println("[INFO] Serving ErrHandler with msg", errmsg)
    // Serve data
	tmpl, _ := template.ParseFiles("templates/error.html")
    err := tmpl.Execute(w, errmsg)
	if err != nil {
		fmt.Println("[ERROR] Error executing template:", err)
	}
}


