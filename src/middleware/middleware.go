package middleware

import (
	"fmt"
	"net/http"
	"blog/src/commander"
)

// logging info about requests
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	  fmt.Println("[INFO] Serving request", r.Method, r.URL)
	  next.ServeHTTP(w, r)
	})
}

// Check if server is about to shutdown (if so, redirects every request to handlers/err_handler)
// Else increase waitgroup in commander
func RedirectMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !commander.Comm.ShouldHandle {
			// Serve err_handler
		}
		commander.Comm.HandlersWG.Add(1)
		defer commander.Comm.HandlersWG.Done()
		next.ServeHTTP(w, r)
	})
}