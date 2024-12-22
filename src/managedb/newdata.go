package managedb

import "errors"

// Create new post in DB from Post struct data
func (db *DB) NewPost(post Post) (Post, error) {
	// post.Id will be ignored
	if post.Content == "" {
		return post, errors.New("empty post.content")
	}
	if post.Description == "" {
		return post, errors.New("empty post.description")
	}
	if post.Title == "" {
		return post, errors.New("empty post.title")
	}
	var postid uint64
	// Create post
	err := db.Database.QueryRow("INSERT INTO Post (Title, Description, Content, PostDate) VALUES " +
        "($1, $2, $3, $4) RETURNING Id", post.Title, post.Description, post.Content, post.PostDate).Scan(&postid)
	// Post was not created
    if err != nil{
        return post, err
    }
	post.Id = postid
	// add tags
	for _, val := range post.Tags { 
		err = db.NewTag(val)
		if err != nil {
			break
		}
		err = db.AddTagToPost(val, post)
		if err != nil {
			break
		}
	}
	// Post created successfully, but tags do not
	if err != nil {
		// Remove post
		db.Database.Exec("DELETE FROM Post WHERE Id = $1", post.Id)
		// Remove all tags created
		for _, val := range post.Tags {
			db.Database.Exec("DELETE FROM TagText WHERE tagtext = $1", val)
		}
		// Info from Tags table will be deleted anyways
		return post, err
	}
	return post, err // Returns post with updated Id
}

// Create new comment in DB from Comment struct data
func (db *DB) NewComment(comm Comment) error {
	if comm.Content == "" {
		return errors.New("empty comm.content")
	}
	if comm.Email == "" {
		return errors.New("empty comm.email")
	}
	var commid uint64
	err := db.Database.QueryRow("INSERT INTO Comment (PostId, Author, Content, CommDate, Email) VALUES " + 
		"($1, $2, $3, $4, $5) RETURNING Id", comm.PostId, comm.Author, comm.Content, comm.CommDate, comm.Email).Scan(&commid)
	// Comment was not created
	if err != nil {
		return err
	}
	return nil
}

// Tell DB that in this post is used this tag
func (db *DB) AddTagToPost(text string, post Post) error {
	_, err := db.Database.Exec("INSERT INTO Tags (PostId, TagId) VALUES ($1, (SELECT DISTINCT Id FROM" +
	" TagText WHERE tagtext = $2))", post.Id, text)
	return err
}

// Adds tag to TagText if needed, does nothing if it exists
func (db *DB) NewTag(text string) error {
	_, err := db.Database.Exec("INSERT INTO TagText (TagText) VALUES ($1) ON CONFLICT DO NOTHING", text)
	return err
}