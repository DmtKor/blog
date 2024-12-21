package handlers

// DELETE handler deletes post and redirects to /browse?page=1&pagesize=10
// If error occurrs, redirect to err_handler with message
// Path: /doc?id=... (id not optional) - just like seepost_handler, but with DELETE method

// All handlers here add 1 to waitgroup and do .Done() in the end