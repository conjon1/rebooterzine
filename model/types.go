package model

// Post represents a single blog post parsed from public_html/posts/*.html.
type Post struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Date     string `json:"date"`
	Excerpt  string `json:"excerpt"`
	URL      string `json:"url"`
	Author   string `json:"author"`
	AuthorID string `json:"authorId"`
}

// AuthorData represents an author parsed from public_html/authors/*.html.
type AuthorData struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Linkedin string `json:"linkedin"`
	Github   string `json:"github"`
}

const (
	PostsDir      = "public_html/posts"
	AuthorsDir    = "public_html/authors"
	JSONOutput    = "public_html/posts.json"
	AuthorsOutput = "public_html/authors.json"
	HTMLOutput    = "public_html/index.html"
	DateFormat    = "January 2, 2006"
)
