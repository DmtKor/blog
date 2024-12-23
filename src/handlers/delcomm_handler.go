package handlers

import (
	"blog/src/commander"
	"context"
	"net/http"
	"strconv"
)

// DELETE /comment?postid=...&commid=... (commid not optional), deletes comment, is called by link/button on page
func DelCommHandler(w http.ResponseWriter, r *http.Request) {
	commid, err := strconv.Atoi(r.URL.Query().Get("commid"))
	if err != nil {
		// Error reading commid parameter
		ctx := context.WithValue(r.Context(), "errormsg", "Unable to read comment ID (" + err.Error() + ").")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	exists, err := commander.Comm.Database.CommExists(uint64(commid))
	if err != nil {
		ctx := context.WithValue(r.Context(), "errormsg", "Unable to check if comment exists (" + err.Error() + ").")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	if !exists {
		ctx := context.WithValue(r.Context(), "errormsg", "Comment you're trying to delete do not exist!")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	err = commander.Comm.Database.DeleteCommentById(uint64(commid))
	if err != nil {
		// Error deleting comment
		ctx := context.WithValue(r.Context(), "errormsg", "Unable to delete comment (" + err.Error() + ").")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	// If comment was deleted successfully 
	postid, err := strconv.Atoi(r.URL.Query().Get("postid"))
	exists, err = commander.Comm.Database.PostExists(uint64(postid))
	if err != nil || !exists {
		// Redirect to browse
		http.Redirect(w, r, "/browse?page=1&pagesize=10", http.StatusSeeOther)
	} else {
		// Redirect back to post
		http.Redirect(w, r, "/doc?id=" + strconv.Itoa(postid), http.StatusSeeOther)
	}
}