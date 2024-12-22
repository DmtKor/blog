package handlers

import (
	"blog/src/commander"
	"blog/src/managedb"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

// POST handler /comment?postid=... (postid not optional) - creates new comment,
// redirects to updated /doc?id=postid page
// Errors -> err_handler
func NewCommHandler(w http.ResponseWriter, r *http.Request) {
	postid, err := strconv.Atoi(chi.URLParam(r, "postid"))
	if err != nil {
		ctx := context.WithValue(r.Context(), "errormsg", "Unable to read post ID (" + err.Error() + ").")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	exists, err := commander.Comm.Database.PostExists(uint64(postid))
	if err != nil {
		ctx := context.WithValue(r.Context(), "errormsg", "Unable to check if post exists (" + err.Error() + ").")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	if !exists {
		ctx := context.WithValue(r.Context(), "errormsg", "Post you're trying to delete do not exist.")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	// Reading form data
	comm := managedb.Comment{} 
	comm.Author = r.FormValue("Author")  // Can be empty
	comm.Content = r.FormValue("Content")
	comm.Email = r.FormValue("Email")
	comm.CommDate = time.Now()
	comm.PostId = uint64(postid)
	if comm.Content == "" {
		ctx := context.WithValue(r.Context(), "errormsg", "Empty comment content.")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	if comm.Email == "" {
		ctx := context.WithValue(r.Context(), "errormsg", "Empty comment email.")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	// Trying to add new comment
	err = commander.Comm.Database.NewComment(comm)
	if err != nil {
		ctx := context.WithValue(r.Context(), "errormsg", "Unable to add comment into DB (" + err.Error() + ").")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	// Success
	http.Redirect(w, r, "/doc?id=" + strconv.Itoa(postid), http.StatusSeeOther)
}