package components

import "github.com/conjon1/rebooterzine/model"

// GetRecent returns up to count posts from the front of the sorted slice.
func GetRecent(posts []model.Post, count int) []model.Post {
	if len(posts) < count {
		return posts
	}
	return posts[:count]
}

// GetPostsByAuthor filters posts by exact author display name match.
func GetPostsByAuthor(posts []model.Post, authorName string) []model.Post {
	var ap []model.Post
	for _, p := range posts {
		if p.Author == authorName {
			ap = append(ap, p)
		}
	}
	return ap
}
