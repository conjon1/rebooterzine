package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"golang.org/x/net/html"

	"github.com/YOUR_USERNAME/rebooter/components"
	"github.com/YOUR_USERNAME/rebooter/model"
)

func main() {
	log.Println("⚡ Starting blog build process...")

	posts, err := findAndParsePosts(model.PostsDir)
	if err != nil {
		log.Fatalf("Error parsing posts: %v", err)
	}
	sortPostsByDate(posts)

	authors, err := findAndParseAuthors(model.AuthorsDir)
	if err != nil {
		log.Fatalf("Error parsing authors: %v", err)
	}

	// Cross-link posts to author IDs
	authorMap := make(map[string]string)
	for _, a := range authors {
		authorMap[a.Name] = a.ID
	}
	for i := range posts {
		if id, ok := authorMap[posts[i].Author]; ok {
			posts[i].AuthorID = id
		}
	}

	if err := writeJSON(posts, model.JSONOutput); err != nil {
		log.Printf("Warning: Could not write posts JSON: %v", err)
	}
	if err := writeJSON(authors, model.AuthorsOutput); err != nil {
		log.Printf("Warning: Could not write authors JSON: %v", err)
	}

	f, err := os.Create(model.HTMLOutput)
	if err != nil {
		log.Fatalf("Error creating index.html: %v", err)
	}
	defer f.Close()

	if err := components.Index(posts, authors).Render(context.Background(), f); err != nil {
		log.Fatalf("❌ Error rendering template: %v", err)
	}

	log.Printf("\n✅ Nooice! Built site with %d posts and %d authors.", len(posts), len(authors))
}

// ── Scanning ──────────────────────────────────────────────────────────────────

func findAndParsePosts(dir string) ([]model.Post, error) {
	var posts []model.Post
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".html") {
			post, err := parsePostHTML(filepath.Join(dir, entry.Name()))
			if err != nil {
				log.Printf("Skipping post %s: %v", entry.Name(), err)
				continue
			}
			posts = append(posts, post)
		}
	}
	return posts, nil
}

func findAndParseAuthors(dir string) ([]model.AuthorData, error) {
	var authors []model.AuthorData
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return authors, nil
		}
		return nil, err
	}
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".html") {
			author, err := parseAuthorHTML(filepath.Join(dir, entry.Name()))
			if err != nil {
				log.Printf("Skipping author %s: %v", entry.Name(), err)
				continue
			}
			authors = append(authors, author)
		}
	}
	return authors, nil
}

// ── Sorting ───────────────────────────────────────────────────────────────────

func sortPostsByDate(posts []model.Post) {
	sort.Slice(posts, func(i, j int) bool {
		t1, _ := time.Parse(model.DateFormat, posts[i].Date)
		t2, _ := time.Parse(model.DateFormat, posts[j].Date)
		return t1.After(t2)
	})
}

// ── Serialisation ─────────────────────────────────────────────────────────────

func writeJSON(data interface{}, path string) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, bytes, 0644)
}

// ── HTML Parsing ──────────────────────────────────────────────────────────────

func parsePostHTML(path string) (model.Post, error) {
	doc, err := parseHTMLFile(path)
	if err != nil {
		return model.Post{}, err
	}

	var post model.Post
	post.ID  = strings.TrimSuffix(filepath.Base(path), ".html")
	post.URL = fmt.Sprintf("posts/%s", filepath.Base(path))

	extractMeta(doc, func(name, content string) {
		switch name {
		case "date":
			post.Date = content
		case "excerpt":
			post.Excerpt = content
		case "author":
			post.Author = content
		}
	})

	var crawler func(*html.Node)
	crawler = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			post.Title = n.FirstChild.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			crawler(c)
		}
	}
	crawler(doc)

	if post.Title == "" || post.Date == "" {
		return model.Post{}, fmt.Errorf("missing title or date metadata")
	}
	if post.Author == "" {
		post.Author = "UNKNOWN"
	}
	return post, nil
}

func parseAuthorHTML(path string) (model.AuthorData, error) {
	doc, err := parseHTMLFile(path)
	if err != nil {
		return model.AuthorData{}, err
	}

	var author model.AuthorData
	author.ID = strings.TrimSuffix(filepath.Base(path), ".html")

	extractMeta(doc, func(name, content string) {
		switch name {
		case "name":
			author.Name = content
		case "linkedin":
			author.Linkedin = content
		case "github":
			author.Github = content
		}
	})

	if author.Name == "" {
		return model.AuthorData{}, fmt.Errorf("missing name metadata")
	}
	return author, nil
}

func parseHTMLFile(path string) (*html.Node, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return html.Parse(f)
}

func extractMeta(n *html.Node, callback func(name, content string)) {
	if n.Type == html.ElementNode && n.Data == "meta" {
		name, content := "", ""
		for _, a := range n.Attr {
			if a.Key == "name" {
				name = a.Val
			}
			if a.Key == "content" {
				content = a.Val
			}
		}
		if name != "" && content != "" {
			callback(name, content)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractMeta(c, callback)
	}
}
