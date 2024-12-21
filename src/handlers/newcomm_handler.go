package handlers

// POST handler /comment?postid=... (postid not optional) - creates new comment, 
// redirects to updated /doc?id=postid page
// Errors -> err_handler

// All handlers here add 1 to waitgroup and do .Done() in the end