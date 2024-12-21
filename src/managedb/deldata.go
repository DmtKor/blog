package managedb

// Comments related to post will be deleted automatically
func (db *DB) DeletePostById(id uint64) error {
	_, err := db.Database.Exec("DELETE FROM Post WHERE Id = $1", id)
	return err
}

func (db *DB) DeleteCommentById(id uint64) error {
	_, err := db.Database.Exec("DELETE FROM Comment WHERE Id = $1", id)
	return err
}

func (db *DB) DeleteUnusedTags() error {
	_, err := db.Database.Exec("DELETE FROM TagText WHERE Id NOT IN (SELECT DISTINCT TagId FROM Tags)")
	return err
}