package handlers

import (
	"blog/src/commander"
	"context"
	"net/http"
	"strconv"
)

// Path: /doc?id=... (id not optional) - just like seepost_handler, but with DELETE method
// DELETE handler deletes post and redirects to /browse?page=1&pagesize=10
// If error occurrs, redirect to err_handler with message
func DelPostHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		// Error reading commid parameter
		ctx := context.WithValue(r.Context(), "errormsg", "Unable to read post ID (" + err.Error() + ").")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	exists, err := commander.Comm.Database.PostExists(uint64(id))
	if err != nil {
		// Error reading commid parameter
		ctx := context.WithValue(r.Context(), "errormsg", "Unable to check if post exists (" + err.Error() + ").")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	if !exists {
		ctx := context.WithValue(r.Context(), "errormsg", "Post you're trying too delete do not exist.")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	err = commander.Comm.Database.DeletePostById(uint64(id))
	if err != nil {
		ctx := context.WithValue(r.Context(), "errormsg", "Unable to delete post (" + err.Error() + ").")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	// Post deleted successfully
	// Redirect to browse
	http.Redirect(w, r, "/browse?page=1&pagesize=10", http.StatusSeeOther)
}