package middleware

import (
	"blog/src/commander"
	"blog/src/handlers"
	"context"
	"fmt"
	"net/http"
)

// logging info about requests
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	  fmt.Println("[INFO] Serving request", r.Method, r.URL)
	  defer fmt.Println("[INFO] Finished serving request", r.Method, r.URL)
	  next.ServeHTTP(w, r)
	})
}

// Check if server is about to shutdown (if so, redirects every request to handlers/err_handler)
// Else increase waitgroup in commander
func RedirectMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !commander.Comm.ShouldHandle {
			ctx := context.WithValue(r.Context(), "errormsg", "Server is shutting down.")
			go handlers.ErrHandler(w, r.WithContext(ctx))
			return
		}
		commander.Comm.HandlersWG.Add(1)
		defer commander.Comm.HandlersWG.Done()
		next.ServeHTTP(w, r)
	})
}