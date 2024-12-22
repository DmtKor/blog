package managedb

import (
	"testing"
	"time"
)

func TestDB(t *testing.T) {
	// Change this data to actual user name and password of your db
	initStr := initstr_temp // If you want to run this test, paste here string like "user=your_username password=your_password dbname=name_of_database sslmode=disable"
	db := DB{}
	err := db.Init(initStr)
	if err != nil {
		t.Fatal("error initialising DB,", err)
	}
	defer db.Close()
	posts := make([]Post, 0)


	// new data

	post, err := db.NewPost(Post{ Title: "title 1", Description: "description 1", Content: "content 1", 
		PostDate: time.Now(), Tags: []string{ "tag a", "tag b", "tag c", "tag d" } })
	if err != nil {
		t.Error("error adding new post 1,", err)
	} else {
		posts = append(posts, post)
	}

	post, err = db.NewPost(Post{ Title: "title 2", Description: "description 2", Content: "content 2", 
		PostDate: time.Now(), Tags: []string{ "tag a", "tag c" } })
	if err != nil {
		t.Error("error adding new post 2,", err)
	} else {
		posts = append(posts, post)
	}

	post, err = db.NewPost(Post{ Title: "title 3", Description: "description 3", Content: "content 3", 
		PostDate: time.Now(), Tags: []string{ "tag b" } })
	if err != nil {
		t.Error("error adding new post 3,", err)
	} else {
		posts = append(posts, post)
	}

	post, err = db.NewPost(Post{ Title: "title 4", Description: "description 4", Content: "content 4", 
		PostDate: time.Now(), Tags: []string{ } })
	if err != nil {
		t.Error("error adding new post 4,", err)
	} else {
		posts = append(posts, post)
	}

	// Check if post exists

	val, err := db.PostExists(post.Id)
	if err != nil {
		t.Log("cannot check if comment exists: ", err)
	} else {
		t.Log("Post", post.Id, "exists? :", val)
	}

	err = db.NewComment(Comment{ PostId: posts[0].Id, Author: "a1", Content: "c1", Email: "e1@mail.com", CommDate: time.Now() })
	if err != nil {
		t.Error("error adding new comment 1,", err)
	}

	err = db.NewComment(Comment{ PostId: posts[1].Id, Author: "a2", Content: "c2", Email: "e2@mail.com", CommDate: time.Now() })
	if err != nil {
		t.Error("error adding new comment 2,", err)
	}

	err = db.NewComment(Comment{ PostId: posts[2].Id, Author: "a3", Content: "c3", Email: "e3@mail.com", CommDate: time.Now() })
	if err != nil {
		t.Error("error adding new comment 3,", err)
	}

	err = db.NewComment(Comment{ PostId: posts[3].Id, Author: "a4", Content: "c4", Email: "e4@mail.com", CommDate: time.Now() })
	if err != nil {
		t.Error("error adding new comment 4,", err)
	}

	err = db.NewComment(Comment{ PostId: posts[3].Id, Author: "a5", Content: "c5", Email: "e5@mail.com", CommDate: time.Now() })
	if err != nil {
		t.Log("wrong comment was not created. ", err)
	}

	// Check if comment exists

	val, err = db.CommExists(2)
	if err != nil {
		t.Log("cannot check if comment exists: ", err)
	} else {
		t.Log("Comment 2 exists? :", val)
	}

	// get data

	posts, err = db.GetAllPosts()
	if err != nil {
		t.Error("error reading posts,", err)
	} else {
		for i, val := range posts {
			t.Log("Post", i, ">", val )
		}
	}

	l, err := db.GetPostsNum()
	t.Log("Posts number:", l)
	if err != nil {
		t.Error("error reading posts number,", err)
	} else {
		pagesize := uint64(2)
		pagenum, err := db.GetPageNum(pagesize)
		if err != nil {
			t.Error("Unable to get number of pages")
		} else {
			t.Log("Number of pages:", pagenum)
			for i := uint64(0); i <= (pagenum) + 1; i++ {
				posts, err := db.GetPostsByPage(uint(i), uint(pagesize))
				if err != nil {
					if i != 0 {
						t.Error("error getting posts from page,", i)
					} else {
						t.Log("posts from page 0 were skipped as intended")
					}
				} else {
					t.Log("Page", i, ">", posts)
				}
			}
		}
	}
	posts, _ = db.GetAllPosts()

	tags, err := db.GetAllTags()
	if err != nil {
		t.Error("error accessing tags,", err)
	} else {
		t.Log("All tags:", tags)
	}

	post, error := db.GetPostById(posts[0].Id)
	if error != nil {
		t.Error("error accessing post ", posts[0].Id, ",", err)
	} else {
		t.Log("Post 1:", post)
		comms, err := db.GetCommentsByPostId(posts[0].Id)
		if err != nil {
			t.Error("error accessing comments of post", posts[0].Id, ",", err)
		} else {
			t.Log("Comments for post", posts[0].Id,":", comms)
		}
	}


	// delete data

	for i := uint64(0); i < 1000; i++ {
		err := db.DeletePostById(i)
		if err != nil {
			if i != 0 {
				t.Error("Error deleting post", i, err)
			} else {
				t.Log("Post 0 was not deleted as intended")
			}
		}
	}
	db.DeleteUnusedTags()

	// Checking if data really was deleted

	var nposts, ncomms, ntags, ntagtexts int
	db.Database.QueryRow("SELECT COUNT(*) FROM Post").Scan(nposts)
	if nposts != 0 {
		t.Error("Not all posts were deleted!")
	}
	db.Database.QueryRow("SELECT COUNT(*) FROM Comment").Scan(ncomms)
	if ncomms != 0 {
		t.Error("Not all comments were deleted!")
	}
	db.Database.QueryRow("SELECT COUNT(*) FROM Tags").Scan(ntags)
	if ntags != 0 {
		t.Error("Not all tag references were deleted!")
	}
	db.Database.QueryRow("SELECT COUNT(*) FROM TagText").Scan(ntagtexts)
	if ntagtexts != 0 {
		t.Error("Not all tags were deleted!")
	}

	// Removing tables
	db.Database.Exec("DROP TABLE Post")
	db.Database.Exec("DROP TABLE Comment")
	db.Database.Exec("DROP TABLE Tags")
	db.Database.Exec("DROP TABLE TagText")
}