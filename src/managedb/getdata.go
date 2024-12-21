package managedb

func (db *DB) GetAllPosts() ([]Post, error) {
	return nil, nil
}

func (db *DB) GetPostsByPage(page uint, pagesize uint) ([]Post, error) {
	return nil, nil
}

func (db *DB) GetPostById(id uint64) (Post, error) {
	return Post{}, nil
}

func (db *DB) GetCommentsByPostId(id uint64) ([]Comment, error) {
	return nil, nil
}

func (db *DB) GetAllTags() ([]string, error) {
	return nil, nil
}