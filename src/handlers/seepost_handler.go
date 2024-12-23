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

type PostView struct {
	Post managedb.Post
	Comments []managedb.Comment
}

// GET handler returns HTML page with post, tags and comments (also DELETE button)
// If error occurrs, redirect to err_handler with message
// Path: /doc?id=... (id is not optional)
func SeePostHandler(w http.ResponseWriter, r *http.Request) {
	view := PostView{}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		ctx := context.WithValue(r.Context(), "errormsg", "Unable to read post ID (" + err.Error() + ").")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	exists, err := commander.Comm.Database.PostExists(uint64(id))
	if err != nil {
		ctx := context.WithValue(r.Context(), "errormsg", "Unable to check if post exists (" + err.Error() + ").")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	if !exists {
		ctx := context.WithValue(r.Context(), "errormsg", "Post you're trying to see do not exist.")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	view.Post, err = commander.Comm.Database.GetPostById(uint64(id))
	if err != nil {
		ctx := context.WithValue(r.Context(), "errormsg", "Unable to access post in DB (" + err.Error() + ").")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	view.Comments, err = commander.Comm.Database.GetCommentsByPostId(view.Post.Id)
	if err != nil {
		ctx := context.WithValue(r.Context(), "errormsg", "Unable to access post comments in DB (" + err.Error() + ").")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	// All data recieved
	// Serve data
	tmpl, _ := template.ParseFiles("templates/post.html")
    err = tmpl.Execute(w, view)
	if err != nil {
		fmt.Println("[ERROR] Error executing template:", err)
	}
}