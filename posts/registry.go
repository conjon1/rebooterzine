package posts

import (
	"github.com/a-h/templ"
	"github.com/conjon1/rebooterzine/model"
)

type PostEntry struct {
	Meta     model.Post
	Renderer func() templ.Component
}

var Registry = make(map[string]PostEntry)

// Register takes a post struct and the templ function, and wires them up.
func Register(meta model.Post, renderFunc func(model.Post) templ.Component) {
	// Automatically generate the URL so you don't have to type it manually
	if meta.URL == "" {
		meta.URL = "post/" + meta.ID
	}

	Registry[meta.ID] = PostEntry{
		Meta:     meta,
		Renderer: func() templ.Component { return renderFunc(meta) },
	}
}

func AllPosts() []model.Post {
	all := make([]model.Post, 0, len(Registry))
	for _, entry := range Registry {
		all = append(all, entry.Meta)
	}
	return all
}
