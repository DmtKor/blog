package handlers

import (
	"blog/src/commander"
	"blog/src/managedb"
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// POST handler adds new post to DB and redirects to created post
// If error occurrs, redirect to err_handler with message
func NewPostHandler(w http.ResponseWriter, r *http.Request) {
	post := managedb.Post{}
	post.Content = r.FormValue("Content")
	post.Description = r.FormValue("Description")
	post.PostDate = time.Now()
	tags := r.FormValue("Tags") // like "apples,bananas"; can be empty
	tags = strings.ReplaceAll(tags, " ", "")
	post.Tags = strings.Split(tags, ",")
	post.Title = r.FormValue("Title")
	// Check for empty fields
	if post.Title == "" {
		ctx := context.WithValue(r.Context(), "errormsg", "Empty post title.")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	if post.Description == "" {
		ctx := context.WithValue(r.Context(), "errormsg", "Empty post description.")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	if post.Content == "" {
		ctx := context.WithValue(r.Context(), "errormsg", "Empty post content.")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	post, err := commander.Comm.Database.NewPost(post)
	if err != nil {
		ctx := context.WithValue(r.Context(), "errormsg", "Unable to add new post (" + err.Error() + ").")
		go ErrHandler(w, r.WithContext(ctx))
		return
	}
	// Success
	// Redirect to created post
	http.Redirect(w, r, "/doc?id=" + strconv.Itoa(int(post.Id)), http.StatusSeeOther)
}