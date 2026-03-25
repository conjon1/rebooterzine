package components

import "github.com/conjon1/rebooterzine/model"

// filters a list of posts and returns only those written by the specified author.
func GetPostsByAuthor(posts []model.Post, authorName string) []model.Post {
	var filtered []model.Post
	for _, p := range posts {
		if p.Author == authorName {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

// assumes the posts slice is already sorted by date.
func GetRecent(posts []model.Post, n int) []model.Post {
	if len(posts) <= n {
		return posts
	}
	return posts[:n]
}

