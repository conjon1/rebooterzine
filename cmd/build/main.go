package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/conjon1/rebooterzine/components"
	"github.com/conjon1/rebooterzine/internal/builder"
	"github.com/conjon1/rebooterzine/internal/parser"
	"github.com/conjon1/rebooterzine/model"
	"github.com/conjon1/rebooterzine/posts"
)

func main() {
	log.Println("⚡ Starting blog build process...")

	// Posts now come from the templ registry, not the HTML file scanner.
	allPosts := posts.AllPosts()
	builder.SortPostsByDate(allPosts)

	authors, err := parser.FindAuthors(model.AuthorsDir)
	if err != nil {
		log.Fatalf("Error parsing authors: %v", err)
	}

	builder.CrossLinkAuthors(allPosts, authors)

	if err := builder.WriteJSON(allPosts, model.JSONOutput); err != nil {
		log.Printf("Warning: Could not write posts JSON: %v", err)
	}
	if err := builder.WriteJSON(authors, model.AuthorsOutput); err != nil {
		log.Printf("Warning: Could not write authors JSON: %v", err)
	}

	// Render index.html
	f, err := os.Create(model.HTMLOutput)
	if err != nil {
		log.Fatalf("Error creating index.html: %v", err)
	}
	defer f.Close()

	if err := components.Index(allPosts, authors).Render(context.Background(), f); err != nil {
		log.Fatalf("❌ Error rendering index: %v", err)
	}

	// Render one HTML file per post → public_html/post/<slug>.html
	postOutDir := "public_html/post"
	if err := os.MkdirAll(postOutDir, 0755); err != nil {
		log.Fatalf("Could not create %s: %v", postOutDir, err)
	}

	for slug, entry := range posts.Registry {
		outPath := fmt.Sprintf("%s/%s.html", postOutDir, slug)
		pf, err := os.Create(outPath)
		if err != nil {
			log.Printf("Warning: could not create %s: %v", outPath, err)
			continue
		}
		if err := entry.Renderer().Render(context.Background(), pf); err != nil {
			log.Printf("Warning: could not render post %s: %v", slug, err)
		}
		pf.Close()
		log.Printf("  wrote %s", outPath)
	}

	log.Printf("\n✅ Nooice! Built site with %d posts and %d authors.", len(allPosts), len(authors))
}
