package managedb

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// For database connection should be created this object
type DB struct {
	Database *sql.DB // DB connection object, result from open()
}

// Open connection, create tables (if they don't exist)/ 
// This function is called in commander package
func (db *DB) Init(initStr string) error {
	var err error
	db.Database, err = sql.Open("postgres", initStr)
	if err != nil {
		return err
	}
	_, err = db.Database.Exec("CREATE TABLE IF NOT EXISTS Post (" + 
							  "Id BIGSERIAL PRIMARY KEY," +
							  "Title VARCHAR(30) CHECK(Title != '')," +
							  "Description VARCHAR(200) CHECK(Description != '')," +
							  "Content TEXT CHECK(Content != '')," +
							  "PostDate TIMESTAMP WITH TIME ZONE" +
							  ")")
	if err != nil {
		return err
	}
	_, err = db.Database.Exec("CREATE TABLE IF NOT EXISTS Comment (" + 
							  "Id BIGSERIAL PRIMARY KEY," +
							  "PostId BIGINT REFERENCES Post (Id) ON DELETE CASCADE," +
							  "Author VARCHAR(30) DEFAULT 'Anonymous'," +
							  "Content VARCHAR(250) CHECK(Content != '')," +
							  "CommDate TIMESTAMP WITH TIME ZONE," +
							  "Email VARCHAR(30) CHECK(Email != '')" +
							  ")")
	if err != nil {
		return err
	}
	_, err = db.Database.Exec("CREATE TABLE IF NOT EXISTS TagText (" + 
							  "Id SERIAL PRIMARY KEY," +
							  "TagText VARCHAR(30) UNIQUE CHECK(TagText != '')" +
							  ")")
	if err != nil {
		return err
	}
	_, err = db.Database.Exec("CREATE TABLE IF NOT EXISTS Tags (" + 
							  "PostId BIGINT REFERENCES Post (Id) ON DELETE CASCADE," +
							  "TagId INT REFERENCES TagText (Id) ON DELETE CASCADE," +
							  "UNIQUE(PostId, TagId)" +
							  ")")
	if err != nil {
		return err
	}
	return nil
}

// Close DB connection
func (db *DB) Close() {
	db.Database.Close()
}