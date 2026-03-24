package parser

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/conjon1/rebooterzine/model"
)

// FindPosts scans dir for .html files and returns parsed posts.
func FindPosts(dir string) ([]model.Post, error) {
	var posts []model.Post

	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return posts, nil
		}
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".html") {
			continue
		}
		post, err := parsePost(filepath.Join(dir, entry.Name()))
		if err != nil {
			log.Printf("Skipping post %s: %v", entry.Name(), err)
			continue
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func parsePost(path string) (model.Post, error) {
	doc, err := parseFile(path)
	if err != nil {
		return model.Post{}, err
	}

	var post model.Post
	post.ID = strings.TrimSuffix(filepath.Base(path), ".html")
	post.URL = "/posts/" + filepath.Base(path)

	extractMeta(doc, func(name, content string) {
		switch name {
		case "title":
			post.Title = content
		case "date":
			post.Date = content
		case "author":
			post.Author = content
		case "excerpt":
			post.Excerpt = content
		}
	})

	if post.Author == "" {
		post.Author = "UNKNOWN"
	}

	if post.Title == "" || post.Date == "" {
		return model.Post{}, fmt.Errorf("missing title or date metadata")
	}

	return post, nil
}
