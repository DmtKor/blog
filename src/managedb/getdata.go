package managedb

func (db *DB) GetAllPosts() ([]Post, error) {
	rows, err := db.Database.Query("SELECT Id, Title, Description, Content, PostDate FROM Post ORDER BY PostDate DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := make([]Post, 0)
	for rows.Next() {
		p := Post{}
		err = rows.Scan(&p.Id, &p.Title, &p.Description, &p.Content, &p.PostDate)
		if err != nil {
			continue
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func (db *DB) GetPostsByPage(page uint, pagesize uint) ([]Post, error) {
	rows, err := db.Database.Query("SELECT Id, Title, Description, Content, PostDate FROM Post " + 
		"ORDER BY PostDate DESC LIMIT $1 OFFSET $2", pagesize, pagesize * (page - 1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := make([]Post, 0)
	for rows.Next() {
		p := Post{}
		err = rows.Scan(&p.Id, &p.Title, &p.Description, &p.Content, &p.PostDate)
		if err != nil {
			continue
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func (db *DB) GetPostsNum() (uint64, error) {
	var res uint64
	err := db.Database.QueryRow("SELECT LEN(*) FROM Post").Scan(&res)
	return res, err
}

func (db *DB) GetPostById(id uint64) (Post, error) {
	var p Post
	err := db.Database.QueryRow("SELECT Id, Title, Description," + 
		" Content, PostDate FROM Post WHERE PostId = $1", id).Scan(&p.Id, &p.Title, &p.Description, &p.Content, &p.PostDate)
	return p, err
}

func (db *DB) GetCommentsByPostId(id uint64) ([]Comment, error) {
	rows, err := db.Database.Query("SELECT Id, PostId, Author, Content, CommDate, Email FROM Comment " + 
		"WHERE PostId = $1 ORDER BY CommDate DESC", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	comms := make([]Comment, 0)
	for rows.Next() {
		c := Comment{}
		err = rows.Scan(&c.Id, &c.PostId, &c.Author, &c.Content, &c.CommDate, &c.Email)
		if err != nil {
			continue
		}
		comms = append(comms, c)
	}
	return comms, nil
}

func (db *DB) GetAllTags() ([]string, error) {
	rows, err := db.Database.Query("SELECT tagtext FROM TagText")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := make([]string, 0)
	var t string
	for rows.Next() {
		err = rows.Scan(&t)
		if err != nil {
			continue
		}
		res = append(res, t)
	}
	return res, nil
}