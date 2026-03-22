package parser

import (
	"os"

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

// extractMeta walks the node tree and calls callback for every <meta name=... content=...> tag.
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
