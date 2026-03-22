package main

import (
	"context"
	"log"
	"os"

	"github.com/conjon1/rebooterzine/components"
	"github.com/conjon1/rebooterzine/internal/builder"
	"github.com/conjon1/rebooterzine/internal/parser"
	"github.com/conjon1/rebooterzine/model"
)

func main() {
	log.Println("⚡ Starting blog build process...")

	posts, err := parser.FindPosts(model.PostsDir)
	if err != nil {
		log.Fatalf("Error parsing posts: %v", err)
	}
	builder.SortPostsByDate(posts)

	authors, err := parser.FindAuthors(model.AuthorsDir)
	if err != nil {
		log.Fatalf("Error parsing authors: %v", err)
	}

	builder.CrossLinkAuthors(posts, authors)

	if err := builder.WriteJSON(posts, model.JSONOutput); err != nil {
		log.Printf("Warning: Could not write posts JSON: %v", err)
	}
	if err := builder.WriteJSON(authors, model.AuthorsOutput); err != nil {
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
