package handlers

// GET handler returns HTML page with post, tags and comments (also DELETE button)
// If error occurrs, redirect to err_handler with message
// Path: /doc?id=... (id is not optional)

// All handlers here add 1 to waitgroup and do .Done() in the end