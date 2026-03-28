package posts

import (
	"github.com/a-h/templ"
	"github.com/conjon1/rebooterzine/model"
)

// PostEntry ties a slug to its metadata and templ renderer.
// Add a new line here every time you create a new post .templ file.
type PostEntry struct {
	Meta     model.Post
	Renderer func() templ.Component
}

// Registry is the single source of truth for all posts.
// Key = the slug used in the URL: /post/<slug>
var Registry = map[string]PostEntry{
	"gnet": {
		Meta: model.Post{
			ID:     "gnet",
			Title:  "Tunneling All System Traffic Through Your Phone with SSH, redsocks, and Mullvad",
			Date:   "March 13, 2026",
			Author: "Connal McInnis",
			AuthorID: "connalmcinnis",
			Excerpt: "Routes all system TCP traffic through a SOCKS5 tunnel over USB tethering to an Android phone running Mullvad VPN.",
			URL:    "post/gnet",
		},
		Renderer: func() templ.Component { return Gnet() },
	},
	"TmobileSpoof": {
		Meta: model.Post{
			ID:     "TmobileSpoof",
			Title:  "T-mobile TTL spoofing",
			Date:   "January 6, 2026",
			Author: "Connal McInnis",
			AuthorID: "connalmcinnis",
			Excerpt: "Right now i am in a rural area where it does not make sense to have WiFi if someone's mobile data is unlimited.",
			URL:    "post/TmobileSpoof",
		},
		Renderer: func() templ.Component { return TmobileSpoof() },
	},
	// ── Add new posts here ───────────────────────────────────────────────────
	// "my-new-post": {
	//     Meta: model.Post{ ... },
	//     Renderer: func() templ.Component { return MyNewPost() },
	// },
}

// AllPosts returns all post metadata as a slice, for the index/blog listing.
func AllPosts() []model.Post {
	posts := make([]model.Post, 0, len(Registry))
	for _, entry := range Registry {
		posts = append(posts, entry.Meta)
	}
	return posts
}
