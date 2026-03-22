package parser

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/conjon1/rebooterzine/model"
)

// FindAuthors scans dir for .html files and returns parsed authors.
func FindAuthors(dir string) ([]model.AuthorData, error) {
	var authors []model.AuthorData

	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return authors, nil
		}
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".html") {
			continue
		}
		author, err := parseAuthor(filepath.Join(dir, entry.Name()))
		if err != nil {
			log.Printf("Skipping author %s: %v", entry.Name(), err)
			continue
		}
		authors = append(authors, author)
	}

	return authors, nil
}

func parseAuthor(path string) (model.AuthorData, error) {
	doc, err := parseFile(path)
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
