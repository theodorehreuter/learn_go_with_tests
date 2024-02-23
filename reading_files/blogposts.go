package blogposts

import "testing/fstest"

type Post struct {
}

func NewPostsFromFS(fileSystem fstest.MapFS) []Post {
	return nil
}

// var posts []blogposts.Post
// posts = blogposts.NewPostsFromFS(someFS)
