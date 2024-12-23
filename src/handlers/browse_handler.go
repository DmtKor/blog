package handlers

import (
	"blog/src/commander"
	"blog/src/managedb"
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type BrowseView struct {
	Page int
	Pagesize int
	Maxpages uint64
	Posts []managedb.Post
}

// GET: serves /browse[?page=...&pagesize=...] (default values - page: 1, pagesize: 10)
// Returns page of posts ordered by time with links to /doc?id=... (seepost_handler)
func BrowseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[DEBUG] Started BrowseHandler")
	view := BrowseView{}
	var err error
	view.Page, err = strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		// Error reading page parameter
		view.Page = 1
	}
	view.Pagesize, err = strconv.Atoi(r.URL.Query().Get("pagesize"))
	if err != nil {
		// Error reading pagesize parameter
		view.Pagesize = 10
	}
	view.Maxpages, err = commander.Comm.Database.GetPageNum(uint64(view.Pagesize))
	if err != nil {
		// Error reading number of pages
		ctx := context.WithValue(r.Context(), "errormsg", "Unable to read number of pages (" + err.Error() + ").")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	view.Posts, err = commander.Comm.Database.GetPostsByPage(uint(view.Page), uint(view.Pagesize))
	if err != nil {
		// Error reading posts on page
		ctx := context.WithValue(r.Context(), "errormsg", "Unable to read posts from DB (" + err.Error() + ").")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	if len(view.Posts) == 0 {
		// Empty page
		ctx := context.WithValue(r.Context(), "errormsg", "Page you're trying to access is empty.")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	// Serve data
	tmpl, err := template.ParseFiles("templates/browse.html")
	if err != nil {
		ctx := context.WithValue(r.Context(), "errormsg", "Unable to parse template.")
		fmt.Println("[ERROR] No template")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
    err = tmpl.Execute(w, view)
	fmt.Println(view)
	if err != nil {
		fmt.Println("[ERROR] Error executing template:", err)
	}
}