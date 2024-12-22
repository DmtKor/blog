package managedb

// Delete post from Post table (and associated tags and comments).
// Data from TagText table should be deleted separately in DeleteUnusedTags
func (db *DB) DeletePostById(id uint64) error {
	_, err := db.Database.Exec("DELETE FROM Post WHERE Id = $1", id)
	return err
}

// Delete comment using its unique ID in Comments table
func (db *DB) DeleteCommentById(id uint64) error {
	_, err := db.Database.Exec("DELETE FROM Comment WHERE Id = $1", id)
	return err
}

// Delete unused tag info from TagText table
func (db *DB) DeleteUnusedTags() error {
	_, err := db.Database.Exec("DELETE FROM TagText WHERE Id NOT IN (SELECT DISTINCT TagId FROM Tags)")
	return err
}