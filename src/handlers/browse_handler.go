package handlers

import (
	"blog/src/commander"
	"blog/src/managedb"
	"context"
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type browseView struct {
	page int
	pagesize int
	maxpages uint64
	posts []managedb.Post
}

// GET: serves /browse[?page=...&pagesize=...] (default values - page: 1, pagesize: 10)
// Returns page of posts ordered by time with links to /doc?id=... (seepost_handler)
func BrowseHandler(w http.ResponseWriter, r *http.Request) {
	view := browseView{}
	var err error
	view.page, err = strconv.Atoi(chi.URLParam(r, "page"))
	if err != nil {
		// Error reading page parameter
		view.page = 1
	}
	view.pagesize, err = strconv.Atoi(chi.URLParam(r, "pagesize"))
	if err != nil {
		// Error reading pagesize parameter
		view.pagesize = 10
	}
	view.maxpages, err = commander.Comm.Database.GetPageNum(uint64(view.pagesize))
	if err != nil {
		// Error reading number of pages
		ctx := context.WithValue(r.Context(), "errormsg", "Unable to read number of pages (" + err.Error() + ").")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	view.posts, err = commander.Comm.Database.GetPostsByPage(uint(view.page), uint(view.pagesize))
	if err != nil {
		// Error reading posts on page
		ctx := context.WithValue(r.Context(), "errormsg", "Unable to read posts from DB (" + err.Error() + ").")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	if len(view.posts) == 0 {
		// Empty page
		ctx := context.WithValue(r.Context(), "errormsg", "Page you're trying to access is empty.")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	// Serve data
	tmpl, _ := template.ParseFiles("templates/browse.html")
    tmpl.Execute(w, view)
}