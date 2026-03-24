package builder

import (
	"sort"
	"time"

	"github.com/conjon1/rebooterzine/model"
)

// SortPostsByDate sorts posts newest-first, in place.
func SortPostsByDate(posts []model.Post) {
	sort.Slice(posts, func(i, j int) bool {
		t1, _ := time.Parse(model.DateFormat, posts[i].Date)
		t2, _ := time.Parse(model.DateFormat, posts[j].Date)
		return t1.After(t2)
	})
}

// CrossLinkAuthors attaches an AuthorID to each post by matching Author name.
func CrossLinkAuthors(posts []model.Post, authors []model.AuthorData) {
	authorMap := make(map[string]string, len(authors))
	for _, a := range authors {
		authorMap[a.Name] = a.ID
	}
	for i := range posts {
		if id, ok := authorMap[posts[i].Author]; ok {
			posts[i].AuthorID = id
		}
	}
}
