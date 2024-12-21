package managedb

import "time"

type Post struct {
    Id uint64          // - Post number
    Title string      
    Description string // - Short description
    Content string     
    PostDate time.Time 
    Tags []string  
}

type Comment struct {
    Id uint64          // - comment id
    PostId uint64   
    Author string      // - Just name, optional, no real 'users'
    Content string     
    CommDate time.Time 
    Email string       
}
