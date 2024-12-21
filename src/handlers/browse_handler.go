package handlers

// GET: serves /browse[?page=...&pagesize=...] (default values - page: 1, pagesize: 10)
// Returns page of posts ordered by time with links to /doc?id=... (seepost_handler)

// All handlers here add 1 to waitgroup and do .Done() in the end