package parser

import (
	"os"
	"strings"

	"golang.org/x/net/html"
)

// parseFile opens a file and returns the parsed HTML node tree.
func parseFile(path string) (*html.Node, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return html.Parse(f)
}

// extractMeta walks the node tree and calls callback for meta tags and the title tag.
func extractMeta(n *html.Node, callback func(name, content string)) {
	if n.Type == html.ElementNode {
		// 1. Handle standard <title> tag if meta title is missing
		if n.Data == "title" && n.FirstChild != nil {
			callback("title", n.FirstChild.Data)
		}

		// 2. Handle <meta> tags with broader attribute support
		if n.Data == "meta" {
			var name, content string
			for _, a := range n.Attr {
				// Check name, property (OG tags), or itemprop
				if a.Key == "name" || a.Key == "property" || a.Key == "itemprop" {
					name = a.Val
				}
				if a.Key == "content" {
					content = a.Val
				}
			}
			if name != "" && content != "" {
				// Normalize keys (e.g., "og:title" -> "title")
				name = strings.ToLower(name)
				name = strings.TrimPrefix(name, "og:")
				callback(name, content)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractMeta(c, callback)
	}
}
